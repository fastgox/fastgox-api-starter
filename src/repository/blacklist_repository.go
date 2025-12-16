package repository

import (
	"time"

	"github.com/fastgox/fastgox-api-starter/src/models/entity"
)

// BlacklistUserRepository 黑名单用户仓储
type BlacklistUserRepository struct {
	BaseRepository[entity.BlacklistUser]
}

// BlacklistRecordRepository 黑名单记录仓储
type BlacklistRecordRepository struct {
	BaseRepository[entity.BlacklistRecord]
}

var (
	BlacklistUserRepo   = &BlacklistUserRepository{BaseRepository: *GetRepository[entity.BlacklistUser]()}
	BlacklistRecordRepo = &BlacklistRecordRepository{BaseRepository: *GetRepository[entity.BlacklistRecord]()}
)

// ========== 用户仓储方法 ==========

// GetByIDCard 根据身份证获取用户
func (r *BlacklistUserRepository) GetByIDCard(idCard string) (*entity.BlacklistUser, error) {
	return r.First("id_card = ?", idCard)
}

// GetByPhone 根据手机号获取用户
func (r *BlacklistUserRepository) GetByPhone(phone string) (*entity.BlacklistUser, error) {
	return r.First("phone = ? AND status = 1", phone)
}

// CheckByIDCard 检查身份证是否在黑名单
func (r *BlacklistUserRepository) CheckByIDCard(idCard string) (*entity.BlacklistUser, error) {
	return r.First("id_card = ? AND status = 1", idCard)
}

// CheckByPhone 检查手机号是否在黑名单
func (r *BlacklistUserRepository) CheckByPhone(phone string) (*entity.BlacklistUser, error) {
	return r.First("phone = ? AND status = 1", phone)
}

// FindOrCreate 查找或创建黑名单用户
func (r *BlacklistUserRepository) FindOrCreate(name, idCard, phone string) (*entity.BlacklistUser, bool, error) {
	user, err := r.GetByIDCard(idCard)
	if err == nil && user != nil {
		// 已存在，更新状态为黑名单中
		if user.Status != 1 {
			now := time.Now()
			r.DB.Model(user).Updates(map[string]interface{}{
				"status":      1,
				"update_time": now,
			})
			user.Status = 1
		}
		return user, false, nil
	}

	// 创建新用户
	now := time.Now()
	newUser := &entity.BlacklistUser{
		Name:       name,
		IDCard:     idCard,
		Phone:      phone,
		Status:     1,
		CreateTime: &now,
		UpdateTime: &now,
	}
	if err := r.Create(newUser); err != nil {
		return nil, false, err
	}
	return newUser, true, nil
}

// UpdateStatus 更新用户状态
func (r *BlacklistUserRepository) UpdateStatus(id int64, status int8) error {
	now := time.Now()
	return r.DB.Model(&entity.BlacklistUser{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":      status,
			"update_time": now,
		}).Error
}

// List 分页查询
func (r *BlacklistUserRepository) List(page, size int, keyword string, status int8) ([]entity.BlacklistUser, int64, error) {
	if keyword != "" {
		return r.Page(page, size, "status = ? AND (name LIKE ? OR id_card LIKE ? OR phone LIKE ?)",
			status, "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}
	return r.Page(page, size, "status = ?", status)
}

// ========== 记录仓储方法 ==========

// AddRecord 添加拉黑记录
func (r *BlacklistRecordRepository) AddRecord(record *entity.BlacklistRecord) error {
	now := time.Now()
	record.CreateTime = &now
	return r.Create(record)
}

// GetRecords 获取用户所有拉黑记录
func (r *BlacklistRecordRepository) GetRecords(userID int64) ([]entity.BlacklistRecord, error) {
	var records []entity.BlacklistRecord
	err := r.DB.Where("user_id = ?", userID).
		Order("create_time DESC").
		Find(&records).Error
	return records, err
}
