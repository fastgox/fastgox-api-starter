package dto

// UVCustomerQueryRequest UV客户查询请求
type UVCustomerQueryRequest struct {
	Page     int `form:"page" binding:"omitempty,min=1"`              // 页码，默认1
	PageSize int `form:"page_size" binding:"omitempty,min=1,max=100"` // 每页数量，默认10
}

// UVCustomerResponse UV客户响应
type UVCustomerResponse struct {
	ID             int64    `json:"id"`                         // 主键ID
	ProductName    *string  `json:"product_name,omitempty"`     // 产品名称
	ProductTag     *string  `json:"product_tag,omitempty"`      // 产品标签(逗号分隔)
	ProductIcon    *string  `json:"product_icon,omitempty"`     // 产品Logo
	MaxAmount      *float32 `json:"max_amount,omitempty"`       // 产品最高额度
	AnnualRate     *float32 `json:"annual_rate,omitempty"`      // 年利率
	DownloadAppURL *string  `json:"download_app_url,omitempty"` // APP下载链接
}

// UVCustomerListResponse UV客户列表响应
type UVCustomerListResponse struct {
	List  []*UVCustomerResponse `json:"list"`  // 客户列表
	Total int64                 `json:"total"` // 总数
	Page  int                   `json:"page"`  // 当前页码
	Size  int                   `json:"size"`  // 每页数量
}

// UVVisitRecordRequest UV客户访问记录请求
type UVVisitRecordRequest struct {
	CustomerID int64 `json:"customer_id" binding:"required,min=1"` // 客户ID
}

// UVVisitRecordResponse UV客户访问记录响应
type UVVisitRecordResponse struct {
	ID         int64  `json:"id"`          // 记录ID
	CustomerID int64  `json:"customer_id"` // 客户ID
	UserID     int64  `json:"user_id"`     // 用户ID
	CreatedAt  string `json:"created_at"`  // 创建时间
	Message    string `json:"message"`     // 响应消息
}

// 撞库请求dto
type CustomerCollideInput struct {
	ChannelCode string `json:"channel_code"` // 渠道编码
	IdCardMd5   string `json:"id_card_md5"`  // 身份证号MD5
	MobileMd5   string `json:"mobile_md5"`   // 手机号MD5
	Key         string `json:"key"`          // 加密密钥
	Url         string `json:"url"`          // 请求地址
}

// 撞库响应dto
type CustomerCollideOut struct {
	Code    int    `json:"code"`    // 响应码，0=成功，其他值表示失败
	Message string `json:"message"` // 响应消息
}

// 订单申请请求dto
type CustomerApplicationApplyInput struct {
	UserId int64 `json:"user_id"` // 用户ID
	//客户Code
	ProductCode string `json:"product_code"` // 客户ID

}

// 撞库响应dto
type CustomerApplicationApplyOut struct {
	Code        int     `json:"code"`        // SUCCESS=成功，其他值表示失败
	Message     string  `json:"message"`     // 响应消息
	CheckResult *bool   `json:"checkResult"` // 授信是否成功,如果审核结果可以立马返回，可以传，就不用调用授信回调
	CreditMoney *int64  `json:"creditMoney"` // 授信成功返回授信额度(分为单位)
	ExpireTime  *int64  `json:"expireTime"`  // 过期时间
	FailReason  *string `json:"failReason"`  // 授信失败拒绝原因
	Msg         string  `json:"msg"`         // 响应消息
}

// 接口加密数据结构响应
type ApiEncryptedResp struct {
	Code    int    `json:"code"`    // 响应码，0=成功，其他值表示失败
	Message string `json:"message"` // 响应消息
	Data    string `json:"data"`    // 加密数据
}
