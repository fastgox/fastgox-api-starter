package i18n_test

import (
	"testing"

	"github.com/fastgox/fastgox-api-starter/src/core/i18n"
)

func TestInit(t *testing.T) {
	// 不应 panic
	i18n.Init()
}

func TestT(t *testing.T) {
	i18n.Init()

	tests := []struct {
		msgID    string
		expected string // 中文默认翻译
	}{
		{i18n.Success, "操作成功"},
		{i18n.ServerError, "服务器内部错误"},
		{i18n.BadRequest, "请求参数错误"},
		{i18n.Unauthorized, "未授权"},
		{i18n.Forbidden, "禁止访问"},
		{i18n.NotFound, "资源不存在"},
		{i18n.TooManyRequests, "请求过于频繁"},
		{i18n.SmsSendSuccess, "验证码发送成功"},
		{i18n.LoginSuccess, "登录成功"},
		{i18n.FileRequired, "请选择要上传的文件"},
		{i18n.FileUploadSuccess, "文件上传成功"},
	}

	for _, tt := range tests {
		t.Run(tt.msgID, func(t *testing.T) {
			result := i18n.T(tt.msgID)
			if result != tt.expected {
				t.Errorf("T(%q) = %q, want %q", tt.msgID, result, tt.expected)
			}
		})
	}
}

func TestTWithLang(t *testing.T) {
	i18n.Init()

	tests := []struct {
		lang     string
		msgID    string
		expected string
	}{
		// 中文
		{"zh", i18n.Success, "操作成功"},
		{"zh", i18n.Unauthorized, "未授权"},
		// 英文
		{"en", i18n.Success, "Success"},
		{"en", i18n.Unauthorized, "Unauthorized"},
		{"en", i18n.BadRequest, "Bad request"},
		{"en", i18n.NotFound, "Resource not found"},
		{"en", i18n.SmsSendSuccess, "Verification code sent"},
	}

	for _, tt := range tests {
		t.Run(tt.lang+"/"+tt.msgID, func(t *testing.T) {
			result := i18n.TWithLang(tt.lang, tt.msgID)
			if result != tt.expected {
				t.Errorf("TWithLang(%q, %q) = %q, want %q", tt.lang, tt.msgID, result, tt.expected)
			}
		})
	}
}

func TestTUnknownKey(t *testing.T) {
	i18n.Init()

	// 未知 key 应返回 key 本身
	unknownKey := "this_key_does_not_exist"
	result := i18n.T(unknownKey)
	if result != unknownKey {
		t.Errorf("未知 key 应返回原文: got=%q, want=%q", result, unknownKey)
	}
}

func TestLangFromString(t *testing.T) {
	i18n.Init()

	tests := []struct {
		input    string
		expected string
	}{
		{"", "zh"},      // 空字符串返回默认语言
		{"en", "en"},    // 标准英文
		{"zh", "zh"},    // 标准中文
		{"invalid!", "zh"}, // 无效语言返回默认
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := i18n.LangFromString(tt.input)
			if result != tt.expected {
				t.Errorf("LangFromString(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestKeysConsistency(t *testing.T) {
	i18n.Init()

	// 所有在 keys.go 中定义的常量都应该在 zh.yaml 中有翻译
	keys := []string{
		i18n.Success, i18n.ServerError,
		i18n.BadRequest, i18n.ParamError, i18n.PhoneFormatError, i18n.CodeFormatError,
		i18n.Unauthorized, i18n.Forbidden, i18n.TokenMissing, i18n.TokenFmtError, i18n.TokenEmpty, i18n.TokenInvalid,
		i18n.NotFound, i18n.TooManyRequests,
		i18n.SmsSendSuccess, i18n.SmsSendFailed, i18n.LoginSuccess, i18n.LoginFailed,
		i18n.FileRequired, i18n.FileUploadSuccess, i18n.FileUploadFailed,
	}

	for _, key := range keys {
		// 中文翻译不应等于 key 本身（说明翻译存在）
		zhResult := i18n.TWithLang("zh", key)
		if zhResult == key {
			t.Errorf("中文翻译缺失: key=%q", key)
		}

		// 英文翻译也不应等于 key 本身
		enResult := i18n.TWithLang("en", key)
		if enResult == key {
			t.Errorf("英文翻译缺失: key=%q", key)
		}
	}
}
