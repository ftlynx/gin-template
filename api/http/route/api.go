package route

import (
	"gin-template/api/http/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ApiPrefix = "/gin-template/v1"
)

type Api struct {
	Name    string          `json:"name"`
	Method  string          `json:"method"`
	Uri     string          `json:"uri"`
	Comment string          `json:"comment"`
	Handler gin.HandlerFunc `json:"-"`
	Group   string          `json:"-"`
}

var apis = []Api{
	{Name: "ping", Group: ApiPrefix, Method: http.MethodPost, Uri: "/ping", Handler: handler.Ping, Comment: "ping"},
}
