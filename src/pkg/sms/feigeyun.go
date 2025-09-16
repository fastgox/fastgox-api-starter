package sms

import (
	"fmt"

	"github.com/fastgox/fastgox-api-starter/src/core/config"
	"github.com/fastgox/fastgox-api-starter/src/pkg"
	"github.com/fastgox/utils/http"
)

// FeigeYunProvider 飞鸽云短信服务提供商
type FeigeYunProvider struct {
}

// GetName 获取提供商名称
func (f *FeigeYunProvider) GetName() string {
	return ProviderNameFeigeyun
}

// Call 统一调用方法
func (f *FeigeYunProvider) Call(input SmsInput) (*SmsOutput, error) {
	switch v := input.(type) {
	case *SendSmsInput:
		return f.SendSms(v)
	case *SendTemplateInput:
		return f.SendTemplate(v)
	default:
		return nil, fmt.Errorf("不支持的输入类型: %T", input)
	}
}

// SendSms 发送验证码短信
func (f *FeigeYunProvider) SendSms(input *SendSmsInput) (*SmsOutput, error) {
	content := fmt.Sprintf("您的验证码为%s，正在登录闪呗APP，若非本人操作，请忽略本短信。", input.Code)
	requestData := map[string]interface{}{
		"apikey":      config.GlobalConfig.SMS.FeigeYun.APIKey,
		"secret":      config.GlobalConfig.SMS.FeigeYun.Secret,
		"mobile":      input.Phone,
		"sign_id":     config.GlobalConfig.SMS.FeigeYun.SignID,
		"template_id": config.GlobalConfig.SMS.FeigeYun.TemplateID,
		"content":     content,
	}
	resp := http.Post[FeigeYunResponse](config.GlobalConfig.SMS.FeigeYun.APIURL, requestData)
	if resp.Error != nil {
		return nil, resp.Error
	}
	body := resp.Body
	if resp.Body.Code == 0 {
		body.Success = true
	}
	return &SmsOutput{
		BaseOutput: pkg.BaseOutput{
			Success: body.Success,
			Code:    fmt.Sprintf("%d", body.Code),
			Message: body.Message,
		},
		MessageID: body.MessageID,
		Fee:       body.Fee,
		Count:     body.Count,
	}, nil
}

// SendTemplate 发送模板短信
func (f *FeigeYunProvider) SendTemplate(input *SendTemplateInput) (*SmsOutput, error) {
	content := input.Params["content"]
	if content == "" {
		return nil, fmt.Errorf("模板短信内容不能为空")
	}

	requestData := map[string]interface{}{
		"apikey":      config.GlobalConfig.SMS.FeigeYun.APIKey,
		"secret":      config.GlobalConfig.SMS.FeigeYun.Secret,
		"mobile":      input.Phone,
		"sign_id":     config.GlobalConfig.SMS.FeigeYun.SignID,
		"template_id": config.GlobalConfig.SMS.FeigeYun.TemplateID,
		"content":     content,
	}
	resp := http.Post[FeigeYunResponse](config.GlobalConfig.SMS.FeigeYun.APIURL, requestData)
	if resp.Error != nil {
		return nil, resp.Error
	}
	body := resp.Body
	return &SmsOutput{
		BaseOutput: pkg.BaseOutput{
			Success: body.Success,
			Code:    fmt.Sprintf("%d", body.Code),
			Message: body.Message,
		},
		MessageID: body.MessageID,
		Fee:       body.Fee,
		Count:     body.Count,
	}, nil
}

// FeigeYunResponse 飞鸽云SMS API响应
type FeigeYunResponse struct {
	Success   bool    `json:"success"`
	MessageID string  `json:"message_id"`
	Code      int     `json:"code"`
	Message   string  `json:"message"`
	Fee       float64 `json:"fee"`
	Count     int     `json:"count"`
}

func init() {
	provider := &FeigeYunProvider{}
	SmsManager.Register(provider.GetName(), provider)
}
