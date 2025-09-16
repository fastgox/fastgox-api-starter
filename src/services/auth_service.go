package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/fastgox/fastgox-api-starter/src/core/config"
	"github.com/fastgox/fastgox-api-starter/src/models/entity"
	"github.com/fastgox/fastgox-api-starter/src/pkg/auth"
	"github.com/fastgox/fastgox-api-starter/src/pkg/ocr"
	"github.com/fastgox/fastgox-api-starter/src/repository"
)

// AuthService 身份认证服务
type AuthService struct {
}

var AuthSvc = &AuthService{}

// VerifyThreeElements 三要素认证（姓名、身份证、手机号）
func (s *AuthService) VerifyThreeElements(userID int64, name, idCard, mobile string) (*auth.AuthOutput, error) {
	// 获取身份认证提供商（智能获取，支持回退机制）
	provider, err := auth.AuthManager.Get(config.GlobalConfig.Auth.Engine)
	if err != nil {
		return nil, fmt.Errorf("没有可用的身份认证服务提供商: %v", err)
	}

	// 构造输入参数
	input := &auth.ThreeElementsInput{
		Name:   name,
		IdCard: idCard,
		Mobile: mobile,
	}

	// 调用三要素认证
	result, err := provider.Call(input)
	if err != nil {
		return nil, err
	}

	// 保存认证记录到数据库
	now := time.Now()
	authStatus := int8(0) // 0-失败, 1-成功
	if result.Success {
		authStatus = 1
	}

	userAuth := &entity.UserAuth{
		UserID:       &userID,
		AuthType:     stringPtr("three_elements"),
		AuthStatus:   &authStatus,
		IdCardNumber: &idCard,
		RealName:     &name,
		AuthTime:     &now,
		CreateTime:   &now,
		UpdateTime:   &now,
	}

	// 保存认证记录到数据库
	if err := repository.UserAuthRepo.Create(userAuth); err != nil {
		// 记录日志但不影响返回结果
		fmt.Printf("保存认证记录失败: %v\n", err)
	}

	// 如果认证成功，更新用户表中的IsAuth字段
	if result.Success {
		if err := repository.UserRepo.UpdateAuthStatus(userID, 1); err != nil {
			fmt.Printf("更新用户认证状态失败: %v\n", err)
			// 这个错误不应该影响认证结果，但需要记录
		} else {
			fmt.Printf("用户认证状态已更新: userID=%d, IsAuth=1\n", userID)
		}
	}

	return result, nil
}

// VerifyThreeElementsWithOCR 通过OCR识别身份证信息进行三要素认证
func (s *AuthService) VerifyThreeElementsWithOCR(userID int64, mobile, idCardFrontUrl, idCardBackUrl string) (*auth.AuthOutput, error) {
	// 获取OCR提供商
	provider, err := ocr.OcrManager.Get(config.GlobalConfig.OCR.Engine)
	if err != nil {
		return nil, fmt.Errorf("没有可用的OCR服务提供商: %v", err)
	}

	// 构造OCR输入参数
	input := &ocr.RecognizeIdCardInput{
		Url:  idCardFrontUrl,
		Body: "",
	}

	// 调用OCR识别
	ocrResult, err := provider.Call(input)
	if err != nil {
		return nil, fmt.Errorf("身份证OCR识别失败: %v", err)
	}

	// 从OCR结果中提取姓名和身份证号
	// 注意：这里需要根据实际的OCR服务返回格式来解析
	// 目前OCR返回的是string格式，可能需要JSON解析
	var ocrData map[string]interface{}
	if err := json.Unmarshal([]byte(ocrResult.Data), &ocrData); err != nil {
		return nil, fmt.Errorf("OCR结果解析失败: %v", err)
	}

	name := ""
	idCard := ""

	if nameVal, exists := ocrData["name"]; exists {
		name = fmt.Sprintf("%v", nameVal)
	}
	if idCardVal, exists := ocrData["id_card"]; exists {
		idCard = fmt.Sprintf("%v", idCardVal)
	}

	if name == "" || idCard == "" {
		return nil, fmt.Errorf("OCR识别结果不完整，无法获取姓名或身份证号")
	}

	// 调用原有的三要素认证方法
	return s.VerifyThreeElements(userID, name, idCard, mobile)
}

// stringPtr 辅助函数，返回字符串指针
func stringPtr(s string) *string {
	return &s
}
