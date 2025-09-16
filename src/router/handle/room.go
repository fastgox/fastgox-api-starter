package handle

import (
	"time"

	"github.com/fastgox/fastgox-api-starter/src/models/dto"
	"github.com/fastgox/fastgox-api-starter/src/router"
	"github.com/fastgox/utils/logger"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
)

// RoomComponent 房间管理组件（示例）
type RoomComponent struct {
	component.Base
}

// Join 加入房间
func (r *RoomComponent) Join(s *session.Session, req *dto.JoinRoomRequest) error {
	// 检查会话中的character信息
	character := getCharacter(s)
	if character == nil {
		return s.Response(&dto.BaseResponse{
			Code:    401,
			Message: "未认证，请先登录",
		})
	}

	logger.Info("用户加入房间请求: %s (ID: %d) 房间: %s", character.Name, character.ID, req.RoomID)

	// 业务逻辑处理
	if req.RoomID == "" {
		return s.Response(&dto.BaseResponse{
			Code:    400,
			Message: "房间ID不能为空",
		})
	}

	// 模拟加入房间成功
	response := &dto.RoomInfoResponse{}
	response.Code = 200
	response.Message = "加入房间成功"
	response.Data.RoomID = req.RoomID
	response.Data.PlayerCount = 1
	response.Data.MaxPlayers = 4
	response.Data.Status = "waiting"
	response.Data.CreatedTime = time.Now()

	return s.Response(response)
}

// Leave 离开房间
func (r *RoomComponent) Leave(s *session.Session, req *dto.LeaveRoomRequest) error {
	// 检查会话中的character信息
	character := getCharacter(s)
	if character == nil {
		return s.Response(&dto.BaseResponse{
			Code:    401,
			Message: "未认证，请先登录",
		})
	}

	logger.Info("用户离开房间请求: %s (ID: %d) 房间: %s", character.Name, character.ID, req.RoomID)

	response := &dto.BaseResponse{
		Code:    200,
		Message: "离开房间成功",
	}

	return s.Response(response)
}

// GetRoomInfo 获取房间信息
func (r *RoomComponent) GetRoomInfo(s *session.Session, req *dto.GetRoomInfoRequest) error {
	// 检查会话中的character信息
	character := getCharacter(s)
	if character == nil {
		return s.Response(&dto.BaseResponse{
			Code:    401,
			Message: "未认证，请先登录",
		})
	}

	logger.Info("用户 %s (ID: %d) 获取房间信息请求: %s", character.Name, character.ID, req.RoomID)

	response := &dto.RoomInfoResponse{}
	response.Code = 200
	response.Message = "获取房间信息成功"
	response.Data.RoomID = req.RoomID
	response.Data.PlayerCount = 2
	response.Data.MaxPlayers = 4
	response.Data.Status = "playing"
	response.Data.CreatedTime = time.Now().Add(-10 * time.Minute)

	return s.Response(response)
}

// 使用 init 函数自动注册组件
func init() {
	// 确保logger被初始化
	logger.InitWithPath("data/logs")
	// 将RoomComponent注册到TCP路由器
	router.Register(&RoomComponent{})
}
