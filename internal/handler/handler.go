package handler

import (
	"github.com/RbPyer/WB0/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}


func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}



func (h *Handler) InitRouting() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		orders := api.Group("/orders") 
		{
			orders.GET("/:id", h.GetOrders)
		}

	}

	return router
}