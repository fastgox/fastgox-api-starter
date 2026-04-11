package middleware

import (
	"strings"

	"github.com/fastgox/fastgox-api-starter/src/core/i18n"
	"github.com/fastgox/fastgox-api-starter/src/core/session"
	"github.com/fastgox/fastgox-api-starter/src/core/tcp"
	"github.com/fastgox/fastgox-api-starter/src/models/dto/response"
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
			response.Unauthorized(c, i18n.TokenMissing)
			c.Abort()
			return
		}

		// 检查Bearer前缀
		if !strings.HasPrefix(authHeader, "Bearer ") {
			response.Unauthorized(c, i18n.TokenFmtError)
			c.Abort()
			return
		}

		// 提取token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			response.Unauthorized(c, i18n.TokenEmpty)
			c.Abort()
			return
		}

		// 验证JWT token的有效性
		claims, err := utils.ValidateJWT(token)
		if err != nil {
			response.Unauthorized(c, i18n.TokenInvalid)
			c.Abort()
			return
		}
		user, _ := repository.UserRepo.GetByID(claims.UserID)

		// 将用户信息存储到上下文中 - 使用新的session包
		session.Manager.SetUserSession(c, user)
		c.Next()
	}
}

// TCPAuthMiddleware TCP鉴权中间件
func TCPAuthMiddleware() tcp.Middleware {
	return func(next tcp.Handler) tcp.Handler {
		return func(ctx *tcp.Context) error {
			token := ctx.GetString("token")
			if token == "" {
				response.Unauthorized(ctx, i18n.TokenMissing)
				return nil
			}

			claims, err := utils.ValidateJWT(token)
			if err != nil {
				response.Unauthorized(ctx, i18n.TokenInvalid)
				return nil
			}

			user, _ := repository.UserRepo.GetByID(claims.UserID)
			if user != nil {
				ctx.Set("user_id", user.ID)
				ctx.Set("user", user)
			}

			return next(ctx)
		}
	}
}
