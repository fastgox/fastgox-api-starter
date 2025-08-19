package services

import (
	"github.com/fastgox/fastgox-api-starter/src/models/entity"
	"github.com/fastgox/fastgox-api-starter/src/repository"
)

// UserService 用户服务
type UserService struct {
	UserRepo *repository.UserRepository
}

var UserSvc *UserService = &UserService{UserRepo: repository.UserRepo}

// GetUserSvc 获取用户服务实例

// GetUserByEmail 根据邮箱获取用户
func (s *UserService) GetUserByEmail(email string) (*entity.User, error) {
	return s.UserRepo.GetByEmail(email)
}

// GetUserByID 根据ID获取用户
func (s *UserService) GetUserByID(id string) (*entity.User, error) {
	return s.UserRepo.GetByID(id)
}

// CreateUser 创建用户
func (s *UserService) CreateUser(user *entity.User) error {
	return s.UserRepo.Create(user)
}

// UpdateUser 更新用户
func (s *UserService) UpdateUser(user *entity.User) error {
	return s.UserRepo.Update(user)
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(id string) error {
	return s.UserRepo.Delete(id)
}

// IsEmailExists 检查邮箱是否已存在
func (s *UserService) IsEmailExists(email string) (bool, error) {
	user, err := s.UserRepo.GetByEmail(email)
	if err != nil {
		return false, err
	}
	return user != nil, nil
}
