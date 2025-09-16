package response

// ApplicationStepResponse 申请步骤响应
type ApplicationStepResponse struct {
	Step     string `json:"step"`      // 当前步骤
	Success  bool   `json:"success"`   // 是否成功
	Message  string `json:"message"`   // 响应消息
	NextStep string `json:"next_step"` // 下一步骤
	Progress int    `json:"progress"`  // 进度百分比
}

// BasicInfoResponse 基本信息查询响应
type BasicInfoResponse struct {
	Education            *int8   `json:"education,omitempty"`              // 学历
	MaritalStatus        *int8   `json:"marital_status,omitempty"`         // 婚姻状态
	MonthlyIncome        *string `json:"monthly_income,omitempty"`         // 月收入字符串
	MonthlyAverageIncome *int8   `json:"monthly_average_income,omitempty"` // 月收入等级
	ZhimaScore           *string `json:"zhima_score,omitempty"`            // 芝麻信用分
	MonthlyIncome2       *string `json:"monthly_income2,omitempty"`        // 月收入2
	SocialSecurity       *string `json:"social_security,omitempty"`        // 社保情况
	ProvidentFund        *string `json:"provident_fund,omitempty"`         // 公积金
	HouseStatus          *string `json:"house_status,omitempty"`           // 住房情况
	HousePrice           *string `json:"house_price,omitempty"`            // 房产价格
	CarStatus            *string `json:"car_status,omitempty"`             // 购车情况
	CarPrice             *string `json:"car_price,omitempty"`              // 购车总额
	CarPurchaseYear      *string `json:"car_purchase_year,omitempty"`      // 购车年限
}

// ContactResponse 联系人信息查询响应
type ContactResponse struct {
	Contacts []ContactInfoResponse `json:"contacts"` // 联系人列表
}

// ContactInfoResponse 联系人信息响应
type ContactInfoResponse struct {
	ID           int64  `json:"id"`            // 联系人ID
	ContactName  string `json:"contact_name"`  // 联系人姓名
	ContactPhone string `json:"contact_phone"` // 联系人电话
	Relationship int8   `json:"relationship"`  // 关系
	ContactType  int8   `json:"contact_type"`  // 联系人类型
}

// OccupationResponse 职业信息查询响应
type OccupationResponse struct {
	Province       *string `json:"province,omitempty"`        // 居住省
	City           *string `json:"city,omitempty"`            // 居住市
	Area           *string `json:"area,omitempty"`            // 居住区
	Address        *string `json:"address,omitempty"`         // 居住详细地址
	CompanyName    *string `json:"company_name,omitempty"`    // 公司名称
	CompanyAddress *string `json:"company_address,omitempty"` // 公司地址
	CompanyPhone   *string `json:"company_phone,omitempty"`   // 公司电话
	OccupationType *int8   `json:"occupation_type,omitempty"` // 职业类别
	Industry       *int8   `json:"industry,omitempty"`        // 行业类型
	LoanPurpose    *string `json:"loan_purpose,omitempty"`    // 借款用途
}
