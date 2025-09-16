package session

import (
	"strings"

	"github.com/fastgox/fastgox-api-starter/src/models/entity"
	"github.com/fastgox/fastgox-api-starter/src/repository"
	"github.com/fastgox/fastgox-api-starter/src/utils"
	"github.com/gin-gonic/gin"
)

// 用户会话相关的键名常量
const (
	key = "user"
)

// UserSession 用户会话信息
type UserSession struct {
	UserID   int64  `json:"user_id"`
	Phone    string `json:"phone"`
	Platform string `json:"platform,omitempty"`
}

// SetUser
func (sm *SessionManager) SetUserSession(c *gin.Context, user *entity.User) error {
	userSession := &UserSession{
		UserID:   user.ID,
		Phone:    user.Phone,
		Platform: user.Platform,
	}
	return sm.SetEntity(c, key, userSession)
}

// GetUser 获取用户会话信息
func (sm *SessionManager) GetUserSession(c *gin.Context) (*UserSession, error) {
	var userSession *UserSession
	err := sm.GetEntity(c, key, &userSession)
	if err != nil {
		return nil, err
	}
	return userSession, nil
}

// GetUser 获取用户会话信息
func (sm *SessionManager) GetUserSessionAndByAuthorization(c *gin.Context) (*UserSession, error) {
	var userSession *UserSession
	err := sm.GetEntity(c, key, &userSession)
	if err != nil {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			userSession = sm.GetUserSessionByHeader(c)
			if userSession != nil {
				return userSession, nil
			}
		}
		return nil, err
	}
	return userSession, nil
}

// get userSession by header
func (sm *SessionManager) GetUserSessionByHeader(c *gin.Context) *UserSession {
	// 获取Authorization头
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil
	}

	// 检查Bearer前缀
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return nil
	}

	// 提取token
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		return nil
	}

	// 验证JWT token的有效性
	claims, err := utils.ValidateJWT(token)
	if err != nil {
		return nil
	}
	user, _ := repository.UserRepo.GetByID(claims.UserID)
	// 将用户信息存储到上下文中 - 使用新的session包
	Manager.SetUserSession(c, user)
	return &UserSession{
		UserID:   user.ID,
		Phone:    user.Phone,
		Platform: user.Platform,
	}
}
