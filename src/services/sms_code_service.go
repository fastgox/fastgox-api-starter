package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fastgox/fastgox-api-starter/src/core/config"
	"github.com/fastgox/fastgox-api-starter/src/models/dto/response"
	"github.com/fastgox/fastgox-api-starter/src/repository"
)

// SmsCodeService 短信验证码服务
type SmsCodeService struct {
}

var SmsCodeSvc = &SmsCodeService{}

// SendSmsCode 发送短信验证码
func (s *SmsCodeService) SendSmsCode(phone string) (*response.SendSmsCodeResult, error) {
	// 1. 验证手机号格式
	if !isValidPhone(phone) {
		return &response.SendSmsCodeResult{
			Success: false,
			Message: "手机号格式不正确",
		}, nil
	}

	// 2. 生成验证码（检查白名单）
	_, code := s.generateCode(phone)

	// 3. 保存验证码到数据库
	smsCode, err := repository.SmsCodeRepo.CreateSmsCode(phone, code, config.GlobalConfig.SmsCode.CodeExpireTime)
	if err != nil {
		return nil, fmt.Errorf("保存验证码失败: %v", err)
	}

	// 4. TODO: 对接实际短信服务商发送短信
	// 可参考 pkg 中 Provider 模式实现短信服务商集成

	// 5. 计算下次可发送时间
	nextSendAt := time.Now().Add(time.Duration(config.GlobalConfig.SmsCode.IntervalLimit) * time.Minute)

	result := &response.SendSmsCodeResult{
		Success:    true,
		Message:    "验证码发送成功",
		ExpireAt:   smsCode.ExpireTime,
		NextSendAt: &nextSendAt,
	}

	// 测试环境返回验证码
	if isDevelopmentMode() {
		result.Code = code
	}

	return result, nil
}

// VerifySmsCode 验证短信验证码
func (s *SmsCodeService) VerifySmsCode(phone, code string) (*response.VerifySmsCodeResult, error) {
	// 1. 验证手机号格式
	if !isValidPhone(phone) {
		return &response.VerifySmsCodeResult{
			Success: false,
			Message: "手机号格式不正确",
		}, nil
	}

	// 2. 验证验证码格式
	if len(code) != config.GlobalConfig.SmsCode.CodeLength {
		return &response.VerifySmsCodeResult{
			Success: false,
			Message: "验证码格式不正确",
		}, nil
	}

	// 3. 验证验证码并标记为已使用
	smsCode, err := repository.SmsCodeRepo.VerifyAndMarkUsed(phone, code)
	if err != nil {
		return nil, fmt.Errorf("验证验证码失败: %v", err)
	}

	if smsCode == nil {
		return &response.VerifySmsCodeResult{
			Success: false,
			Message: "验证码错误或已过期",
		}, nil
	}

	return &response.VerifySmsCodeResult{
		Success: true,
		Message: "验证码验证成功",
	}, nil
}

// generateCode 生成验证码
func (s *SmsCodeService) generateCode(phone string) (bool, string) {
	cfg := config.GlobalConfig.SmsCode

	// 检查是否在白名单中
	for _, whitePhone := range cfg.Whitelist {
		if phone == whitePhone {
			// 白名单手机号返回固定验证码
			if cfg.WhitelistCode != "" {
				return true, cfg.WhitelistCode
			}
			break
		}
	}

	// 非白名单手机号生成随机验证码
	codeLength := cfg.CodeLength
	maxValue := 1
	for i := 0; i < codeLength; i++ {
		maxValue *= 10
	}

	code := rand.Intn(maxValue)
	format := fmt.Sprintf("%%0%dd", codeLength)
	return false, fmt.Sprintf(format, code)
}

// isValidPhone 验证手机号格式
func isValidPhone(phone string) bool {
	if len(phone) != 11 {
		return false
	}

	// 简单的手机号格式验证
	if phone[0] != '1' {
		return false
	}

	// 验证是否全为数字
	for _, char := range phone {
		if char < '0' || char > '9' {
			return false
		}
	}

	return true
}

// isDevelopmentMode 判断是否为开发模式
func isDevelopmentMode() bool {
	// 从配置中读取环境信息
	return config.GlobalConfig.App.Env == "dev" || config.GlobalConfig.App.Debug
}
