package database_test

import (
	"testing"

	"github.com/fastgox/fastgox-api-starter/src/core/database"
)

func TestConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  *database.Config
		wantErr bool
	}{
		{
			"合法配置",
			&database.Config{
				Driver: "postgres",
				Host:   "localhost",
				Port:   5432,
				User:   "root",
				DBName: "testdb",
			},
			false,
		},
		{
			"空Host",
			&database.Config{
				Driver: "postgres",
				Host:   "",
				Port:   5432,
				User:   "root",
				DBName: "testdb",
			},
			true,
		},
		{
			"无效端口-0",
			&database.Config{
				Driver: "postgres",
				Host:   "localhost",
				Port:   0,
				User:   "root",
				DBName: "testdb",
			},
			true,
		},
		{
			"无效端口-超上限",
			&database.Config{
				Driver: "postgres",
				Host:   "localhost",
				Port:   99999,
				User:   "root",
				DBName: "testdb",
			},
			true,
		},
		{
			"空User",
			&database.Config{
				Driver: "postgres",
				Host:   "localhost",
				Port:   5432,
				User:   "",
				DBName: "testdb",
			},
			true,
		},
		{
			"空DBName",
			&database.Config{
				Driver: "postgres",
				Host:   "localhost",
				Port:   5432,
				User:   "root",
				DBName: "",
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfigValidateAutoFix(t *testing.T) {
	cfg := &database.Config{
		Driver:      "postgres",
		Host:        "localhost",
		Port:        5432,
		User:        "root",
		DBName:      "testdb",
		MaxOpenConn: -1,
		MaxIdleConn: 0,
	}

	err := cfg.Validate()
	if err != nil {
		t.Fatalf("合法配置验证失败: %v", err)
	}

	if cfg.MaxOpenConn != 100 {
		t.Errorf("MaxOpenConn 应自动修正为 100: got=%d", cfg.MaxOpenConn)
	}
	if cfg.MaxIdleConn != 10 {
		t.Errorf("MaxIdleConn 应自动修正为 10: got=%d", cfg.MaxIdleConn)
	}
}

func TestConfigDSN(t *testing.T) {
	tests := []struct {
		name   string
		config *database.Config
		want   string // 只检查包含的关键词
	}{
		{
			"PostgreSQL DSN",
			&database.Config{
				Driver:   "postgres",
				Host:     "localhost",
				Port:     5432,
				User:     "root",
				Password: "secret",
				DBName:   "testdb",
				SSLMode:  "disable",
				Timezone: "Asia/Shanghai",
			},
			"host=localhost",
		},
		{
			"MySQL DSN",
			&database.Config{
				Driver:   "mysql",
				Host:     "localhost",
				Port:     3306,
				User:     "root",
				Password: "secret",
				DBName:   "testdb",
				Timezone: "Asia/Shanghai",
			},
			"tcp(localhost:3306)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dsn := tt.config.DSN()
			if dsn == "" {
				t.Fatal("DSN 不应为空")
			}
			if !contains(dsn, tt.want) {
				t.Errorf("DSN 应包含 %q: got=%q", tt.want, dsn)
			}
		})
	}
}

func TestDefaultConfig(t *testing.T) {
	cfg := database.DefaultConfig()
	if cfg.Driver != "postgres" {
		t.Errorf("默认 Driver 应为 postgres: got=%q", cfg.Driver)
	}
	if cfg.Port != 5432 {
		t.Errorf("默认 Port 应为 5432: got=%d", cfg.Port)
	}
	if cfg.MaxOpenConn != 100 {
		t.Errorf("默认 MaxOpenConn 应为 100: got=%d", cfg.MaxOpenConn)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsImpl(s, substr))
}

func containsImpl(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
