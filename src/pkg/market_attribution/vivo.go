package market_attribution

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/fastgox/fastgox-api-starter/src/core/config"
	"github.com/fastgox/fastgox-api-starter/src/pkg"
	"github.com/fastgox/utils/http"
)

// VivoProvider vivo渠道归因服务提供商
type VivoProvider struct {
	APIKey    string
	SecretKey string
	BaseURL   string
}

// VivoActivationRequest vivo激活上报请求
type VivoActivationRequest struct {
	IMEI       string  `json:"imei"`        // 设备IMEI
	OAID       string  `json:"oaid"`        // 设备OAID
	ClientIP   string  `json:"client_ip"`   // 客户端IP
	UA         string  `json:"ua"`          // 用户代理
	ConvTime   int64   `json:"conv_time"`   // 转化时间
	ConvWeight float64 `json:"conv_weight"` // 转化权重
	RequestId  string  `json:"request_id"`  // 请求ID
	CreativeId string  `json:"creative_id"` // 创意ID
}

// VivoRegistrationRequest vivo注册上报请求
type VivoRegistrationRequest struct {
	IMEI       string  `json:"imei"`        // 设备IMEI
	OAID       string  `json:"oaid"`        // 设备OAID
	ClientIP   string  `json:"client_ip"`   // 客户端IP
	UA         string  `json:"ua"`          // 用户代理
	ConvTime   int64   `json:"conv_time"`   // 转化时间
	ConvWeight float64 `json:"conv_weight"` // 转化权重
	RequestId  string  `json:"request_id"`  // 请求ID
	CreativeId string  `json:"creative_id"` // 创意ID
}

// VivoResponse vivo API响应
type VivoResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		RequestId string `json:"request_id"`
	} `json:"data"`
}

// NewVivoProvider 创建vivo归因服务提供商
func NewVivoProvider() *VivoProvider {
	// 如果配置未初始化，返回默认配置
	if config.GlobalConfig == nil {
		return &VivoProvider{
			APIKey:    "default-vivo-api-key",
			SecretKey: "default-vivo-secret-key",
			BaseURL:   "https://api.vivo.com/attribution",
		}
	}

	return &VivoProvider{
		APIKey:    config.GlobalConfig.AppMarket.Vivo.APIKey,
		SecretKey: config.GlobalConfig.AppMarket.Vivo.SecretKey,
		BaseURL:   config.GlobalConfig.AppMarket.Vivo.BaseURL,
	}
}

// GetName 获取提供商名称
func (v *VivoProvider) GetName() string {
	return ProviderNameVivo
}

// Call 类型安全的调用方法
func (v *VivoProvider) Call(input *EventInput) (*MarketAttributionOutput, error) {
	return v.ReportEvent(input)
}

// ReportEvent 上报事件数据（统一处理激活和注册）
func (v *VivoProvider) ReportEvent(input *EventInput) (*MarketAttributionOutput, error) {
	if input == nil {
		return nil, fmt.Errorf("事件上报输入参数不能为空")
	}

	// 设置默认值
	if input.ConvTime == 0 {
		input.ConvTime = time.Now().UnixMilli()
	}
	if input.ConvWeight == 0 {
		input.ConvWeight = 1.0
	}
	if input.RequestID == "" {
		input.RequestID = "default_request_id"
	}
	if input.CreativeId == "" {
		input.CreativeId = "default_creative_id"
	}

	var url string
	var requestData map[string]interface{}

	switch input.EventType {
	case EventTypeActivation:
		// 构建vivo激活请求
		request := &VivoActivationRequest{
			IMEI:       input.IMEI,
			OAID:       input.OAID,
			ClientIP:   input.ClientIP,
			UA:         input.UA,
			ConvTime:   input.ConvTime,
			ConvWeight: input.ConvWeight,
			RequestId:  input.RequestID,
			CreativeId: input.CreativeId,
		}
		url = fmt.Sprintf("%s/activation", v.BaseURL)
		requestData = v.buildRequestData(request)

	case EventTypeRegistration:
		// 构建vivo注册请求
		request := &VivoRegistrationRequest{
			IMEI:       input.IMEI,
			OAID:       input.OAID,
			ClientIP:   input.ClientIP,
			UA:         input.UA,
			ConvTime:   input.ConvTime,
			ConvWeight: input.ConvWeight,
			RequestId:  input.RequestID,
			CreativeId: input.CreativeId,
		}
		url = fmt.Sprintf("%s/registration", v.BaseURL)
		requestData = v.buildRequestData(request)

	default:
		return nil, fmt.Errorf("不支持的事件类型: %s", input.EventType)
	}

	// 发送HTTP请求
	resp := http.Post[VivoResponse](url, requestData)
	if resp.Error != nil {
		return nil, fmt.Errorf("vivo%s上报请求失败: %w", input.EventType, resp.Error)
	}

	body := resp.Body
	return &MarketAttributionOutput{
		BaseOutput: pkg.BaseOutput{
			Success:   body.Success,
			Code:      fmt.Sprintf("%d", body.Code),
			Message:   body.Message,
			RequestID: body.Data.RequestId,
		},
		Platform: "vivo",
	}, nil
}

// buildRequestData 构建请求数据，添加认证信息
func (v *VivoProvider) buildRequestData(data interface{}) map[string]interface{} {
	// 将结构体转换为map
	jsonData, _ := json.Marshal(data)
	var requestData map[string]interface{}
	json.Unmarshal(jsonData, &requestData)

	// 添加认证信息
	requestData["api_key"] = v.APIKey
	requestData["secret_key"] = v.SecretKey
	requestData["timestamp"] = time.Now().Unix()

	return requestData
}

// 在包初始化时注册vivo提供商
func init() {
	provider := NewVivoProvider()
	AttributionManager.Register(provider.GetName(), provider)
}
