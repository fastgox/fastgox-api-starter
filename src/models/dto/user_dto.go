package dto

// UserInfoResponse 用户信息响应
type UserInfoResponse struct {
	SaveStatus  SaveStatus `json:"save_status"`  // 保存状态
	Quota       int        `json:"quota"`        // 额度
	ApplyStatus int        `json:"apply_status"` // 申请状态：-1=测试，0=未申请，1=匹配中，2=撞库失败，3=等待审核，4=审核成功，5=审核失败
	URL         string     `json:"url"`          // 当申请状态为审核成功时，返回跳转URL
}

// SaveStatus 保存状态
type SaveStatus struct {
	Auth    int `json:"auth"`    // 认证状态 0=未认证，1=已认证
	Base    int `json:"base"`    // 基本信息：0=未保存，1=已保存
	Contact int `json:"contact"` // 紧急联系人：0=未保存，1=已保存
	Job     int `json:"job"`     // 职业信息：0=未保存，1=已保存
}
