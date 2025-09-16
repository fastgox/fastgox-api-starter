package sms

import (
	"fmt"

	"github.com/fastgox/fastgox-api-starter/src/pkg"
)

// AliyunSmsProvider 阿里云短信服务提供商
type AliyunSmsProvider struct {
}

// GetName 获取提供商名称
func (a *AliyunSmsProvider) GetName() string {
	return ProviderNameAliyun
}

// Call 统一调用方法
func (a *AliyunSmsProvider) Call(input SmsInput) (*SmsOutput, error) {
	switch v := input.(type) {
	case *SendSmsInput:
		return a.SendSms(v)
	case *SendTemplateInput:
		return a.SendTemplate(v)
	default:
		return nil, fmt.Errorf("不支持的输入类型: %T", input)
	}
}

// SendSms 发送验证码短信
func (a *AliyunSmsProvider) SendSms(input *SendSmsInput) (*SmsOutput, error) {
	// 这里是阿里云短信的具体实现
	// 为了演示，这里只是返回一个模拟结果
	return &SmsOutput{
		BaseOutput: pkg.BaseOutput{
			Success:  true,
			Code:     "200",
			Message:  "发送成功",
			Provider: a.GetName(),
		},
		MessageID: "aliyun-msg-" + input.Phone,
		Fee:       0.05,
		Count:     1,
	}, nil
}

// SendTemplate 发送模板短信
func (a *AliyunSmsProvider) SendTemplate(input *SendTemplateInput) (*SmsOutput, error) {
	// 这里是阿里云模板短信的具体实现
	return &SmsOutput{
		BaseOutput: pkg.BaseOutput{
			Success:  true,
			Code:     "200",
			Message:  "模板短信发送成功",
			Provider: a.GetName(),
		},
		MessageID: "aliyun-template-" + input.Phone,
		Fee:       0.05,
		Count:     1,
	}, nil
}

// 自动注册到管理器
func init() {
	provider := &AliyunSmsProvider{}
	SmsManager.Register(provider.GetName(), provider)
}
