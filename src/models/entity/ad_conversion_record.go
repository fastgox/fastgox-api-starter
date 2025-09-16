package entity

import "time"

// 媒介类型枚举
const (
	MediumCPC       = "cpc"       // 按点击付费广告
	MediumCPM       = "cpm"       // 按展示付费广告
	MediumDisplay   = "display"   // 展示广告
	MediumSocial    = "social"    // 社交媒体
	MediumSearch    = "search"    // 搜索引擎
	MediumEmail     = "email"     // 邮件营销
	MediumReferral  = "referral"  // 推荐流量
	MediumOrganic   = "organic"   // 自然流量
	MediumDirect    = "direct"    // 直接访问
	MediumVideo     = "video"     // 视频广告
	MediumBanner    = "banner"    // 横幅广告
	MediumAffiliate = "affiliate" // 联盟营销
)

// AdConversionRecord 广告转化记录实体
// 对应数据库表 t_ad_conversion_record
// @Description 广告转化记录信息
type AdConversionRecord struct {
	ID              int64      `gorm:"column:id;primaryKey;autoIncrement;comment:主键ID" json:"id"`
	CreatedAt       *time.Time `gorm:"column:created_at;comment:创建时间" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`
	AdID            string     `gorm:"column:ad_id;size:255;not null;comment:广告ID" json:"ad_id"`
	ChannelCode     string     `gorm:"column:channel_code;size:100;not null;comment:渠道代码" json:"channel_code"`
	UserID          *int64     `gorm:"column:user_id;comment:用户ID，关联t_user表" json:"user_id,omitempty"`
	ConversionType  int8       `gorm:"column:conversion_type;not null;comment:转化类型:1=首次打开app,2=注册,3=实名,4=留资" json:"conversion_type"`
	IsRepeatConvert bool       `gorm:"column:is_repeat_convert;default:0;comment:是否重复转化" json:"is_repeat_convert"`
	Platform        *string    `gorm:"column:platform;size:50;comment:平台来源(ios/android)" json:"platform,omitempty"`
	DeviceID        *string    `gorm:"column:device_id;size:255;comment:设备唯一标识" json:"device_id,omitempty"`
	IP              *string    `gorm:"column:ip;size:80;comment:用户IP地址" json:"ip,omitempty"`
	ConversionTime  *time.Time `gorm:"column:conversion_time;comment:转化时间" json:"conversion_time,omitempty"`
	Medium          *string    `gorm:"column:medium;size:100;comment:媒介类型" json:"medium,omitempty"`
}

// TableName 指定表名
func (AdConversionRecord) TableName() string {
	return "t_ad_conversion_record"
}

// IsValidConversionType 检查转化类型是否有效
func (a *AdConversionRecord) IsValidConversionType() bool {
	return a.ConversionType >= 1 && a.ConversionType <= 4
}

// GetConversionTypeName 获取转化类型名称
func (a *AdConversionRecord) GetConversionTypeName() string {
	switch a.ConversionType {
	case 1:
		return "首次打开app"
	case 2:
		return "注册"
	case 3:
		return "实名认证"
	case 4:
		return "留资"
	default:
		return "未知"
	}
}

// GetChannelName 获取渠道名称（直接返回渠道代码）
func (a *AdConversionRecord) GetChannelName() string {
	return a.ChannelCode
}

// IsValidMediumType 检查媒介类型是否有效
func (a *AdConversionRecord) IsValidMediumType() bool {
	if a.Medium == nil {
		return true // 允许为空
	}

	validMediums := []string{
		MediumCPC,
		MediumCPM,
		MediumDisplay,
		MediumSocial,
		MediumSearch,
		MediumEmail,
		MediumReferral,
		MediumOrganic,
		MediumDirect,
		MediumVideo,
		MediumBanner,
		MediumAffiliate,
	}

	for _, validMedium := range validMediums {
		if *a.Medium == validMedium {
			return true
		}
	}
	return false
}

// GetMediumTypeName 获取媒介类型中文名称
func (a *AdConversionRecord) GetMediumTypeName() string {
	if a.Medium == nil {
		return "未知"
	}

	mediumMap := map[string]string{
		MediumCPC:       "按点击付费",
		MediumCPM:       "按展示付费",
		MediumDisplay:   "展示广告",
		MediumSocial:    "社交媒体",
		MediumSearch:    "搜索引擎",
		MediumEmail:     "邮件营销",
		MediumReferral:  "推荐流量",
		MediumOrganic:   "自然流量",
		MediumDirect:    "直接访问",
		MediumVideo:     "视频广告",
		MediumBanner:    "横幅广告",
		MediumAffiliate: "联盟营销",
	}

	if name, exists := mediumMap[*a.Medium]; exists {
		return name
	}
	return *a.Medium
}
