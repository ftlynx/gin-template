package exception

const (
	CodeOk = 0

	CodeBadRequest       = 400
	CodeNotAuthorization = 401
	CodeForbidden        = 403
)

type codeMapping struct {
	httpCode int
	explain  string
}

var codeMap = map[int]codeMapping{
	CodeOk:               {200, "成功"},
	CodeBadRequest:       {400, "参数错误"},
	CodeForbidden:        {403, "权限拒绝"},
	CodeNotAuthorization: {401, "用户未登陆"},
}
