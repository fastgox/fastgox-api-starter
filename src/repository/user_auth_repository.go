package repository

import "github.com/fastgox/fastgox-api-starter/src/models/entity"

// UserAuthRepository 用户认证仓储
type UserAuthRepository struct {
	BaseRepository[entity.UserAuth]
}

var UserAuthRepo = &UserAuthRepository{BaseRepository: *GetRepository[entity.UserAuth]()}

// GetByUserID 根据用户ID获取认证信息
func (r *UserAuthRepository) GetByUserID(userID int64) (*entity.UserAuth, error) {
	return r.First("user_id = ?", userID)
}

// GetByUserIDAndAuthType 根据用户ID和认证类型获取认证信息
func (r *UserAuthRepository) GetByUserIDAndAuthType(userID int64, authType string) (*entity.UserAuth, error) {
	return r.First("user_id = ? AND auth_type = ?", userID, authType)
}

// GetByIdCardNumber 根据身份证号获取认证信息
func (r *UserAuthRepository) GetByIdCardNumber(idCardNumber string) (*entity.UserAuth, error) {
	return r.First("id_card_number = ?", idCardNumber)
}
