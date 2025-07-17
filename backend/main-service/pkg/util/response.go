package util

import "github.com/gin-gonic/gin"

func APIResponse(message string, code int, data interface{}) gin.H {
	return gin.H{
		"message": message,
		"code":    code,
		"data":    data,
	}
}