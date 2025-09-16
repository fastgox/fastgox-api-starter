package handle

import (
	"fmt"
	"time"

	"github.com/fastgox/fastgox-api-starter/src/models/dto"
	"github.com/fastgox/fastgox-api-starter/src/models/entity"
	"github.com/fastgox/fastgox-api-starter/src/router"
	"github.com/fastgox/utils/logger"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
)

// AuthComponent 认证组件
type AuthComponent struct {
	component.Base
}

// validateToken 验证Token并返回用户对象
func (a *AuthComponent) validateToken(token string) (*entity.User, error) {
	// 这里应该实现真正的Token验证逻辑
	// 暂时返回一个模拟的用户对象
	if token == "" {
		return nil, fmt.Errorf("token不能为空")
	}

	// 模拟用户对象
	user := &entity.User{
		BaseModel: entity.BaseModel{
			ID:        "user_12345",
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
		Name:  "测试用户",
		Email: "test@example.com",
		Role:  "user",
	}

	return user, nil
}

// Login 用户登录认证
func (a *AuthComponent) Login(s *session.Session, req *dto.LoginRequest) error {
	logger.Info("TCP用户登录请求: %s", req.Token)
	fmt.Printf("🔥 [AUTH] TCP用户登录请求: %s\n", req.Token)

	// 验证Token（这里可以调用JWT验证或查询数据库）
	user, err := a.validateToken(req.Token)
	if err != nil {
		return s.Response(&dto.BaseResponse{
			Code:    401,
			Message: "认证失败: " + err.Error(),
		})
	}

	characterSession := &dto.CharacterSession{
		ID:   int64(123),
		Name: user.Name,
	}
	setCharacter(s, characterSession)

	// 响应登录成功
	response := &dto.LoginResponse{}
	response.Code = 200
	response.Message = "TCP认证成功"
	response.Data.UserID = 123
	response.Data.LoginTime = time.Now()

	logger.Info("用户TCP认证成功: %s (ID: %s)", user.Name, user.ID)
	return s.Response(response)
}

// HeartBeat 心跳检测（需要认证）
func (a *AuthComponent) HeartBeat(s *session.Session, req *dto.HeartBeatRequest) error {
	// 检查会话中的character信息
	character := getCharacter(s)
	if character == nil {
		return s.Response(&dto.BaseResponse{
			Code:    401,
			Message: "未认证，请先登录",
		})
	}

	logger.Info("心跳检测: 用户 %s (ID: %d)", character.Name, character.ID)

	response := &dto.HeartBeatResponse{}
	response.Code = 200
	response.Message = "heartbeat"
	response.Data.Timestamp = time.Now()
	response.Data.UserID = character.ID

	return s.Response(response)
}

// 使用 init 函数自动注册组件
func init() {
	// 确保logger被初始化
	logger.InitWithPath("data/logs")
	router.Register(&AuthComponent{})
}
