package market_attribution

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/fastgox/fastgox-api-starter/src/core/config"
	"github.com/fastgox/fastgox-api-starter/src/pkg"
	"github.com/fastgox/utils/http"
)

// XiaomiProvider 小米渠道归因服务提供商
type XiaomiProvider struct {
	APIKey    string
	SecretKey string
	BaseURL   string
}

// XiaomiActivationRequest 小米激活上报请求
type XiaomiActivationRequest struct {
	ClientIP   string  `json:"client_ip"`   // 客户端IP
	UA         string  `json:"ua"`          // 用户代理(设备信息)
	ConvTime   int64   `json:"conv_time"`   // 转化时间(毫秒时间戳)
	OAID       string  `json:"oaid"`        // 设备OAID
	ConvWeight float64 `json:"conv_weight"` // 转化权重(可选)
	AdId       int64   `json:"ad_id"`       // 广告ID(可选)
}

// XiaomiRegistrationRequest 小米注册上报请求
type XiaomiRegistrationRequest struct {
	ClientIP   string  `json:"client_ip"`   // 客户端IP
	UA         string  `json:"ua"`          // 用户代理
	ConvTime   int64   `json:"conv_time"`   // 转化时间(毫秒时间戳)
	OAID       string  `json:"oaid"`        // 设备OAID
	ConvWeight float64 `json:"conv_weight"` // 转化权重(可选)
	AdId       int64   `json:"ad_id"`       // 广告ID(可选)
}

// XiaomiResponse 小米API响应
type XiaomiResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		RequestId string `json:"request_id"`
	} `json:"data"`
}

// NewXiaomiProvider 创建小米归因服务提供商
func NewXiaomiProvider() *XiaomiProvider {
	// 如果配置未初始化，返回默认配置
	if config.GlobalConfig == nil {
		return &XiaomiProvider{
			APIKey:    "default-xiaomi-api-key",
			SecretKey: "default-xiaomi-secret-key",
			BaseURL:   "https://api.xiaomi.com/attribution",
		}
	}

	return &XiaomiProvider{
		APIKey:    config.GlobalConfig.AppMarket.Xiaomi.APIKey,
		SecretKey: config.GlobalConfig.AppMarket.Xiaomi.SecretKey,
		BaseURL:   config.GlobalConfig.AppMarket.Xiaomi.BaseURL,
	}
}

// GetName 获取提供商名称
func (x *XiaomiProvider) GetName() string {
	return ProviderNameXiaomi
}

// Call 类型安全的调用方法
func (x *XiaomiProvider) Call(input *EventInput) (*MarketAttributionOutput, error) {
	return x.ReportEvent(input)
}

// ReportEvent 上报事件数据（统一处理激活和注册）
func (x *XiaomiProvider) ReportEvent(input *EventInput) (*MarketAttributionOutput, error) {
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

	var url string
	var requestData map[string]interface{}

	switch input.EventType {
	case EventTypeActivation:
		// 构建小米激活请求
		request := &XiaomiActivationRequest{
			ClientIP:   input.ClientIP,
			UA:         input.UA,
			ConvTime:   input.ConvTime,
			OAID:       input.OAID,
			ConvWeight: input.ConvWeight,
			AdId:       input.AdId,
		}
		url = fmt.Sprintf("%s/activation", x.BaseURL)
		requestData = x.buildRequestData(request)

	case EventTypeRegistration:
		// 构建小米注册请求
		request := &XiaomiRegistrationRequest{
			ClientIP:   input.ClientIP,
			UA:         input.UA,
			ConvTime:   input.ConvTime,
			OAID:       input.OAID,
			ConvWeight: input.ConvWeight,
			AdId:       input.AdId,
		}
		url = fmt.Sprintf("%s/registration", x.BaseURL)
		requestData = x.buildRequestData(request)

	default:
		return nil, fmt.Errorf("不支持的事件类型: %s", input.EventType)
	}

	// 发送HTTP请求
	resp := http.Post[XiaomiResponse](url, requestData)
	if resp.Error != nil {
		return nil, fmt.Errorf("小米%s上报请求失败: %w", input.EventType, resp.Error)
	}

	body := resp.Body
	return &MarketAttributionOutput{
		BaseOutput: pkg.BaseOutput{
			Success:   body.Success,
			Code:      fmt.Sprintf("%d", body.Code),
			Message:   body.Message,
			RequestID: body.Data.RequestId,
		},
		Platform: "xiaomi",
	}, nil
}

// buildRequestData 构建请求数据，添加认证信息
func (x *XiaomiProvider) buildRequestData(data interface{}) map[string]interface{} {
	// 将结构体转换为map
	jsonData, _ := json.Marshal(data)
	var requestData map[string]interface{}
	json.Unmarshal(jsonData, &requestData)

	// 添加认证信息
	requestData["api_key"] = x.APIKey
	requestData["secret_key"] = x.SecretKey
	requestData["timestamp"] = time.Now().Unix()

	return requestData
}

// 在包初始化时注册小米提供商
func init() {
	provider := NewXiaomiProvider()
	AttributionManager.Register(provider.GetName(), provider)
}
