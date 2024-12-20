package handler

import (
	"github.com/Njrctr/javacode_test_golang_junior/pkg/service"
	"github.com/gin-gonic/gin"

	_ "github.com/Njrctr/javacode_test_golang_junior/docs"
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

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// v1.POST("/", h.createList)
			// v1.GET("/", h.getAllLists)
			// v1.GET("/:list_id", h.getListById)
			// v1.PUT("/:list_id", h.updateList)
			// v1.DELETE("/:list_id", h.deleteList)

			wallets := api.Group("/wallets")
			{
				// wallets.GET("/:item_id", h.getItemById)
				// wallets.PUT("/:item_id", h.updateItem)
				// wallets.DELETE("/:item_id", h.deleteItem)
			}

		}

	}
	return router
}
