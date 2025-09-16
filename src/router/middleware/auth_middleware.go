package middleware

import (
	"net/http"
	"strings"

	"github.com/fastgox/fastgox-api-starter/src/core/session"
	"github.com/fastgox/fastgox-api-starter/src/repository"
	"github.com/fastgox/fastgox-api-starter/src/utils"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware JWT认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization头
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "缺少认证token",
			})
			c.Abort()
			return
		}

		// 检查Bearer前缀
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "token格式错误",
			})
			c.Abort()
			return
		}

		// 提取token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "token不能为空",
			})
			c.Abort()
			return
		}

		// 验证JWT token的有效性
		claims, err := utils.ValidateJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "token无效: " + err.Error(),
			})
			c.Abort()
			return
		}
		user, _ := repository.UserRepo.GetByID(claims.UserID)

		// 将用户信息存储到上下文中 - 使用新的session包
		session.Manager.SetUserSession(c, user)
		c.Next()
	}
}
