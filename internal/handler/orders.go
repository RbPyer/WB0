package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)



func (h *Handler) GetOrder(c *gin.Context) {
	data, ok := h.cache.Get(c.Param("order_uid"))
	if !ok {
		err := errors.New("no such order with this uid")
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(200, data)
}