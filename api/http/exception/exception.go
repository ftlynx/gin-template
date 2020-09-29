package exception

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

type Exception interface {
	error
	Code() int
	HttpCode() int
	Explain() string
	SetMessage(string, ...interface{})
	SetExplain(string, ...interface{})
}

type exception struct {
	code     int
	httpCode int
	message  string
	explain  string
}

func (e *exception) Error() string {
	return e.message
}

func (e *exception) Code() int {
	return e.code
}

func (e *exception) HttpCode() int {
	return e.httpCode
}

func (e *exception) Explain() string {
	return e.explain
}

func (e *exception) SetMessage(format string, a ...interface{}) {
	e.message = fmt.Sprintf(format, a...)
	return
}

func (e *exception) SetExplain(format string, a ...interface{}) {
	e.explain = fmt.Sprintf(format, a...)
	return
}

//正常
func NewOk() Exception {
	code := CodeOk
	return &exception{
		httpCode: codeMap[code].httpCode,
		code:     code,
		explain:  codeMap[code].explain,
	}
}

//参数错误
func NewBadRequest(message string) Exception {
	code := CodeBadRequest
	return &exception{
		httpCode: codeMap[code].httpCode,
		code:     code,
		message:  message,
		explain:  codeMap[code].explain,
	}
}


func New404() Exception {
	return &exception{
		httpCode: http.StatusNotFound,
		code:     http.StatusNotFound,
		message:  "",
		explain:  "404 page not found",
	}
}

func New405() Exception {
	return &exception{
		httpCode: http.StatusMethodNotAllowed,
		code:     http.StatusMethodNotAllowed,
		message:  "",
		explain:  "405 method not allowed",
	}
}
func New500() Exception {
	return &exception{
		httpCode: http.StatusInternalServerError,
		code:     http.StatusInternalServerError,
		message:  "",
		explain:  "internal server error",
	}
}

//通用
func NewException(code int, message string) Exception {
	if _, ok := codeMap[code]; !ok {
		//如果code不存在定义的map中
		return &exception{
			httpCode: 498, //乱定义了一个http code
			code:     code,
			message:  message,
			explain:  "未知的错误",
		}
	}
	return &exception{
		httpCode: codeMap[code].httpCode,
		code:     code,
		message:  message,
		explain:  codeMap[code].explain,
	}
}

//根据错误返回预先定义的错误
func AutoException(err error) Exception{
	switch true {
	case errors.Is(err, sql.ErrNoRows):
		return &exception{
			httpCode: 500,
			code:     1501,
			message:  "数据库记录不存在",
			explain:  err.Error(),
		}
	case errors.Is(err, sql.ErrConnDone):
		return &exception{
			httpCode: 500,
			code:     1502,
			message:  "数据库连接已关闭",
			explain:  err.Error(),
		}
	case errors.Is(err, sql.ErrTxDone):
		return &exception{
			httpCode: 500,
			code:     1503,
			message:  "数据库事务提交失败",
			explain:  err.Error(),
		}
	default:
		return &exception{
			httpCode: 500,
			code:     1500,
			message:  "服务器内部错误",
			explain:  err.Error(),
		}
	}
}
