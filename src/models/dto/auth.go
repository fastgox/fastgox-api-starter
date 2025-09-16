package dto

import "time"

// =============================================================================
// TCP协议 - 认证相关请求响应结构体
// =============================================================================

// LoginRequest TCP登录请求
type LoginRequest struct {
	Token string `json:"token"` // JWT Token或其他认证凭据
}

// LoginResponse TCP登录响应
type LoginResponse struct {
	BaseResponse
	Data struct {
		UserID    int64     `json:"userId"`
		LoginTime time.Time `json:"loginTime"`
	} `json:"data"`
}

// HeartBeatRequest 心跳请求
type HeartBeatRequest struct {
	// 心跳通常不需要额外参数，但可以包含客户端状态信息
	ClientTime time.Time `json:"clientTime,omitempty"`
}

// HeartBeatResponse 心跳响应
type HeartBeatResponse struct {
	BaseResponse
	Data struct {
		Timestamp time.Time `json:"timestamp"`
		UserID    int64     `json:"userId"`
	} `json:"data"`
}

// LogoutRequest 登出请求
type LogoutRequest struct {
	// 通常登出不需要参数，会话信息从session获取
}

// LogoutResponse 登出响应
type LogoutResponse struct {
	BaseResponse
	Data struct {
		LogoutTime time.Time `json:"logoutTime"`
	} `json:"data"`
}
