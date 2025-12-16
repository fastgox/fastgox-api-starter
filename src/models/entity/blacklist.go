package entity

import "time"

// BlacklistUser 黑名单用户 - 对应 t_blacklist_user 表
// 存储用户三要素信息
type BlacklistUser struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name       string     `gorm:"column:name;type:varchar(50);index" json:"name"`             // 姓名
	IDCard     string     `gorm:"column:id_card;type:varchar(50);uniqueIndex" json:"id_card"` // 身份证号（唯一）
	Phone      string     `gorm:"column:phone;type:varchar(50);index" json:"phone"`           // 手机号
	Status     int8       `gorm:"column:status;type:tinyint;default:1" json:"status"`         // 状态: 1-黑名单中 0-已解除
	CreateTime *time.Time `gorm:"column:create_time;type:datetime(3)" json:"create_time"`
	UpdateTime *time.Time `gorm:"column:update_time;type:datetime(3)" json:"update_time"`
}

func (BlacklistUser) TableName() string {
	return "t_blacklist_user"
}

// IsBlocked 是否在黑名单中
func (b *BlacklistUser) IsBlocked() bool {
	return b.Status == 1
}

// BlacklistRecord 黑名单记录 - 对应 t_blacklist_record 表
// 存储每次拉黑的详细记录（纯历史记录）
type BlacklistRecord struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID     int64      `gorm:"column:user_id;index" json:"user_id"`              // 关联黑名单用户ID
	Reason     string     `gorm:"column:reason;type:varchar(500)" json:"reason"`    // 拉黑原因
	Source     string     `gorm:"column:source;type:varchar(100)" json:"source"`    // 来源渠道
	Operator   string     `gorm:"column:operator;type:varchar(50)" json:"operator"` // 操作人
	Remark     string     `gorm:"column:remark;type:varchar(500)" json:"remark"`    // 备注
	CreateTime *time.Time `gorm:"column:create_time;type:datetime(3)" json:"create_time"`
}

func (BlacklistRecord) TableName() string {
	return "t_blacklist_record"
}
