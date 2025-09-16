package response

import (
	"time"

	"github.com/fastgox/fastgox-api-starter/src/models/entity"
)

// LoginWithSmsResult 短信登录结果
type LoginWithSmsResult struct {
	Success   bool         `json:"success"`              // 是否成功
	Message   string       `json:"message"`              // 消息
	User      *entity.User `json:"user,omitempty"`       // 用户信息
	Token     string       `json:"token,omitempty"`      // JWT令牌
	IsNewUser bool         `json:"is_new_user"`          // 是否为新用户
	ExpiresAt *time.Time   `json:"expires_at,omitempty"` // 令牌过期时间
}

// SendLoginSmsResult 发送登录短信结果
type SendLoginSmsResult struct {
	Success    bool       `json:"success"`                // 是否成功
	Message    string     `json:"message"`                // 消息
	ExpireAt   *time.Time `json:"expire_at,omitempty"`    // 过期时间
	NextSendAt *time.Time `json:"next_send_at,omitempty"` // 下次可发送时间
}
