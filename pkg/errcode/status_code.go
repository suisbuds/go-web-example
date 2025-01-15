package errcode

// 将错误码映射成 HTTP 状态码
var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服务器内部错误")
	InvalidParams             = NewError(10000001, "参数错误")
	NotFound                  = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "鉴权失败, 找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "鉴权失败, Token 错误")
	UnauthorizedTokenTimeout  = NewError(10000005, "鉴权失败, Token 超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败, Token 生成失败")
	UnauthorizedTokenInvalidSigningMethod = NewError(10000007, "鉴权失败，签名方法错误")
	TooManyRequests           = NewError(10000008, "请求过多")
	EnvVarNotSet              = NewError(10000009, "环境变量未设置")
	RequestTimeout                      = NewError(10000010, "请求超时") 
)

