package handler

import (
	"gin-template/api/http/exception"
	"gin-template/api/http/response"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	response.JSON(c, exception.NewOk())
}