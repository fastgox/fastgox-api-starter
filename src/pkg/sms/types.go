package sms

import (
	"github.com/fastgox/fastgox-api-starter/src/pkg"
)

// 提供商名称常量
const (
	ProviderNameFeigeyun = "feigeyun" // 飞鸽云提供商
	ProviderNameAliyun   = "aliyun"   // 阿里云短信提供商
	// 可以继续添加其他提供商...
)

// 全局SMS服务提供商管理器
var SmsManager = pkg.NewManager[pkg.Provider[SmsInput, *SmsOutput]](ProviderNameFeigeyun)

// SmsInput SMS输入联合类型
type SmsInput interface {
	isSmsInput()
}

// SendSmsInput 发送短信输入
type SendSmsInput struct {
	pkg.BaseInput
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

// isSmsInput 实现 SmsInput 接口
func (s *SendSmsInput) isSmsInput() {}

// SendTemplateInput 发送模板短信输入
type SendTemplateInput struct {
	pkg.BaseInput
	Phone      string            `json:"phone"`
	TemplateID string            `json:"template_id"`
	Params     map[string]string `json:"params"`
}

// isSmsInput 实现 SmsInput 接口
func (s *SendTemplateInput) isSmsInput() {}

// SmsOutput 短信发送响应
type SmsOutput struct {
	pkg.BaseOutput
	MessageID string  `json:"message_id"` // 消息ID
	Fee       float64 `json:"fee"`        // 费用
	Count     int     `json:"count"`      // 发送数量
}
