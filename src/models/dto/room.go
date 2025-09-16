package dto

import "time"

// =============================================================================
// TCP协议 - 房间相关请求结构体
// =============================================================================

// JoinRoomRequest 加入房间请求
type JoinRoomRequest struct {
	RoomID string `json:"roomId"`
	// UserID 从认证会话中获取，无需客户端传递
}

// LeaveRoomRequest 离开房间请求
type LeaveRoomRequest struct {
	RoomID string `json:"roomId"`
	// UserID 从认证会话中获取，无需客户端传递
}

// GetRoomInfoRequest 获取房间信息请求
type GetRoomInfoRequest struct {
	RoomID string `json:"roomId"`
}

// =============================================================================
// TCP协议 - 房间相关响应结构体
// =============================================================================

// RoomInfoResponse 房间信息响应
type RoomInfoResponse struct {
	BaseResponse
	Data struct {
		RoomID      string    `json:"roomId"`
		PlayerCount int       `json:"playerCount"`
		MaxPlayers  int       `json:"maxPlayers"`
		Status      string    `json:"status"` // waiting, playing, finished
		CreatedTime time.Time `json:"createdTime"`
	} `json:"data"`
}
