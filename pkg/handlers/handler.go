package handler

import (
	_ "github.com/Njrctr/javacode_test_golang_junior/docs"
	"github.com/Njrctr/javacode_test_golang_junior/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentify)
	{
		v1 := api.Group("/v1")
		{
			wallet := v1.Group("/wallet")
			{
				// wallet.POST("/", h.depositOrWithdraw)
				wallet.POST("/", h.updateWallet)
				wallet.POST("/new", h.createWallet)
				wallet.GET("/", h.getAllWallets)
				wallet.DELETE("/:wallet_uuid", h.deleteWallet)
			}

			wallets := v1.Group("/wallets")
			{
				wallets.GET("/:wallet_uuid", h.GetBalanceByUUID)
			}

			admin := v1.Group("/admin", h.adminIdentify)
			{
				wallet := admin.Group("/wallet")
				{
					wallet.POST("/", h.updateWalletAdmin)
					wallet.POST("/new", h.createWalletToUser)
					wallet.GET("/:user_id", h.getAllWalletsByUser)
					wallet.PUT("/block", h.blockWallet)
				}

				wallets := admin.Group("/wallets")
				{
					wallets.GET("/:wallet_uuid", h.getWalletByUUID)
				}
			}

		}

	}
	return router
}
