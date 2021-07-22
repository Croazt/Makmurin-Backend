package utils

import (
	"github.com/gin-gonic/gin"
)

type JsonFormat struct {
	Success bool        `json:"success" form:"success" binding:"required"`
	Data    interface{} `json:"data" form:"data" binding:"required"`
	Message string      `json:"Message" form:"message" binding:"required"`
}

func ErrorHandler(c *gin.Context, Message string, Status int) {
	c.AbortWithStatusJSON(Status, JsonFormat{
		Success: false,
		Data:    nil,
		Message: Message,
	})
}

func SuccessHandler(c *gin.Context, Data interface{}, Message string) {
	c.JSON(200, JsonFormat{
		Success: true,
		Data:    Data,
		Message: Message,
	})
}
