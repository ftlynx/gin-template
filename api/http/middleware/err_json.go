package middleware

import (
	"gin-template/api/http/exception"
	"gin-template/api/http/response"
	"github.com/gin-gonic/gin"
	"net/http"
)


/*
exception处理，用于解决用户请求返回为文本的情况，需要转换成json
比如说请求一个不存在的接口，gin默认返回 404 page 文本
*/
func ExceptionJson(c *gin.Context) {
	c.Next()
	switch c.Writer.Status() {
	case http.StatusNotFound:
		c.Abort()
		response.JSON(c, exception.New404())
	case http.StatusMethodNotAllowed:
		c.Abort()
		response.JSON(c, exception.New405())
	default:
		return
	}
	return
}
