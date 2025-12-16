package services

import (
	"errors"

	"github.com/fastgox/fastgox-api-starter/src/models/entity"
	"github.com/fastgox/fastgox-api-starter/src/repository"
)

// BlacklistService 黑名单服务
type BlacklistService struct{}

var Blacklist = &BlacklistService{}

// CheckResult 检查结果
type CheckResult struct {
	IsBlocked bool                     `json:"is_blocked"`        // 是否被拦截
	User      *entity.BlacklistUser    `json:"user,omitempty"`    // 黑名单用户
	Records   []entity.BlacklistRecord `json:"records,omitempty"` // 有效的拉黑记录
	HitField  string                   `json:"hit_field"`         // 命中字段
}

// Check 检查三要素是否在黑名单
func (s *BlacklistService) Check(name, idCard, phone string) (*CheckResult, error) {
	result := &CheckResult{IsBlocked: false}

	var user *entity.BlacklistUser

	// 优先检查身份证
	if idCard != "" {
		user, _ = repository.BlacklistUserRepo.CheckByIDCard(idCard)
		if user != nil {
			result.HitField = "id_card"
		}
	}

	// 检查手机号
	if user == nil && phone != "" {
		user, _ = repository.BlacklistUserRepo.CheckByPhone(phone)
		if user != nil {
			result.HitField = "phone"
		}
	}

	if user == nil {
		return result, nil
	}

	// 获取拉黑记录
	records, _ := repository.BlacklistRecordRepo.GetRecords(user.ID)

	result.IsBlocked = true
	result.User = user
	result.Records = records

	return result, nil
}

// AddBlacklistRequest 添加黑名单请求
type AddBlacklistRequest struct {
	Name     string `json:"name" binding:"required"`    // 姓名
	IDCard   string `json:"id_card" binding:"required"` // 身份证号
	Phone    string `json:"phone" binding:"required"`   // 手机号
	Reason   string `json:"reason"`                     // 原因
	Source   string `json:"source"`                     // 来源
	Operator string `json:"operator"`                   // 操作人
	Remark   string `json:"remark"`                     // 备注
}

// Add 添加黑名单（新增一条拉黑记录）
func (s *BlacklistService) Add(req *AddBlacklistRequest) error {
	if req.IDCard == "" {
		return errors.New("身份证号不能为空")
	}

	// 查找或创建黑名单用户
	user, _, err := repository.BlacklistUserRepo.FindOrCreate(req.Name, req.IDCard, req.Phone)
	if err != nil {
		return err
	}

	// 创建拉黑记录
	record := &entity.BlacklistRecord{
		UserID:   user.ID,
		Reason:   req.Reason,
		Source:   req.Source,
		Operator: req.Operator,
		Remark:   req.Remark,
	}

	return repository.BlacklistRecordRepo.AddRecord(record)
}

// Remove 解除黑名单（更新用户状态）
func (s *BlacklistService) Remove(userID int64) error {
	return repository.BlacklistUserRepo.UpdateStatus(userID, 0)
}

// RemoveByIDCard 根据身份证解除黑名单
func (s *BlacklistService) RemoveByIDCard(idCard string) error {
	user, err := repository.BlacklistUserRepo.GetByIDCard(idCard)
	if err != nil || user == nil {
		return errors.New("用户不存在")
	}
	return s.Remove(user.ID)
}

// GetUserWithRecords 获取用户及其所有拉黑记录
type BlacklistUserDetail struct {
	User    *entity.BlacklistUser    `json:"user"`
	Records []entity.BlacklistRecord `json:"records"`
}

func (s *BlacklistService) GetUserWithRecords(userID int64) (*BlacklistUserDetail, error) {
	user, err := repository.BlacklistUserRepo.GetByID(userID)
	if err != nil || user == nil {
		return nil, errors.New("用户不存在")
	}

	records, _ := repository.BlacklistRecordRepo.GetRecords(userID)

	return &BlacklistUserDetail{
		User:    user,
		Records: records,
	}, nil
}

// List 分页查询黑名单用户
func (s *BlacklistService) List(page, size int, keyword string, status int8) ([]entity.BlacklistUser, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}
	return repository.BlacklistUserRepo.List(page, size, keyword, status)
}

// BatchAdd 批量添加黑名单
func (s *BlacklistService) BatchAdd(items []AddBlacklistRequest) (int, []string) {
	successCount := 0
	var errs []string

	for _, item := range items {
		if err := s.Add(&item); err != nil {
			errs = append(errs, item.IDCard+": "+err.Error())
		} else {
			successCount++
		}
	}

	return successCount, errs
}
