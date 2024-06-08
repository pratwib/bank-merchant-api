package main

import (
	"bank-merchant-api/controller"
	"bank-merchant-api/middleware"
	"bank-merchant-api/model/domain"
	"bank-merchant-api/repository"
	"bank-merchant-api/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=bank_merchant sslmode=disable password=wibietea234")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.AutoMigrate(&domain.Customer{}, &domain.Merchant{}, &domain.History{})

	router := gin.Default()
	router.Use(middleware.Logger())

	jwtSecret := "your_jwt_secret"
	customerRepo := repository.NewCustomerRepository(db)
	merchantRepo := repository.NewMerchantRepository(db)
	historyRepo := repository.NewHistoryRepository(db)
	customerService := service.NewCustomerService(customerRepo, historyRepo, jwtSecret)
	merchantService := service.NewMerchantService(merchantRepo, historyRepo)
	customerController := controller.NewCustomerController(customerService)
	merchantController := controller.NewMerchantController(merchantService)

	router.POST("/auth/register", customerController.Register)
	router.POST("/auth/login", customerController.Login)
	router.POST("/customers/balance/add", middleware.Auth(jwtSecret), customerController.AddBalance)
	router.POST("/customers/payment/pay", middleware.Auth(jwtSecret), customerController.Payment)

	router.POST("/merchants", middleware.Auth(jwtSecret), merchantController.CreateMerchant)
	router.GET("/merchants/:id", middleware.Auth(jwtSecret), merchantController.GetMerchant)
	router.DELETE("/merchants/:id", middleware.Auth(jwtSecret), merchantController.DeleteMerchant)

	router.Run(":8080")
}
