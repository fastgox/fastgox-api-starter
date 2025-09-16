package request

// ThreeElementsVerifyRequest 三要素认证请求
type ThreeElementsVerifyRequest struct {
	Name   string `json:"name" binding:"required"`   // 姓名
	IdCard string `json:"idcard" binding:"required"` // 身份证号
}
