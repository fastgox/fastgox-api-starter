package ocr

import (
	"encoding/json"
	"fmt"

	"github.com/fastgox/fastgox-api-starter/src/core/config"
	"github.com/fastgox/fastgox-api-starter/src/pkg"
	"github.com/fastgox/fastgox-api-starter/src/pkg/transport/qiandun"
)

// QiandunOcrProvider 千盾OCR服务提供商
type QiandunOcrProvider struct {
	client *qiandun.Client
}

// NewQiandunOcrProvider 创建千盾OCR服务提供商
func NewQiandunOcrProvider() *QiandunOcrProvider {
	return &QiandunOcrProvider{}
}

// 延迟初始化客户端
func (q *QiandunOcrProvider) ensureClient() error {
	if q.client != nil {
		return nil
	}

	if config.GlobalConfig == nil {
		return fmt.Errorf("配置未初始化")
	}

	cfg := config.GlobalConfig.OCR.Qiandun
	transportConfig := &qiandun.Config{
		AppKey:    cfg.AppKey,
		AppSecret: cfg.AppSecret,
		BaseURL:   cfg.BaseURL,
		IsProd:    cfg.IsProd,
	}

	q.client = qiandun.NewClient(transportConfig)
	return nil
}

// GetName 获取提供商名称
func (q *QiandunOcrProvider) GetName() string {
	return ProviderNameQiandun
}

// Call 类型安全的调用方法
func (q *QiandunOcrProvider) Call(input *RecognizeIdCardInput) (*OcrOutput, error) {
	return q.RecognizeIdCard(input)
}

// QiandunIdCardRequest 千盾身份证识别请求
type QiandunIdCardRequest struct {
	ImageBase64 string `json:"imageBase64,omitempty"`
	ImageUrl    string `json:"imageUrl,omitempty"`
	CardSide    string `json:"cardSide,omitempty"`
}

// QiandunIdCardResult 千盾身份证识别结果
type QiandunIdCardResult struct {
	Name         string `json:"name"`         // 姓名
	Sex          string `json:"sex"`          // 性别
	Nation       string `json:"nation"`       // 民族
	Birth        string `json:"birth"`        // 出生日期
	Address      string `json:"address"`      // 地址
	PsnIdCardNum string `json:"psnIdCardNum"` // 身份证号码
	Authority    string `json:"authority"`    // 发证机关
	ValidDate    string `json:"validDate"`    // 证件有效期
}

// QiandunOcrResponse 千盾OCR响应（包含具体的OCR结果类型）
type QiandunOcrResponse struct {
	Result   QiandunIdCardResult `json:"result"`
	Code     string              `json:"code"`
	Msg      string              `json:"msg"`
	LogKey   string              `json:"logKey"`
	Ts       string              `json:"ts"`
	Success  bool                `json:"success"`
	CostTime int64               `json:"costTime"` // 耗时（毫秒）
}

// RecognizeIdCard 身份证识别
func (q *QiandunOcrProvider) RecognizeIdCard(input *RecognizeIdCardInput) (*OcrOutput, error) {
	// 确保客户端已初始化
	if err := q.ensureClient(); err != nil {
		return nil, err
	}

	// 处理图片参数 - 优先使用新的字段名
	imageBase64 := input.ImageBase64
	imageUrl := input.ImageUrl

	// 如果新字段为空，使用旧字段
	if imageBase64 == "" {
		imageBase64 = input.Body
	}
	if imageUrl == "" {
		imageUrl = input.Url
	}

	// 验证参数
	if imageBase64 != "" && imageUrl != "" {
		return nil, fmt.Errorf("imageBase64和imageUrl参数不能同时提供，请选择其中一种方式")
	}
	if imageBase64 == "" && imageUrl == "" {
		return nil, fmt.Errorf("imageBase64和imageUrl参数不能同时为空，请提供其中一种方式")
	}

	// 构建千盾API请求格式
	qiandunReq := QiandunIdCardRequest{
		ImageBase64: imageBase64,
		ImageUrl:    imageUrl,
		CardSide:    input.CardSide,
	}

	// 使用transport层发送请求
	resp, err := q.client.SendRequest("/open/api/ocr/idCard", qiandunReq)
	if err != nil {
		return nil, err
	}

	// 检查业务状态
	if !resp.Success || resp.Code != "200" {
		return nil, fmt.Errorf("千盾OCR识别失败: %s", resp.Msg)
	}

	// 解析具体的OCR结果
	var ocrResult QiandunIdCardResult
	if resultBytes, err := json.Marshal(resp.Result); err != nil {
		return nil, fmt.Errorf("解析OCR结果失败: %v", err)
	} else if err := json.Unmarshal(resultBytes, &ocrResult); err != nil {
		return nil, fmt.Errorf("解析OCR结果失败: %v", err)
	}

	// 构造统一响应格式
	result := &OcrOutput{
		BaseOutput: pkg.BaseOutput{
			Success:   resp.Success,
			Message:   resp.Msg,
			RequestID: resp.LogKey,
			Provider:  "qiandun",
			CostTime:  resp.CostTime,
		},
	}

	// 将结果转换为JSON字符串
	resultBytes, _ := json.Marshal(ocrResult)
	result.Data = string(resultBytes)

	return result, nil
}

func init() {
	// 注册千盾OCR服务提供商
	provider := NewQiandunOcrProvider()
	OcrManager.Register(provider.GetName(), provider)
}
