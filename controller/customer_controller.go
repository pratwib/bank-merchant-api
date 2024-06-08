package controller

import (
	"bank-merchant-api/model/domain"
	"bank-merchant-api/model/web"
	"bank-merchant-api/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type CustomerController struct {
	service service.CustomerService
}

func NewCustomerController(service service.CustomerService) *CustomerController {
	return &CustomerController{service: service}
}

func (controller *CustomerController) Register(ctx *gin.Context) {
	var req web.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.Response{Status: "error", Message: err.Error()})
		return
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.Response{Status: "error", Message: err.Error()})
		return
	}
	customer := domain.Customer{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Balance:  0.0,
	}
	if err := controller.service.Register(&customer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, web.Response{Status: "success", Message: "customer registered"})
}

func (controller *CustomerController) Login(ctx *gin.Context) {
	var req web.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.Response{Status: "error", Message: err.Error()})
		return
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.Response{Status: "error", Message: err.Error()})
		return
	}
	token, err := controller.service.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, web.Response{Status: "error", Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, web.Response{Status: "success", Message: "logged in", Data: gin.H{"token": token}})
}

func (controller *CustomerController) AddBalance(ctx *gin.Context) {
	var req web.AddBalanceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.Response{Status: "error", Message: err.Error()})
		return
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.Response{Status: "error", Message: err.Error()})
		return
	}
	if err := controller.service.AddBalance(req.CustomerID, req.Amount); err != nil {
		ctx.JSON(http.StatusInternalServerError, web.Response{Status: "error", Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, web.Response{Status: "success", Message: "balance added"})
}

func (controller *CustomerController) Payment(ctx *gin.Context) {
	var req web.PaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.Response{Status: "error", Message: err.Error()})
		return
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.Response{Status: "error", Message: err.Error()})
		return
	}
	if err := controller.service.ProcessPayment(req.CustomerID, req.MerchantID, req.Amount); err != nil {
		ctx.JSON(http.StatusInternalServerError, web.Response{Status: "error", Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, web.Response{Status: "success", Message: "payment processed"})
}
