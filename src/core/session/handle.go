package session

import (
	"github.com/gin-gonic/gin"
)

// SetUserSession 设置用户会话
func SetUserSession(c *gin.Context, user interface{}) {
	c.Set(string(UserKey), user)
}

// GetUserSession 获取用户会话
func GetUserSession(c *gin.Context) interface{} {
	value, exists := c.Get(string(UserKey))
	if !exists {
		return nil
	}
	return value
}
