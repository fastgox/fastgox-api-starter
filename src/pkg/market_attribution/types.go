package market_attribution

import (
	"github.com/fastgox/fastgox-api-starter/src/pkg"
)

// EventType 事件类型
type EventType string

const (
	EventTypeActivation   EventType = "activation"   // 激活事件
	EventTypeRegistration EventType = "registration" // 注册事件
)

// 提供商名称常量
const (
	ProviderNameVivo   = "vivo"   // vivo提供商
	ProviderNameXiaomi = "xiaomi" // 小米提供商
)

// 默认提供商
const DefaultProvider = ProviderNameVivo

// 全局应用市场归因服务提供商管理器
var AttributionManager = pkg.NewManager[pkg.Provider[*EventInput, *MarketAttributionOutput]](DefaultProvider)

// EventInput 事件上报输入参数
type EventInput struct {
	pkg.BaseInput
	EventType  EventType `json:"event_type"`  // 事件类型
	ClientIP   string    `json:"client_ip"`   // 客户端IP
	UA         string    `json:"ua"`          // 用户代理(设备信息)
	ConvTime   int64     `json:"conv_time"`   // 转化时间(毫秒时间戳)
	OAID       string    `json:"oaid"`        // 设备OAID
	IMEI       string    `json:"imei"`        // 设备IMEI
	ConvWeight float64   `json:"conv_weight"` // 转化权重(可选)
	AdId       int64     `json:"ad_id"`       // 广告ID(可选)
	CreativeId string    `json:"creative_id"` // 创意ID(vivo需要)
}

// MarketAttributionOutput 应用市场归因上报响应
type MarketAttributionOutput struct {
	pkg.BaseOutput
	Platform string `json:"platform"` // 平台名称
}

// AttributionError 归因错误
type AttributionError struct {
	Platform string
	Message  string
}

func (e *AttributionError) Error() string {
	return "归因上报错误 [" + e.Platform + "]: " + e.Message
}
