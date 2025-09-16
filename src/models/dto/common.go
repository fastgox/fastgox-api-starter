package dto

// =============================================================================
// TCP协议 - 通用响应结构体
// =============================================================================

// BaseResponse TCP协议基础响应
type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrorResponse TCP协议错误响应
type ErrorResponse struct {
	BaseResponse
	Error struct {
		Type    string `json:"type"`
		Details string `json:"details"`
	} `json:"error"`
}

// SuccessResponse TCP协议成功响应
type SuccessResponse struct {
	BaseResponse
	Data interface{} `json:"data,omitempty"`
}

// TODO: 在这里可以添加更多TCP协议相关的通用结构体
// 比如：认证、聊天、游戏逻辑等相关的DTO
