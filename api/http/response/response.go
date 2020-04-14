package response

import (
	"gin-template/api/http/exception"
	"github.com/gin-gonic/gin"
)

type Response struct {
	HttpCode    int         `json:"-"`
	Code        int         `json:"code"`
	CodeExplain string      `json:"code_explain,omitempty"` //错误信息，前端展示
	Message     string      `json:"message,omitempty"`      //错误信息，用于debug
	Data        interface{} `json:"data,omitempty"`
}

// PageData 数据分页数据
type PageData struct {
	PageSize   int         `json:"page_size"`   // 总共多少页
	TotalCount int         `json:"total_count"` // 总共多少条
	PageIndex  int         `json:"page_index"`  // 当前页
	Start      int         `json:"start"`       // 开始位置
	End        int         `json:"end"`         // 结束位置
	List       interface{} `json:"list"`        // 页面数据
}

/*
	将gin的c.JSON封装一层
*/
func JSON(c *gin.Context, e exception.Exception, data ...interface{}) {
	resp := Response{
		HttpCode:    e.HttpCode(),
		Code:        e.Code(),
		CodeExplain: e.Explain(),
		Message:     e.Error(),
	}
	if len(data) > 0 {
		resp.Data = data[0]
	}
	c.JSON(resp.HttpCode, resp)
}


