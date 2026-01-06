package services

import (
	"errors"
	"time"

	"github.com/fastgox/fastgox-api-starter/src/models/dto/request"
	"github.com/fastgox/fastgox-api-starter/src/models/dto/response"
	"github.com/fastgox/fastgox-api-starter/src/models/entity"
	"github.com/fastgox/fastgox-api-starter/src/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

var AuthSvc = &AuthService{}

// Login 登录
func (s *AuthService) Login(req *request.LoginRequest) (*response.LoginResponse, error) {
	auth, err := repository.AuthRepo.GetActiveByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if auth == nil {
		return nil, errors.New("账号不存在或未激活")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("密码错误")
	}

	user, err := repository.UserRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if user != nil {
		_ = repository.UserRepo.UpdateLastActiveAt(user.ID)
	}

	// TODO: 生成JWT Token
	expiresAt := time.Now().Add(24 * time.Hour).Unix()

	return &response.LoginResponse{
		Token:     "token_placeholder",
		ExpiresAt: expiresAt,
		User:      toUserResponse(user),
	}, nil
}

func toUserResponse(user *entity.User) *response.UserResponse {
	if user == nil {
		return nil
	}
	return &response.UserResponse{
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		Role:            user.Role,
		ProfileImageURL: user.ProfileImageURL,
		Phone:           user.Phone,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		LastActiveAt:    user.LastActiveAt,
	}
}
