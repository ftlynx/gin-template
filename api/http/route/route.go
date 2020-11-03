package route

import (
	"fmt"
	"gin-template/api/http/middleware"
	"gin-template/internal/global"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func MyRoute() error {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(middleware.AddCorsHeader)
	r.Use(middleware.PrometheusMetric)
	r.Use(middleware.GetPanic)
	r.Use(middleware.ExceptionJson)
	userApi := r.Group(ApiPrefix)
	for _, v := range apis {
		switch v.Group {
		case ApiPrefix:
			switch v.Method {
			case http.MethodGet:
				userApi.GET(v.Uri, v.Handler)
			case http.MethodPost:
				userApi.POST(v.Uri, v.Handler)
			case http.MethodPut:
				userApi.PUT(v.Uri, v.Handler)
			case http.MethodDelete:
				userApi.DELETE(v.Uri, v.Handler)
			default:
				return fmt.Errorf("%s method not support", v.Name)
			}
		default:
			return fmt.Errorf("%s group not support", v.Name)
		}
	}
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return r.Run(global.Conf.App.Listen)
}
