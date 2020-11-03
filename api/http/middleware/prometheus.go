package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"strings"
	"time"
)

var (
	HttpReqDuration *prometheus.HistogramVec
	HttpReqTotal    *prometheus.CounterVec
)

func init() {
	HttpReqDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:        "http_request_duration_seconds",
		Help:        "The HTTP request latencies in secondes.",
		ConstLabels: nil,
		Buckets:     nil,
	}, []string{"method", "path"})

	HttpReqTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests made.",
	}, []string{"method", "path", "status"})

	prometheus.MustRegister(
		HttpReqDuration,
		HttpReqTotal,
	)
}

func parsePath(path string) string {
	itemList := strings.Split(path, "/")
	if len(itemList) > 4 {
		return strings.Join(itemList[0:3], "/")
	}
	return path
}


func PrometheusMetric(c *gin.Context) {
	tBegin := time.Now()
	c.Next()

	duration := float64(time.Since(tBegin)) / float64(time.Second)
	path := parsePath(c.Request.URL.Path)

	//请求计数+1
	HttpReqTotal.With(prometheus.Labels{
		"method": c.Request.Method,
		"path": path,
		"status": strconv.Itoa(c.Writer.Status()),
	}).Inc()
	//记录本身请求处理时间
	HttpReqDuration.With(prometheus.Labels{
		"method": c.Request.Method,
		"path": path,
	}).Observe(duration)
}