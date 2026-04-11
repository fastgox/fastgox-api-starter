package testhelper

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"

	"github.com/fastgox/fastgox-api-starter/src/core/config"
	"github.com/fastgox/fastgox-api-starter/src/core/i18n"
	"github.com/fastgox/fastgox-api-starter/src/router"
	_ "github.com/fastgox/fastgox-api-starter/src/router/handle"
	"github.com/gin-gonic/gin"
)

var initOnce sync.Once

// Setup 初始化测试环境（全局只执行一次）
// 设置 APP_ENV=test 后触发 router init 链
func Setup() {
	initOnce.Do(func() {
		os.Setenv("APP_ENV", "test")
		gin.SetMode(gin.TestMode)
		// router 包的 init() 已在 import 时执行，
		// 此处确保 config/i18n 已初始化
		_ = config.GlobalConfig
		_ = router.Engine
	})
}

// SetupWithoutDB 初始化不依赖数据库的测试环境
// 只加载配置和 i18n，不触发 router/database init
func SetupWithoutDB() {
	initOnce.Do(func() {
		os.Setenv("APP_ENV", "test")
		gin.SetMode(gin.TestMode)
		_ = config.InitConfig()
		i18n.Init()
	})
}

// PerformRequest 执行 HTTP 请求并返回 recorder
func PerformRequest(method, path string, body string, headers ...map[string]string) *httptest.ResponseRecorder {
	var req *http.Request
	if body != "" {
		req = httptest.NewRequest(method, path, strings.NewReader(body))
	} else {
		req = httptest.NewRequest(method, path, nil)
	}
	req.Header.Set("Content-Type", "application/json")

	// 应用自定义 headers
	for _, h := range headers {
		for k, v := range h {
			req.Header.Set(k, v)
		}
	}

	w := httptest.NewRecorder()
	router.Engine.ServeHTTP(w, req)
	return w
}

// PerformAuthRequest 执行带 JWT token 的认证请求
func PerformAuthRequest(method, path, body, token string) *httptest.ResponseRecorder {
	return PerformRequest(method, path, body, map[string]string{
		"Authorization": "Bearer " + token,
	})
}
