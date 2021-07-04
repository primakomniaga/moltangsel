package utils

import (
	"github.com/gin-gonic/gin"
)

func ResponseError(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"status":  false,
		"message": msg,
		"code":    code,
	})
}

func ResponseSuccess(c *gin.Context, code int, payload interface{}) {
	c.JSON(code, payload)
}
