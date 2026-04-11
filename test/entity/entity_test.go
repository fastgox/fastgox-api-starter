package entity_test

import (
	"testing"
	"time"

	"github.com/fastgox/fastgox-api-starter/src/models/entity"
)

// ===== User 实体测试 =====

func TestUserTableName(t *testing.T) {
	u := entity.User{}
	if u.TableName() != "t_user" {
		t.Errorf("TableName() = %q, want %q", u.TableName(), "t_user")
	}
}

func TestUserIsAuthenticated(t *testing.T) {
	tests := []struct {
		isAuth   int8
		expected bool
	}{
		{0, false},
		{1, true},
		{2, false},
	}

	for _, tt := range tests {
		u := &entity.User{IsAuth: tt.isAuth}
		if u.IsAuthenticated() != tt.expected {
			t.Errorf("IsAuthenticated() with isAuth=%d: got=%v, want=%v", tt.isAuth, u.IsAuthenticated(), tt.expected)
		}
	}
}

func TestUserIsActive(t *testing.T) {
	tests := []struct {
		status   int8
		expected bool
	}{
		{0, false},
		{1, true},
		{2, false},
	}

	for _, tt := range tests {
		u := &entity.User{Status: tt.status}
		if u.IsActive() != tt.expected {
			t.Errorf("IsActive() with status=%d: got=%v, want=%v", tt.status, u.IsActive(), tt.expected)
		}
	}
}

// ===== SmsCode 实体测试 =====

func TestSmsCodeTableName(t *testing.T) {
	s := entity.SmsCode{}
	if s.TableName() != "t_sms_codes" {
		t.Errorf("TableName() = %q, want %q", s.TableName(), "t_sms_codes")
	}
}

func TestSmsCodeIsExpired(t *testing.T) {
	// nil ExpireTime → 过期
	s := &entity.SmsCode{ExpireTime: nil}
	if !s.IsExpired() {
		t.Error("ExpireTime 为 nil 时应该返回过期")
	}

	// 过去的时间 → 过期
	past := time.Now().Add(-1 * time.Hour)
	s = &entity.SmsCode{ExpireTime: &past}
	if !s.IsExpired() {
		t.Error("过去的时间应该返回过期")
	}

	// 未来的时间 → 未过期
	future := time.Now().Add(1 * time.Hour)
	s = &entity.SmsCode{ExpireTime: &future}
	if s.IsExpired() {
		t.Error("未来的时间不应该返回过期")
	}
}

func TestSmsCodeIsValid(t *testing.T) {
	future := time.Now().Add(1 * time.Hour)
	past := time.Now().Add(-1 * time.Hour)

	tests := []struct {
		name     string
		code     *entity.SmsCode
		expected bool
	}{
		{"未使用且未过期", &entity.SmsCode{IsUsed: false, ExpireTime: &future}, true},
		{"已使用", &entity.SmsCode{IsUsed: true, ExpireTime: &future}, false},
		{"已过期", &entity.SmsCode{IsUsed: false, ExpireTime: &past}, false},
		{"已使用且已过期", &entity.SmsCode{IsUsed: true, ExpireTime: &past}, false},
		{"ExpireTime为nil", &entity.SmsCode{IsUsed: false, ExpireTime: nil}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.code.IsValid() != tt.expected {
				t.Errorf("IsValid() = %v, want %v", tt.code.IsValid(), tt.expected)
			}
		})
	}
}
