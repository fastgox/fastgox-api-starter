package response

import "time"

// SendSmsCodeResult 发送验证码结果
type SendSmsCodeResult struct {
	Success    bool       `json:"success"`                // 是否成功
	Message    string     `json:"message"`                // 消息
	Code       string     `json:"code,omitempty"`         // 验证码（测试环境可返回）
	ExpireAt   *time.Time `json:"expire_at,omitempty"`    // 过期时间
	NextSendAt *time.Time `json:"next_send_at,omitempty"` // 下次可发送时间
}

// VerifySmsCodeResult 验证验证码结果
type VerifySmsCodeResult struct {
	Success bool   `json:"success"` // 是否成功
	Message string `json:"message"` // 消息
}
