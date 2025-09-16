package auth

import (
	"encoding/json"
	"fmt"

	"github.com/fastgox/fastgox-api-starter/src/core/config"
	"github.com/fastgox/fastgox-api-starter/src/pkg"
	"github.com/fastgox/fastgox-api-starter/src/pkg/transport/qiandun"
)

// QiandunAuthProvider 钱盾身份认证服务提供商
type QiandunAuthProvider struct {
	client *qiandun.Client
}

// NewQiandunAuthProvider 创建钱盾身份认证服务提供商
func NewQiandunAuthProvider() *QiandunAuthProvider {
	return &QiandunAuthProvider{}
}

// 延迟初始化客户端
func (q *QiandunAuthProvider) ensureClient() error {
	if q.client != nil {
		return nil
	}

	if config.GlobalConfig == nil {
		return fmt.Errorf("配置未初始化")
	}

	cfg := config.GlobalConfig.Auth.Qiandun
	transportConfig := &qiandun.Config{
		AppKey:    cfg.AppKey,
		AppSecret: cfg.AppSecret,
		BaseURL:   "", // auth配置中没有BaseURL，使用默认值
		IsProd:    cfg.IsProd,
	}

	q.client = qiandun.NewClient(transportConfig)
	return nil
}

// GetName 获取提供商名称
func (q *QiandunAuthProvider) GetName() string {
	return ProviderNameQiandun
}

// Call 类型安全的调用方法
func (q *QiandunAuthProvider) Call(input *ThreeElementsInput) (*AuthOutput, error) {
	return q.VerifyThreeElements(input)
}

// VerifyThreeElements 三要素验证（姓名、身份证、手机号）
func (q *QiandunAuthProvider) VerifyThreeElements(input *ThreeElementsInput) (*AuthOutput, error) {
	// 确保客户端已初始化
	if err := q.ensureClient(); err != nil {
		return nil, err
	}

	// 构建钱盾API请求格式
	qiandunReq := QiandunThreeElementsRequest{
		PsnName:       input.Name,
		PsnIDCardType: "CRED_PSN_CH_IDCARD", // 固定值：中国身份证
		PsnIDCardNum:  input.IdCard,
		PsnMobile:     input.Mobile,
	}

	// 使用transport层发送请求
	resp, err := q.client.SendRequest("/open/psnCert/3element", qiandunReq)
	if err != nil {
		return nil, err
	}

	// 检查业务状态
	if !resp.Success || resp.Code != "200" {
		return nil, fmt.Errorf("千盾认证失败: %s", resp.Msg)
	}

	// 解析具体的认证结果
	var authResult QiandunResult
	if resultBytes, err := json.Marshal(resp.Result); err != nil {
		return nil, fmt.Errorf("解析认证结果失败: %v", err)
	} else if err := json.Unmarshal(resultBytes, &authResult); err != nil {
		return nil, fmt.Errorf("解析认证结果失败: %v", err)
	}

	// 构造统一响应格式
	result := &AuthOutput{
		BaseOutput: pkg.BaseOutput{
			Success:  resp.Success,
			Message:  resp.Msg,
			Code:     resp.Code,
			Provider: "qiandun",
			CostTime: resp.CostTime,
		},
		VerifyStatus: authResult.VerifyStatus,
		FlowNo:       authResult.FlowNo,
		ServiceId:    authResult.ServiceId,
		LogKey:       resp.LogKey,
	}

	return result, nil
}

// QiandunThreeElementsRequest 钱盾三要素验证请求
type QiandunThreeElementsRequest struct {
	PsnName       string `json:"psnName"`       // 姓名
	PsnIDCardType string `json:"psnIDCardType"` // 证件类型，固定值：CRED_PSN_CH_IDCARD
	PsnIDCardNum  string `json:"psnIDCardNum"`  // 证件号码
	PsnMobile     string `json:"psnMobile"`     // 手机号
}

// QiandunResult 钱盾认证验证结果
type QiandunResult struct {
	VerifyStatus string `json:"verifyStatus"` // passed-通过，un_passed-未通过
	FlowNo       string `json:"flowNo"`       // 验证流水号
	ServiceId    string `json:"serviceId"`    // 本次服务流程ID
}

func init() {
	// 注册钱盾身份认证服务提供商
	provider := NewQiandunAuthProvider()
	AuthManager.Register(provider.GetName(), provider)
}
