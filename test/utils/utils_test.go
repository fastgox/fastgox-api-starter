package utils_test

import (
	"testing"
	"time"

	"github.com/fastgox/fastgox-api-starter/src/utils"
)

func TestEncryptDecrypt(t *testing.T) {
	// 生成一个合法的 base64 编码 AES-128 密钥（16字节）
	key := "MTIzNDU2Nzg5MDEyMzQ1Ng==" // base64("1234567890123456")

	tests := []struct {
		name string
		data string
	}{
		{"空字符串", ""},
		{"普通文本", "hello world"},
		{"中文文本", "你好世界"},
		{"JSON数据", `{"key":"value","num":123}`},
		{"长文本", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encrypted, err := utils.Encrypt([]byte(tt.data), key)
			if err != nil {
				t.Fatalf("加密失败: %v", err)
			}

			if encrypted == "" && tt.data != "" {
				t.Fatal("加密结果不应为空")
			}

			decrypted, err := utils.Decrypt(encrypted, key)
			if err != nil {
				t.Fatalf("解密失败: %v", err)
			}

			if string(decrypted) != tt.data {
				t.Errorf("解密结果不匹配: got=%q, want=%q", string(decrypted), tt.data)
			}
		})
	}
}

func TestEncryptWithInvalidKey(t *testing.T) {
	_, err := utils.Encrypt([]byte("test"), "invalid-base64!")
	if err == nil {
		t.Fatal("应返回密钥错误")
	}
}

func TestDecryptWithInvalidKey(t *testing.T) {
	_, err := utils.Decrypt("validbase64==", "invalid-base64!")
	if err == nil {
		t.Fatal("应返回密钥错误")
	}
}

func TestDecryptWithInvalidData(t *testing.T) {
	key := "MTIzNDU2Nzg5MDEyMzQ1Ng=="
	_, err := utils.Decrypt("not-valid-base64!", key)
	if err == nil {
		t.Fatal("应返回格式错误")
	}
}

func TestDecryptWithShortData(t *testing.T) {
	key := "MTIzNDU2Nzg5MDEyMzQ1Ng=="
	// base64 编码的短数据（少于 NonceSize）
	_, err := utils.Decrypt("AQID", key)
	if err == nil {
		t.Fatal("应返回格式错误")
	}
}

func TestBcryptHashAndCheck(t *testing.T) {
	password := "mySecurePassword123!"

	hash := utils.BcryptHash(password)
	if hash == "" {
		t.Fatal("哈希结果不应为空")
	}

	// 正确密码
	if !utils.BcryptCheck(password, hash) {
		t.Error("正确密码应验证通过")
	}

	// 错误密码
	if utils.BcryptCheck("wrongPassword", hash) {
		t.Error("错误密码应验证失败")
	}
}

func TestMD5V(t *testing.T) {
	result := utils.MD5V([]byte("hello"))
	if result == "" {
		t.Fatal("MD5 结果不应为空")
	}
	// MD5("hello") = 5d41402abc4b2a76b9719d911017c592
	expected := "5d41402abc4b2a76b9719d911017c592"
	if result != expected {
		t.Errorf("MD5 结果不匹配: got=%s, want=%s", result, expected)
	}
}

func TestGenerateOrderNo(t *testing.T) {
	orderNo := utils.GenerateOrderNo()

	// 基本格式检查
	if len(orderNo) == 0 {
		t.Fatal("订单号不应为空")
	}

	if orderNo[:3] != "ORD" {
		t.Errorf("订单号应以 ORD 开头: got=%s", orderNo)
	}

	// 唯一性检查
	seen := make(map[string]bool)
	for i := 0; i < 1000; i++ {
		no := utils.GenerateOrderNo()
		if seen[no] {
			t.Fatalf("订单号重复: %s (iteration %d)", no, i)
		}
		seen[no] = true
	}
}

func TestGenerateOrderNoConcurrency(t *testing.T) {
	const goroutines = 10
	const perGoroutine = 100

	results := make(chan string, goroutines*perGoroutine)

	for g := 0; g < goroutines; g++ {
		go func() {
			for i := 0; i < perGoroutine; i++ {
				results <- utils.GenerateOrderNo()
			}
		}()
	}

	seen := make(map[string]bool)
	for i := 0; i < goroutines*perGoroutine; i++ {
		no := <-results
		if seen[no] {
			t.Fatalf("并发下订单号重复: %s", no)
		}
		seen[no] = true
	}
}

func TestGetRealIP(t *testing.T) {
	// GetRealIP 依赖 gin.Context，跳过（需集成测试）
	t.Skip("需要 gin.Context，在集成测试中覆盖")
}

// BenchmarkGenerateOrderNo 基准测试
func BenchmarkGenerateOrderNo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.GenerateOrderNo()
	}
}

func BenchmarkBcryptHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.BcryptHash("testPassword123")
	}
}

func BenchmarkEncryptDecrypt(b *testing.B) {
	key := "MTIzNDU2Nzg5MDEyMzQ1Ng=="
	data := []byte("benchmark test data for encryption")

	for i := 0; i < b.N; i++ {
		enc, _ := utils.Encrypt(data, key)
		utils.Decrypt(enc, key)
	}
}

// ===== JWT 测试（需要配置初始化） =====

// TestJWTWithMockConfig 在不依赖全局配置的情况下验证 JWT 工具函数的逻辑
// 注意：完整的 JWT 测试需要集成测试环境
func TestJWTTokenExpiry(t *testing.T) {
	_ = time.Now()
	// JWT 测试需要初始化 config.GlobalConfig
	// 在集成测试中覆盖
	t.Skip("JWT 测试依赖 GlobalConfig，在集成测试中覆盖")
}
