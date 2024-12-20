package handler

import (
	"github.com/Njrctr/javacode_test_golang_junior/pkg/service"
	"github.com/gin-gonic/gin"
	// _ "github.com/Njrctr/javacode_test_golang_junior/docs"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			wallet := v1.Group("/wallet")
			{
				wallet.POST("/", h.updateWallet)
			}

			wallets := v1.Group("/wallets")
			{
				wallets.GET("/:wallet_uuid", h.getWalletById)
			}

		}

	}
	return router
}
