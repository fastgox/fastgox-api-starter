package middleware

import (
	"github.com/fastgox/fastgox-api-starter/src/core/i18n"
	"github.com/gin-gonic/gin"
)

const LangKey = "lang"

// I18nMiddleware 从 Accept-Language header 解析语言并存入 context
func I18nMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := i18n.LangFromGin(c)
		c.Set(LangKey, lang)
		c.Next()
	}
}
