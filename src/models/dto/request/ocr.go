package request

// RecognizeIdCardRequest 身份证识别请求
type RecognizeIdCardRequest struct {
	Url      string `json:"url"`      // 图片URL
	Body     string `json:"body"`     // 图片Base64编码
	CardSide string `json:"cardSide"` // 卡面方向：front/back
}

// OpenIdCardRequest 开放接口身份证识别请求
type OpenIdCardRequest struct {
	ImageBase64 string `json:"imageBase64,omitempty"` // 身份证图片的base64编码
	ImageUrl    string `json:"imageUrl,omitempty"`    // 身份证图片的url地址
	CardSide    string `json:"cardSide,omitempty"`    // 卡面方向：front/back
}
