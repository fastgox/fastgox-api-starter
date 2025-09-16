package geolocation

import (
	"github.com/fastgox/fastgox-api-starter/src/pkg"
)

// 提供商名称常量
const (
	ProviderNameAmap = "amap" // 高德地图提供商
	// 可以继续添加其他提供商...
	// ProviderNameBaidu = "baidu" // 百度地图提供商
	// ProviderNameTencent = "tencent" // 腾讯地图提供商
)

// 全局地理位置服务提供商管理器
var GeolocationManager = pkg.NewManager[pkg.Provider[GeolocationInput, *GeolocationOutput]](ProviderNameAmap)

// GeolocationInput 地理位置输入联合类型
type GeolocationInput interface {
	isGeolocationInput()
}

// IPLocationInput IP地理位置查询输入
type IPLocationInput struct {
	pkg.BaseInput
	IP string `json:"ip"` // IP地址
}

// isGeolocationInput 实现 GeolocationInput 接口
func (i *IPLocationInput) isGeolocationInput() {}

// GeolocationOutput 地理位置查询响应
type GeolocationOutput struct {
	pkg.BaseOutput
	IP        string  `json:"ip"`        // IP地址
	City      string  `json:"city"`      // 城市
	Region    string  `json:"region"`    // 省份/地区
	Country   string  `json:"country"`   // 国家
	Latitude  float64 `json:"latitude"`  // 纬度
	Longitude float64 `json:"longitude"` // 经度
}
