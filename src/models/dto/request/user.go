package request

// SendLoginSmsRequest 发送登录短信请求
type SendLoginSmsRequest struct {
	Phone string `json:"phone" binding:"required"` // 手机号
}

// LoginWithSmsRequest 短信登录请求
type LoginWithSmsRequest struct {
	Phone string `json:"phone" binding:"required"` // 手机号
	Code  string `json:"code" binding:"required"`  // 验证码
}
