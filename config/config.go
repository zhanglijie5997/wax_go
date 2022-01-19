package config

import "github.com/gin-gonic/gin"

const (
	EmailIsNotValidated = 10001	// 邮箱错误
	PasswordIsNotValidate = 10002 // 密码错误
	IsNotRegister = 10003	// 未注册
	IsRegistered = 10004	// email 注册过
	EmailIsRegister = 10005 // email注册过
	RegisterFailed = 10005 // 注册错误
	NeedLogin = 10006 // 需要登录
	LoginFailed = 10007 // 登录失败
	TokenExpired = 10008 // token过期
	RequestError = 10009 // 请求错误
	AccountIsNotHave = 10010 // 账号未关联
	Success = 20000
)

type Status struct {
	Code int `json:"json"`
	Msg string `json:"msg"`
}

// Status200 200 - 请求成功
func Status200() int {
	return 200
}

// Status400 400 - 请求格式错误
func Status400() Status {
	return Status{Code: 400, Msg: "Invalid format"}
}

// Status_404 404 - 请求路径不存在
func Status_404() Status {
	return Status{Code: 404, Msg: "Path does not exist"}
}

// ErrMethods  无此方法
func ErrMethods(c *gin.Context)  {
	c.Status(405)
}