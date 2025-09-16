package session

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SessionManager 会话管理器，提供所有会话操作方法
type SessionManager struct{}

// NewSessionManager 创建会话管理器
func NewSessionManager() *SessionManager {
	return &SessionManager{}
}

// Set 设置会话值
func (sm *SessionManager) Set(c *gin.Context, key string, value interface{}) {
	c.Set(key, value)
}

// GetValue 获取原始值
func (sm *SessionManager) GetValue(c *gin.Context, key string) (interface{}, bool) {
	return c.Get(key)
}

// GetString 获取字符串值
func (sm *SessionManager) GetString(c *gin.Context, key string) (string, error) {
	value, exists := c.Get(key)
	if !exists {
		return "", errors.New("键不存在: " + key)
	}

	if str, ok := value.(string); ok {
		return str, nil
	}

	return "", errors.New("值不是字符串类型: " + key)
}

// GetInt64 获取int64值
func (sm *SessionManager) GetInt64(c *gin.Context, key string) (int64, error) {
	value, exists := c.Get(key)
	if !exists {
		return 0, errors.New("键不存在: " + key)
	}

	switch v := value.(type) {
	case int64:
		return v, nil
	case string:
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, errors.New("字符串转换为int64失败: " + key)
		}
		return id, nil
	case int:
		return int64(v), nil
	case int32:
		return int64(v), nil
	default:
		return 0, errors.New("值不是数字类型: " + key)
	}
}

// SetJSON 设置JSON值
func (sm *SessionManager) SetJSON(c *gin.Context, key string, value interface{}) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}
	sm.Set(c, key, string(jsonData))
	return nil
}

// GetJSON 获取JSON值并反序列化到指定类型
func (sm *SessionManager) GetJSON(c *gin.Context, key string, result interface{}) error {
	jsonStr, err := sm.GetString(c, key)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(jsonStr), result)
	if err != nil {
		return errors.New("JSON反序列化失败: " + key)
	}

	return nil
}

// Remove 移除会话值
func (sm *SessionManager) Remove(c *gin.Context, key string) {
	sm.Set(c, key, nil)
}

// Clear 清除多个会话值
func (sm *SessionManager) Clear(c *gin.Context, keys ...string) {
	for _, key := range keys {
		sm.Set(c, key, nil)
	}
}

// Exists 检查键是否存在
func (sm *SessionManager) Exists(c *gin.Context, key string) bool {
	_, exists := c.Get(key)
	return exists
}

// GetEntity 获取实体（通过JSON反序列化）
func (sm *SessionManager) GetEntity(c *gin.Context, key string, result interface{}) error {
	// 先尝试直接获取
	if value, ok := sm.GetValue(c, key); ok {
		// 如果直接获取到的就是目标类型，直接返回
		if targetValue, ok := value.(interface{}); ok {
			// 这里可以添加类型检查逻辑
			_ = targetValue
		}
	}

	// 尝试从JSON格式获取（使用 key+"_json" 键）
	return sm.GetJSON(c, key+"_json", result)
}

// SetEntity 设置实体（同时保存原始值和JSON）
func (sm *SessionManager) SetEntity(c *gin.Context, key string, entity interface{}) error {
	// 直接设置原始值
	sm.Set(c, key, entity)

	// 同时也保存为JSON格式，便于跨请求访问
	return sm.SetJSON(c, key+"_json", entity)
}

// GetEntityDirect 直接获取实体（不尝试JSON反序列化）
func (sm *SessionManager) GetEntityDirect(c *gin.Context, key string) (interface{}, error) {
	value, ok := sm.GetValue(c, key)
	if !ok {
		return nil, errors.New("实体不存在: " + key)
	}
	return value, nil
}

// 全局会话管理器实例
var Manager = NewSessionManager()
