package session

// ContextKey 类型定义，用于避免key冲突
type ContextKey string

// 聊天相关
const (

	// QueryKey 查询内容
	UserKey ContextKey = "user"
	// chatRequest
	ChatSessionKey ContextKey = "chat_session"
)
