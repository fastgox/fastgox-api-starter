package request

// CreateAdConversionRequest 创建广告转化记录请求
type CreateAdConversionRequest struct {
	AdID           string `json:"ad_id" binding:"omitempty"`                      // 广告ID（可选，Android无真实广告ID时为空）
	ChannelCode    string `json:"channel_code" binding:"required"`                // 渠道代码
	ConversionType int8   `json:"conversion_type" binding:"required,min=1,max=4"` // 转化类型:1=首次打开app,2=注册,3=实名,4=留资
	Platform       string `json:"platform" binding:"omitempty"`                   // 平台来源 ios/android
	DeviceID       string `json:"device_id" binding:"omitempty"`                  // 设备ID
	Medium         string `json:"medium" binding:"omitempty"`                     // 媒介类型 广告/搜索
}
