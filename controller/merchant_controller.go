package controller

import (
	"bank-merchant-api/model/domain"
	"bank-merchant-api/model/web"
	"bank-merchant-api/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type MerchantController struct {
	service service.MerchantService
}

func NewMerchantController(service service.MerchantService) *MerchantController {
	return &MerchantController{service: service}
}

func (controller *MerchantController) CreateMerchant(ctx *gin.Context) {
	var req web.MerchantRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.Response{Status: "error", Message: err.Error()})
		return
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.Response{Status: "error", Message: err.Error()})
		return
	}
	merchant := domain.Merchant{
		Name:    req.Name,
		Balance: 0,
	}
	if err := controller.service.CreateMerchant(&merchant); err != nil {
		ctx.JSON(http.StatusInternalServerError, web.Response{Status: "error", Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, web.Response{Status: "success", Message: "merchant created"})
}

func (controller *MerchantController) GetMerchant(ctx *gin.Context) {
	id := ctx.Param("id")
	merchant, err := controller.service.GetMerchant(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.Response{Status: "error", Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, merchant)
}

func (controller *MerchantController) DeleteMerchant(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := controller.service.DeleteMerchant(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, web.Response{Status: "error", Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, web.Response{Status: "success", Message: "merchant deleted"})
}
