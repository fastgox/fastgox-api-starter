package ocr

import (
	"github.com/fastgox/fastgox-api-starter/src/pkg"
)

// 提供商名称常量
const (
	ProviderNameAliyun  = "aliyun"  // 阿里云提供商
	ProviderNameQiandun = "qiandun" // 千盾提供商
)

// 默认提供商
const DefaultProvider = ProviderNameAliyun

// 全局OCR服务提供商管理器
var OcrManager = pkg.NewManager[pkg.Provider[*RecognizeIdCardInput, *OcrOutput]](DefaultProvider)

// RecognizeIdCardInput 身份证识别输入
type RecognizeIdCardInput struct {
	pkg.BaseInput
	Url         string `json:"url"`         // 图片URL
	Body        string `json:"body"`        // 图片Base64编码
	ImageBase64 string `json:"imageBase64"` // 千盾OCR使用的Base64字段
	ImageUrl    string `json:"imageUrl"`    // 千盾OCR使用的URL字段
	CardSide    string `json:"cardSide"`    // 卡面方向：front/back
}

// OcrOutput OCR识别响应
type OcrOutput struct {
	pkg.BaseOutput
	Data string `json:"data"` // 识别结果数据
}
