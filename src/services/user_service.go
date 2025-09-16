package services

import (
	"fmt"
	"time"

	"github.com/fastgox/fastgox-api-starter/src/models/dto/response"
	"github.com/fastgox/fastgox-api-starter/src/models/entity"
	"github.com/fastgox/fastgox-api-starter/src/repository"
	"github.com/fastgox/fastgox-api-starter/src/utils"
)

// UserService 用户服务
type UserService struct {
}

var UserSvc = &UserService{}

// SendLoginSms 发送登录短信验证码
func (s *UserService) SendLoginSms(phone string) (*response.SendLoginSmsResult, error) {
	// 调用短信验证码服务发送验证码
	result, err := SmsCodeSvc.SendSmsCode(phone)
	if err != nil {
		return nil, err
	}

	return &response.SendLoginSmsResult{
		Success:    result.Success,
		Message:    result.Message,
		ExpireAt:   result.ExpireAt,
		NextSendAt: result.NextSendAt,
	}, nil
}

// LoginWithSms 使用短信验证码登录
func (s *UserService) LoginWithSms(phone, code string) (*response.LoginWithSmsResult, error) {
	// 1. 验证短信验证码
	verifyResult, err := SmsCodeSvc.VerifySmsCode(phone, code)
	if err != nil {
		return nil, fmt.Errorf("验证短信验证码失败: %v", err)
	}

	if !verifyResult.Success {
		return &response.LoginWithSmsResult{
			Success: false,
			Message: verifyResult.Message,
		}, nil
	}

	// 2. 查找或创建用户
	user, isNewUser, err := UserSvc.FindOrCreateUserByPhone(phone)
	if err != nil {
		return nil, fmt.Errorf("查找或创建用户失败: %v", err)
	}

	// 3. 检查用户状态
	if user.Status == 0 {
		return &response.LoginWithSmsResult{
			Success: false,
			Message: "账户已被禁用，请联系客服",
		}, nil
	}

	// 4. 更新用户登录时间
	if err := UserSvc.UpdateLoginTime(user.ID); err != nil {
		// 更新失败不影响登录，只记录日志
		fmt.Printf("更新用户登录时间失败: %v\n", err)
	}

	// 5. 生成JWT令牌
	token, expiresAt, err := s.generateJWTToken(user)
	if err != nil {
		return nil, fmt.Errorf("生成JWT令牌失败: %v", err)
	}

	return &response.LoginWithSmsResult{
		Success:   true,
		Message:   "登录成功",
		User:      user,
		Token:     token,
		IsNewUser: isNewUser,
		ExpiresAt: expiresAt,
	}, nil
}

// generateJWTToken 生成JWT令牌
func (s *UserService) generateJWTToken(user *entity.User) (string, *time.Time, error) {
	return utils.GenerateJWT(user.ID, user.Phone)
}

// FindOrCreateUserByPhone 根据手机号查找或创建用户（业务逻辑）
func (s *UserService) FindOrCreateUserByPhone(phone string) (*entity.User, bool, error) {
	return repository.UserRepo.FindOrCreateUserByPhone(phone)
}

// UpdateLoginTime 更新用户登录时间
func (s *UserService) UpdateLoginTime(userID int64) error {
	return repository.UserRepo.UpdateLoginTime(userID)
}
