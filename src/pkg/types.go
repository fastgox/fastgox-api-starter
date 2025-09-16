package pkg

import "time"

// BaseInput 通用输入结构体基类
type BaseInput struct {
	RequestID string    `json:"request_id,omitempty"` // 请求ID
	Timestamp time.Time `json:"timestamp,omitempty"`  // 请求时间戳
}

// BaseOutput 通用输出结构体基类
type BaseOutput struct {
	Success   bool      `json:"success"`              // 操作是否成功
	Code      string    `json:"code,omitempty"`       // 响应代码
	Message   string    `json:"message,omitempty"`    // 响应消息
	RequestID string    `json:"request_id,omitempty"` // 请求ID
	Provider  string    `json:"provider,omitempty"`   // 服务提供商名称
	CostTime  int64     `json:"cost_time,omitempty"`  // 耗时（毫秒）
	Timestamp time.Time `json:"timestamp,omitempty"`  // 响应时间戳
}
