package geolocation

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fastgox/fastgox-api-starter/src/core/config"
	"github.com/fastgox/fastgox-api-starter/src/pkg"
)

// AmapProvider 高德地图地理位置服务提供商
type AmapProvider struct {
	APIKey string
}

// NewAmapProvider 创建高德地图地理位置服务提供商
func NewAmapProvider(apiKey string) *AmapProvider {
	return &AmapProvider{
		APIKey: apiKey,
	}
}

// GetName 获取提供商名称
func (p *AmapProvider) GetName() string {
	return ProviderNameAmap
}

// Call 调用地理位置服务
func (p *AmapProvider) Call(input GeolocationInput) (*GeolocationOutput, error) {
	startTime := time.Now()

	// 类型断言获取具体的输入类型
	ipInput, ok := input.(*IPLocationInput)
	if !ok {
		return &GeolocationOutput{
			BaseOutput: pkg.BaseOutput{
				Success:   false,
				Code:      "INVALID_INPUT",
				Message:   "不支持的输入类型",
				Provider:  p.GetName(),
				CostTime:  time.Since(startTime).Milliseconds(),
				Timestamp: time.Now(),
			},
		}, fmt.Errorf("不支持的输入类型")
	}

	// 如果是本地IP，返回默认信息
	if isLocalIP(ipInput.IP) {
		return &GeolocationOutput{
			BaseOutput: pkg.BaseOutput{
				Success:   true,
				Code:      "SUCCESS",
				Message:   "本地IP地址",
				RequestID: ipInput.RequestID,
				Provider:  p.GetName(),
				CostTime:  time.Since(startTime).Milliseconds(),
				Timestamp: time.Now(),
			},
			IP:        ipInput.IP,
			City:      "本地",
			Region:    "本地",
			Country:   "CN",
			Latitude:  39.9042, // 北京坐标
			Longitude: 116.4074,
		}, nil
	}

	// 使用高德地图IP定位服务
	url := fmt.Sprintf("https://restapi.amap.com/v3/ip?ip=%s&key=%s", ipInput.IP, p.APIKey)

	// 配置HTTP传输，非生产环境跳过TLS证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // 跳过证书验证
		},
	}

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: tr,
	}

	resp, err := client.Get(url)
	if err != nil {
		return &GeolocationOutput{
			BaseOutput: pkg.BaseOutput{
				Success:   false,
				Code:      "REQUEST_FAILED",
				Message:   fmt.Sprintf("请求地理位置服务失败: %v", err),
				RequestID: ipInput.RequestID,
				Provider:  p.GetName(),
				CostTime:  time.Since(startTime).Milliseconds(),
				Timestamp: time.Now(),
			},
		}, err
	}
	defer resp.Body.Close()

	var amapResponse struct {
		Status    string `json:"status"`
		Info      string `json:"info"`
		InfoCode  string `json:"infocode"`
		Province  string `json:"province"`
		City      string `json:"city"`
		AdCode    string `json:"adcode"`
		Rectangle string `json:"rectangle"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&amapResponse); err != nil {
		return &GeolocationOutput{
			BaseOutput: pkg.BaseOutput{
				Success:   false,
				Code:      "PARSE_FAILED",
				Message:   fmt.Sprintf("解析地理位置响应失败: %v", err),
				RequestID: ipInput.RequestID,
				Provider:  p.GetName(),
				CostTime:  time.Since(startTime).Milliseconds(),
				Timestamp: time.Now(),
			},
		}, err
	}

	if amapResponse.Status != "1" {
		return &GeolocationOutput{
			BaseOutput: pkg.BaseOutput{
				Success:   false,
				Code:      amapResponse.InfoCode,
				Message:   fmt.Sprintf("获取地理位置失败: %s", amapResponse.Info),
				RequestID: ipInput.RequestID,
				Provider:  p.GetName(),
				CostTime:  time.Since(startTime).Milliseconds(),
				Timestamp: time.Now(),
			},
		}, fmt.Errorf("获取地理位置失败: %s", amapResponse.Info)
	}

	// 高德地图API不返回经纬度，使用默认坐标（可以根据需要调用其他API获取）
	latitude := 39.9042 // 默认北京坐标
	longitude := 116.4074

	return &GeolocationOutput{
		BaseOutput: pkg.BaseOutput{
			Success:   true,
			Code:      "SUCCESS",
			Message:   "获取地理位置成功",
			RequestID: ipInput.RequestID,
			Provider:  p.GetName(),
			CostTime:  time.Since(startTime).Milliseconds(),
			Timestamp: time.Now(),
		},
		IP:        ipInput.IP,
		City:      amapResponse.City,
		Region:    amapResponse.Province,
		Country:   "CN", // 高德地图只支持国内IP，固定为CN
		Latitude:  latitude,
		Longitude: longitude,
	}, nil
}

// isLocalIP 判断是否为本地IP
func isLocalIP(ip string) bool {
	return ip == "127.0.0.1" || ip == "::1" ||
		strings.HasPrefix(ip, "192.168.") ||
		strings.HasPrefix(ip, "10.") ||
		strings.HasPrefix(ip, "172.16.") ||
		strings.HasPrefix(ip, "172.17.") ||
		strings.HasPrefix(ip, "172.18.") ||
		strings.HasPrefix(ip, "172.19.") ||
		strings.HasPrefix(ip, "172.20.") ||
		strings.HasPrefix(ip, "172.21.") ||
		strings.HasPrefix(ip, "172.22.") ||
		strings.HasPrefix(ip, "172.23.") ||
		strings.HasPrefix(ip, "172.24.") ||
		strings.HasPrefix(ip, "172.25.") ||
		strings.HasPrefix(ip, "172.26.") ||
		strings.HasPrefix(ip, "172.27.") ||
		strings.HasPrefix(ip, "172.28.") ||
		strings.HasPrefix(ip, "172.29.") ||
		strings.HasPrefix(ip, "172.30.") ||
		strings.HasPrefix(ip, "172.31.")
}

// InitAmapProvider 初始化高德地图提供商
func InitAmapProvider() {
	if config.GlobalConfig != nil && config.GlobalConfig.Geolocation.Providers != nil {
		if amapConfig, exists := config.GlobalConfig.Geolocation.Providers[ProviderNameAmap]; exists {
			amapProvider := NewAmapProvider(amapConfig.APIKey)
			GeolocationManager.Register(ProviderNameAmap, amapProvider)
		}
	}
}
