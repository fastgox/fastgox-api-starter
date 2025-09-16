package repository

import (
	"fmt"
	"time"

	"github.com/fastgox/fastgox-api-starter/src/models/entity"
)

// SmsCodeRepository 短信验证码仓储
type SmsCodeRepository struct {
	BaseRepository[entity.SmsCode]
}

var SmsCodeRepo = &SmsCodeRepository{BaseRepository: *GetRepository[entity.SmsCode]()}

// GetValidCodeByPhone 根据手机号获取有效的验证码
func (r *SmsCodeRepository) GetValidCodeByPhone(phone string) (*entity.SmsCode, error) {
	var smsCode entity.SmsCode
	err := r.DB.Where("phone = ? AND is_used = ? AND expire_time > ?",
		phone, false, time.Now()).
		Order("created_at DESC").
		First(&smsCode).Error

	if err != nil {
		return nil, err
	}
	return &smsCode, nil
}

// CreateSmsCode 创建短信验证码
func (r *SmsCodeRepository) CreateSmsCode(phone, code string, expireMinutes int) (*entity.SmsCode, error) {
	now := time.Now()
	expireTime := now.Add(time.Duration(expireMinutes) * time.Minute)

	smsCode := &entity.SmsCode{
		Phone:      phone,
		Code:       code,
		ExpireTime: &expireTime,
		IsUsed:     false,
		CreatedAt:  &now,
		UpdatedAt:  &now,
	}

	err := r.Create(smsCode)
	if err != nil {
		return nil, err
	}
	return smsCode, nil
}

// MarkAsUsed 标记验证码为已使用
func (r *SmsCodeRepository) MarkAsUsed(id uint) error {
	now := time.Now()
	return r.DB.Model(&entity.SmsCode{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_used":    true,
			"updated_at": now,
		}).Error
}

// CountTodayCodesByPhone 统计今天某手机号发送的验证码数量
func (r *SmsCodeRepository) CountTodayCodesByPhone(phone string) (int64, error) {
	today := time.Now().Format("2006-01-02")
	startTime := today + " 00:00:00"
	endTime := today + " 23:59:59"

	var count int64
	err := r.DB.Model(&entity.SmsCode{}).
		Where("phone = ? AND created_at BETWEEN ? AND ?", phone, startTime, endTime).
		Count(&count).Error

	return count, err
}

// CountRecentCodesByPhone 统计最近N分钟内某手机号发送的验证码数量
func (r *SmsCodeRepository) CountRecentCodesByPhone(phone string, minutes int) (int64, error) {
	recentTime := time.Now().Add(-time.Duration(minutes) * time.Minute)

	var count int64
	err := r.DB.Model(&entity.SmsCode{}).
		Where("phone = ? AND created_at > ?", phone, recentTime).
		Count(&count).Error

	return count, err
}

// VerifyCode 验证验证码
func (r *SmsCodeRepository) VerifyCode(phone, code string) (*entity.SmsCode, error) {
	// 获取有效的验证码
	smsCode, err := r.GetValidCodeByPhone(phone)
	if err != nil {
		return nil, err
	}

	// 验证验证码是否匹配
	if smsCode.Code != code {
		return nil, nil // 验证码不匹配
	}

	return smsCode, nil
}

// VerifyAndMarkUsed 验证验证码并标记为已使用
func (r *SmsCodeRepository) VerifyAndMarkUsed(phone, code string) (*entity.SmsCode, error) {
	// 验证验证码
	smsCode, err := r.VerifyCode(phone, code)
	if err != nil {
		return nil, err
	}

	if smsCode == nil {
		return nil, nil // 验证码不匹配或无效
	}

	// 标记为已使用
	if err := r.MarkAsUsed(smsCode.ID); err != nil {
		return nil, err
	}

	return smsCode, nil
}

// CheckSendLimits 检查发送限制
func (r *SmsCodeRepository) CheckSendLimits(phone string, dailyLimit, intervalMinutes, recentLimit, recentMinutes int) error {
	// 1. 检查今日发送次数限制
	todayCount, err := r.CountTodayCodesByPhone(phone)
	if err != nil {
		return err
	}
	if todayCount >= int64(dailyLimit) {
		return fmt.Errorf("今日发送次数已达上限，请明天再试")
	}

	// 2. 检查最近发送间隔（如果间隔限制大于0才检查）
	if intervalMinutes > 0 {
		recentCount, err := r.CountRecentCodesByPhone(phone, intervalMinutes)
		if err != nil {
			return err
		}
		if recentCount > 0 {
			return fmt.Errorf("发送过于频繁，请%d分钟后再试", intervalMinutes)
		}
	}

	// 3. 检查最近时间窗口内的发送次数
	recentWindowCount, err := r.CountRecentCodesByPhone(phone, recentMinutes)
	if err != nil {
		return err
	}
	if recentWindowCount >= int64(recentLimit) {
		return fmt.Errorf("发送过于频繁，请%d分钟后再试", recentMinutes)
	}

	return nil
}
