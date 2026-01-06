package repository

import "github.com/fastgox/fastgox-api-starter/src/models/entity"

// AuthRepository 认证仓储
type AuthRepository struct {
	*BaseRepository[entity.Auth]
}

var AuthRepo = &AuthRepository{}

func (r *AuthRepository) getRepo() *BaseRepository[entity.Auth] {
	if r.BaseRepository == nil {
		r.BaseRepository = GetRepository[entity.Auth]()
	}
	return r.BaseRepository
}

// GetByID 根据ID获取认证信息
func (r *AuthRepository) GetByID(id string) (*entity.Auth, error) {
	return r.getRepo().First("id = ?", id)
}

// GetByEmail 根据邮箱获取认证信息
func (r *AuthRepository) GetByEmail(email string) (*entity.Auth, error) {
	return r.getRepo().First("email = ?", email)
}

// GetByPhone 根据手机号获取认证信息
func (r *AuthRepository) GetByPhone(phone string) (*entity.Auth, error) {
	return r.getRepo().First("phone = ?", phone)
}

// ExistsByEmail 检查邮箱是否已注册
func (r *AuthRepository) ExistsByEmail(email string) (bool, error) {
	count, err := r.getRepo().Count("email = ?", email)
	return count > 0, err
}

// UpdatePassword 更新密码
func (r *AuthRepository) UpdatePassword(id string, hashedPassword string) error {
	return r.getRepo().DB.Model(&entity.Auth{}).
		Where("id = ?", id).
		Update("password", hashedPassword).Error
}

// UpdateActiveStatus 更新激活状态
func (r *AuthRepository) UpdateActiveStatus(id string, active bool) error {
	return r.getRepo().DB.Model(&entity.Auth{}).
		Where("id = ?", id).
		Update("active", active).Error
}

// GetActiveByEmail 获取已激活的认证信息
func (r *AuthRepository) GetActiveByEmail(email string) (*entity.Auth, error) {
	return r.getRepo().First("email = ? AND active = ?", email, true)
}
