package entity

import "time"

// UserAuth 用户认证实体
// 对应数据库表 t_user_auth
type UserAuth struct {
	ID            int64      `gorm:"column:id;primaryKey" json:"id"`
	UserID        *int64     `gorm:"column:user_id" json:"user_id,omitempty"`
	AuthType      *string    `gorm:"column:auth_type" json:"auth_type,omitempty"`
	AuthStatus    *int8      `gorm:"column:auth_status" json:"auth_status,omitempty"`
	IdCardNumber  *string    `gorm:"column:id_card_number" json:"id_card_number,omitempty"`
	RealName      *string    `gorm:"column:real_name" json:"real_name,omitempty"`
	IdCardFront   *string    `gorm:"column:id_card_front" json:"id_card_front,omitempty"`
	IdCardBack    *string    `gorm:"column:id_card_back" json:"id_card_back,omitempty"`
	FaceURL       *string    `gorm:"column:face_url" json:"face_url,omitempty"`
	FaceScore     *string    `gorm:"column:face_score" json:"face_score,omitempty"`
	Confidence    *string    `gorm:"column:confidence" json:"confidence,omitempty"`
	FaceTime      *time.Time `gorm:"column:face_time" json:"face_time,omitempty"`
	NativePlace   *string    `gorm:"column:native_place" json:"native_place,omitempty"`
	EffectiveDate *string    `gorm:"column:effective_date" json:"effective_date,omitempty"`
	Gender        *string    `gorm:"column:gender" json:"gender,omitempty"`
	Birthday      *string    `gorm:"column:birthday" json:"birthday,omitempty"`
	Nation        *string    `gorm:"column:nation" json:"nation,omitempty"`
	Issuing       *string    `gorm:"column:issuing" json:"issuing,omitempty"`
	Age           *string    `gorm:"column:age" json:"age,omitempty"`
	AuthTime      *time.Time `gorm:"column:auth_time" json:"auth_time,omitempty"`
	CreateTime    *time.Time `gorm:"column:create_time" json:"create_time,omitempty"`
	UpdateTime    *time.Time `gorm:"column:update_time" json:"update_time,omitempty"`
}

func (UserAuth) TableName() string {
	return "t_user_auth"
}
