package handler

import (
	"github.com/RbPyer/WB0/internal/cache"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	cache *cache.Cache
}


func NewHandler(cache *cache.Cache) *Handler {
	return &Handler{cache: cache}
}



func (h *Handler) InitRouting() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		orders := api.Group("/orders") 
		{
			orders.GET("/:order_uid", h.GetOrder)
		}

	}

	return router
}