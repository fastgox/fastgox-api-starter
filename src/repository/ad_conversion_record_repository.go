package repository

import (
	"time"

	"github.com/fastgox/fastgox-api-starter/src/models/entity"
)

// AdConversionRecordRepository 广告转化记录仓储
type AdConversionRecordRepository struct {
	BaseRepository[entity.AdConversionRecord]
}

var AdConversionRecordRepo = &AdConversionRecordRepository{BaseRepository: *GetRepository[entity.AdConversionRecord]()}

// GetByAdIDAndDeviceID 根据广告ID和设备ID获取转化记录
func (r *AdConversionRecordRepository) GetByAdIDAndDeviceID(adID, deviceID string) (*entity.AdConversionRecord, error) {
	return r.First("ad_id = ? AND device_id = ?", adID, deviceID)
}

// GetByChannelCode 根据渠道代码获取转化记录列表
func (r *AdConversionRecordRepository) GetByChannelCode(channelCode string, limit int) ([]*entity.AdConversionRecord, error) {
	var records []*entity.AdConversionRecord
	err := r.DB.Where("channel_code = ?", channelCode).
		Order("created_at DESC").
		Limit(limit).
		Find(&records).Error
	return records, err
}

// GetByUserID 根据用户ID获取转化记录列表
func (r *AdConversionRecordRepository) GetByUserID(userID int64) ([]*entity.AdConversionRecord, error) {
	var records []*entity.AdConversionRecord
	err := r.DB.Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&records).Error
	return records, err
}

// CountByChannelCodeAndDateRange 按渠道代码和日期范围统计转化数量
func (r *AdConversionRecordRepository) CountByChannelCodeAndDateRange(channelCode string, startTime, endTime time.Time) (int64, error) {
	var count int64
	err := r.DB.Model(&entity.AdConversionRecord{}).
		Where("channel_code = ? AND created_at BETWEEN ? AND ?", channelCode, startTime, endTime).
		Count(&count).Error
	return count, err
}

// CheckRepeatConversion 检查是否为重复转化
func (r *AdConversionRecordRepository) CheckRepeatConversion(adID, deviceID, ip string, userID *int64) (bool, error) {
	var count int64
	query := r.DB.Model(&entity.AdConversionRecord{}).Where("ad_id = ?", adID)

	// 优先使用设备ID，如果设备ID为空则使用IP
	if deviceID != "" {
		query = query.Where("device_id = ?", deviceID)
	} else if ip != "" {
		query = query.Where("ip = ?", ip)
	}

	if userID != nil && *userID > 0 {
		query = query.Where("user_id = ?", *userID)
	}

	err := query.Count(&count).Error
	return count > 0, err
}

// CreateWithRepeatCheck 创建转化记录并检查重复
func (r *AdConversionRecordRepository) CreateWithRepeatCheck(record *entity.AdConversionRecord) error {
	// 获取设备ID和IP
	deviceID := ""
	if record.DeviceID != nil {
		deviceID = *record.DeviceID
	}

	ip := ""
	if record.IP != nil {
		ip = *record.IP
	}

	// 检查是否重复转化（优先使用设备ID，如果为空则使用IP）
	isRepeat, err := r.CheckRepeatConversion(record.AdID, deviceID, ip, record.UserID)
	if err != nil {
		return err
	}

	record.IsRepeatConvert = isRepeat
	now := time.Now()
	if record.ConversionTime == nil {
		record.ConversionTime = &now
	}

	return r.Create(record)
}

// CheckDeviceHasHistory 检查设备是否有历史转化记录
func (r *AdConversionRecordRepository) CheckDeviceHasHistory(adID, deviceID, ip string) (bool, error) {
	var count int64
	query := r.DB.Model(&entity.AdConversionRecord{}).Where("ad_id = ?", adID)

	// 优先使用设备ID，如果设备ID为空则使用IP
	if deviceID != "" {
		query = query.Where("device_id = ?", deviceID)
	} else if ip != "" {
		query = query.Where("ip = ?", ip)
	}

	err := query.Count(&count).Error
	return count > 0, err
}