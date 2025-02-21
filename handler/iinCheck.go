package handler

import (
	"net/http"

	sr "Proger30/task/utils/service.response"

	"github.com/gin-gonic/gin"
)

func (h *Handler) IinCheck(c *gin.Context) {
	response, err := h.Service.IinCheck(c, c.Param("iin"))
	if err != nil {
		// log here...
		c.JSON(http.StatusUnprocessableEntity, sr.Error(-11, err.Error())) // ToDo: правильный ответ ошибки
		return
	}
	c.JSON(http.StatusOK, sr.OkWithData("success", response))
}
