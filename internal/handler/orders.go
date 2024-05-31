package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)



func (h *Handler) GetOrder(c *gin.Context) {
	order_uid := c.Param("order_uid")
	data, ok := h.cache.Get(order_uid)
	if !ok {
		err := errors.New("no such order with this uid")
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(200, data)
}