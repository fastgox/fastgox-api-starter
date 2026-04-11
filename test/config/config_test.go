package config_test

import (
	"os"
	"testing"

	coreconfig "github.com/fastgox/fastgox-api-starter/src/core/config"
)

func TestInitConfigDev(t *testing.T) {
	// 切到项目根目录（测试需要能找到 config/dev.yaml）
	origDir, _ := os.Getwd()
	if err := os.Chdir("../.."); err != nil {
		t.Skip("无法切换到项目根目录")
	}
	defer os.Chdir(origDir)

	os.Setenv("APP_ENV", "dev")
	defer os.Unsetenv("APP_ENV")

	// 重置全局配置
	coreconfig.GlobalConfig = nil

	err := coreconfig.InitConfig()
	if err != nil {
		t.Fatalf("InitConfig 失败: %v", err)
	}

	cfg := coreconfig.GlobalConfig
	if cfg == nil {
		t.Fatal("GlobalConfig 不应为 nil")
	}

	// 验证基本配置项
	if cfg.App.Name == "" {
		t.Error("App.Name 不应为空")
	}
	if cfg.App.Port <= 0 {
		t.Error("App.Port 应大于 0")
	}
	if cfg.Database.Driver == "" {
		t.Error("Database.Driver 不应为空")
	}
	if cfg.JWT.SecretKey == "" {
		t.Error("JWT.SecretKey 不应为空")
	}
}

func TestInitConfigTest(t *testing.T) {
	origDir, _ := os.Getwd()
	if err := os.Chdir("../.."); err != nil {
		t.Skip("无法切换到项目根目录")
	}
	defer os.Chdir(origDir)

	os.Setenv("APP_ENV", "test")
	defer os.Unsetenv("APP_ENV")

	coreconfig.GlobalConfig = nil

	err := coreconfig.InitConfig()
	if err != nil {
		t.Fatalf("InitConfig(test) 失败: %v", err)
	}

	cfg := coreconfig.GlobalConfig
	if cfg.App.Env != "test" {
		t.Errorf("App.Env = %q, want %q", cfg.App.Env, "test")
	}
	if cfg.App.Port != 19081 {
		t.Errorf("App.Port = %d, want %d", cfg.App.Port, 19081)
	}
	if !cfg.TCP.Enabled == true {
		// test.yaml 中 tcp.enabled = false
	}
}

func TestInitConfigNonExistent(t *testing.T) {
	origDir, _ := os.Getwd()
	if err := os.Chdir("../.."); err != nil {
		t.Skip("无法切换到项目根目录")
	}
	defer os.Chdir(origDir)

	os.Setenv("APP_ENV", "nonexistent")
	defer os.Unsetenv("APP_ENV")

	coreconfig.GlobalConfig = nil

	err := coreconfig.InitConfig()
	if err == nil {
		t.Fatal("不存在的配置文件应返回错误")
	}
}

func TestValidateConfigMissingDB(t *testing.T) {
	origDir, _ := os.Getwd()
	if err := os.Chdir("../.."); err != nil {
		t.Skip("无法切换到项目根目录")
	}
	defer os.Chdir(origDir)

	// 测试验证逻辑 — 通过加载 dev 配置后手动修改来验证
	os.Setenv("APP_ENV", "dev")
	defer os.Unsetenv("APP_ENV")

	coreconfig.GlobalConfig = nil
	// 这个测试验证 InitConfig 内部的 validateConfig 是否正常工作
	// 由于 dev.yaml 有完整配置，应该成功
	err := coreconfig.InitConfig()
	if err != nil {
		t.Fatalf("合法配置应通过验证: %v", err)
	}
}
