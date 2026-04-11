package tcp

import "encoding/json"

// Message TCP消息
type Message struct {
	Cmd  string          `json:"cmd"`            // 命令路由标识
	Seq  string          `json:"seq,omitempty"`  // 请求序列号
	Data json.RawMessage `json:"data,omitempty"` // 业务数据
}

// Response TCP响应消息
type Response struct {
	Cmd     string      `json:"cmd"`
	Seq     string      `json:"seq,omitempty"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Handler TCP消息处理器
type Handler func(ctx *Context) error

// Middleware TCP中间件
type Middleware func(Handler) Handler
