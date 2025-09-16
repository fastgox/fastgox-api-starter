package response

import "time"

// AdConversionResponse 广告转化记录响应
type AdConversionResponse struct {
	ID                 int64      `json:"id"`                        // 主键ID
	AdID               string     `json:"ad_id"`                     // 广告ID
	ChannelCode        string     `json:"channel_code"`              // 渠道代码
	ChannelName        string     `json:"channel_name"`              // 渠道名称
	UserID             *int64     `json:"user_id,omitempty"`         // 用户ID
	ConversionType     int8       `json:"conversion_type"`           // 转化类型
	ConversionTypeName string     `json:"conversion_type_name"`      // 转化类型名称
	IsRepeatConvert    bool       `json:"is_repeat_convert"`         // 是否重复转化
	Platform           *string    `json:"platform,omitempty"`        // 平台来源
	DeviceID           *string    `json:"device_id,omitempty"`       // 设备ID
	IP                 *string    `json:"ip,omitempty"`              // IP地址
	ConversionTime     *time.Time `json:"conversion_time,omitempty"` // 转化时间
	Source             *string    `json:"source,omitempty"`          // 流量来源
	Medium             *string    `json:"medium,omitempty"`          // 媒介类型
	Campaign           *string    `json:"campaign,omitempty"`        // 活动名称
	CreatedAt          *time.Time `json:"created_at,omitempty"`      // 创建时间
	
	// 设备历史提交状态：该设备是否曾经提交过任何转化记录（用于智能重复安装处理）
	DeviceHasHistory   bool       `json:"device_has_history"`        // 设备是否有历史提交记录
}
