package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fastgox/fastgox-api-starter/src/core/config"
	"github.com/fastgox/fastgox-api-starter/src/models/dto/response"
	"github.com/fastgox/fastgox-api-starter/src/pkg/sms"
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

	// // 2. 检查发送频率限制
	// cfg := config.GlobalConfig.SmsCode
	// if err := repository.SmsCodeRepo.CheckSendLimits(phone, cfg.DailyLimit, cfg.IntervalLimit, cfg.RecentLimit, cfg.RecentMinutes); err != nil {
	// 	return &response.SendSmsCodeResult{
	// 		Success: false,
	// 		Message: err.Error(),
	// 	}, nil
	// }

	// 3. 生成验证码（检查白名单）
	isWhiteListPhone, code := s.generateCode(phone)

	// 4. 保存验证码到数据库
	smsCode, err := repository.SmsCodeRepo.CreateSmsCode(phone, code, config.GlobalConfig.SmsCode.CodeExpireTime)
	if err != nil {
		return nil, fmt.Errorf("保存验证码失败: %v", err)
	}

	// 5. 发送短信（这里可以集成实际的短信服务）
	if !isWhiteListPhone {
		if err := s.sendSmsToPhone(phone, code); err != nil {
			return nil, fmt.Errorf("发送短信失败: %v", err)
		}
	}

	// 6. 计算下次可发送时间
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

// sendSmsToPhone 发送短信到手机号
func (s *SmsCodeService) sendSmsToPhone(phone, code string) error {
	cfg := config.GlobalConfig.SmsCode

	// 检查是否在白名单中，白名单手机号跳过实际发送
	for _, whitePhone := range cfg.Whitelist {
		if phone == whitePhone {
			// 白名单手机号，跳过实际发送，直接返回成功
			return nil
		}
	}

	// 非白名单手机号，执行实际发送
	// 获取SMS提供商（智能获取，支持回退机制）
	provider, err := sms.SmsManager.Get(config.GlobalConfig.SMS.Engine)
	if err != nil {
		return fmt.Errorf("没有可用的SMS服务提供商: %v", err)
	}

	// 构造输入参数
	input := &sms.SendSmsInput{
		Phone: phone,
		Code:  code,
	}

	// 调用发送短信
	result, err := provider.Call(input)
	if err != nil {
		return fmt.Errorf("发送短信失败: %v", err)
	}

	if !result.BaseOutput.Success {
		return fmt.Errorf("发送短信失败: %s", result.BaseOutput.Message)
	}

	return nil
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
