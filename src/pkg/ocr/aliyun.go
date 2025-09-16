package ocr

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ocr_api20210707 "github.com/alibabacloud-go/ocr-api-20210707/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/fastgox/fastgox-api-starter/src/core/config"
	"github.com/fastgox/fastgox-api-starter/src/pkg"
)

// AliyunOcrProvider 阿里云OCR服务提供商
type AliyunOcrProvider struct {
}

// NewAliyunOcrProvider 创建阿里云OCR服务提供商
func NewAliyunOcrProvider() *AliyunOcrProvider {
	return &AliyunOcrProvider{}
}

// GetName 获取提供商名称
func (a *AliyunOcrProvider) GetName() string {
	return ProviderNameAliyun
}

// Call 类型安全的调用方法
func (a *AliyunOcrProvider) Call(input *RecognizeIdCardInput) (*OcrOutput, error) {
	return a.RecognizeIdCard(input)
}

// createOCRClient 创建阿里云OCR客户端
func (a *AliyunOcrProvider) createOCRClient() (*ocr_api20210707.Client, error) {
	ocrConfig := &openapi.Config{
		AccessKeyId:     tea.String(config.GlobalConfig.OCR.Aliyun.AccessKeyID),
		AccessKeySecret: tea.String(config.GlobalConfig.OCR.Aliyun.AccessKeySecret),
	}

	// 设置endpoint
	endpoint := config.GlobalConfig.OCR.Aliyun.Endpoint
	if endpoint == "" {
		endpoint = "ocr-api.cn-hangzhou.aliyuncs.com"
	}
	ocrConfig.Endpoint = tea.String(endpoint)

	return ocr_api20210707.NewClient(ocrConfig)
}

// RecognizeIdCard 身份证识别
func (a *AliyunOcrProvider) RecognizeIdCard(input *RecognizeIdCardInput) (*OcrOutput, error) {
	ocrClient, err := a.createOCRClient()
	if err != nil {
		return nil, fmt.Errorf("创建OCR客户端失败: %v", err)
	}

	recognizeRequest := &ocr_api20210707.RecognizeAllTextRequest{
		Type: tea.String("IdCard"),
	}

	// 设置图片来源 - URL和Body二选一，不可同时为空
	if input.Url != "" && input.Body != "" {
		return nil, fmt.Errorf("URL和Body参数不能同时提供，请选择其中一种方式")
	}

	if input.Url == "" && input.Body == "" {
		return nil, fmt.Errorf("URL和Body参数不能同时为空，请提供其中一种方式")
	}

	if input.Url != "" {
		// 验证URL格式
		if _, err := url.Parse(input.Url); err != nil {
			return nil, fmt.Errorf("图片URL格式不正确: %v", err)
		}

		// 检查URL是否为HTTPS（阿里云OCR推荐使用HTTPS）
		if !strings.HasPrefix(input.Url, "http://") && !strings.HasPrefix(input.Url, "https://") {
			return nil, fmt.Errorf("图片URL必须以http://或https://开头")
		}

		recognizeRequest.Url = tea.String(input.Url)
	}

	// 处理Base64编码的图片数据
	if input.Body != "" {
		// 去除可能的Data URL前缀（如 "data:image/jpeg;base64,"）
		base64Data := input.Body
		if strings.Contains(base64Data, ",") {
			// 如果包含逗号，取逗号后面的部分作为Base64数据
			parts := strings.Split(base64Data, ",")
			if len(parts) > 1 {
				base64Data = parts[1]
			}
		}

		// 将Base64字符串解码为二进制数据
		bodyBytes, err := base64.StdEncoding.DecodeString(base64Data)
		if err != nil {
			return nil, fmt.Errorf("Base64解码失败: %v", err)
		}

		// 将[]byte转换为io.Reader
		recognizeRequest.Body = bytes.NewReader(bodyBytes)
	}

	runtime := &util.RuntimeOptions{}

	resp, err := ocrClient.RecognizeAllTextWithOptions(recognizeRequest, runtime)
	if err != nil {
		var sdkError = &tea.SDKError{}
		if t, ok := err.(*tea.SDKError); ok {
			sdkError = t
		} else {
			sdkError.Message = tea.String(err.Error())
		}

		// 解析错误详情
		var errorData interface{}
		if sdkError.Data != nil {
			d := json.NewDecoder(strings.NewReader(tea.StringValue(sdkError.Data)))
			d.Decode(&errorData)
		}

		return nil, fmt.Errorf("阿里云OCR识别失败: %s", tea.StringValue(sdkError.Message))
	}

	// 构造响应
	result := &OcrOutput{
		BaseOutput: pkg.BaseOutput{
			Success: true,
			Message: "识别成功",
		},
	}

	if resp.Body != nil {
		if resp.Body.Data != nil {
			// 将数据转换为JSON字符串
			dataBytes, _ := json.Marshal(resp.Body.Data)
			result.Data = string(dataBytes)
		}
		if resp.Body.RequestId != nil {
			result.BaseOutput.RequestID = tea.StringValue(resp.Body.RequestId)
		}
	}

	return result, nil
}

func init() {
	// 注册阿里云OCR服务提供商
	provider := NewAliyunOcrProvider()
	OcrManager.Register(provider.GetName(), provider)
}
