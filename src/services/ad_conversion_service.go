package services

import (
	"time"

	"github.com/fastgox/fastgox-api-starter/src/core/session"
	"github.com/fastgox/fastgox-api-starter/src/models/dto/request"
	"github.com/fastgox/fastgox-api-starter/src/models/dto/response"
	"github.com/fastgox/fastgox-api-starter/src/models/entity"
	"github.com/fastgox/fastgox-api-starter/src/repository"
	"github.com/fastgox/fastgox-api-starter/src/utils"
	"github.com/gin-gonic/gin"
)

// AdConversionService 广告转化记录服务
type AdConversionService struct {
}

var AdConversionSvc = &AdConversionService{}

// CreateAdConversion 创建广告转化记录
func (s *AdConversionService) CreateAdConversion(c *gin.Context, req *request.CreateAdConversionRequest) (*response.AdConversionResponse, error) {
	// 构建实体对象
	now := time.Now()

	userSession, err := session.Manager.GetUserSessionAndByAuthorization(c)
	var userID *int64
	// 如果用户已登录，记录UserID；如果未登录，UserID为nil（这是正常的）
	if err == nil && userSession != nil {
		userID = &userSession.UserID
	}

	record := &entity.AdConversionRecord{
		AdID:           req.AdID,
		ChannelCode:    req.ChannelCode,
		UserID:         userID, // 安全处理：已登录时记录UserID，未登录时为nil
		ConversionType: req.ConversionType,
		ConversionTime: &now,
	}

	// 设置可选字段
	if req.Platform != "" {
		record.Platform = &req.Platform
	}
	if req.DeviceID != "" {
		record.DeviceID = &req.DeviceID
	}
	ip := utils.GetRealIP(c)
	record.IP = &ip

	if req.Medium != "" {
		record.Medium = &req.Medium
	}

	// 检查设备是否有历史记录（在创建新记录之前）
	deviceID := ""
	if record.DeviceID != nil {
		deviceID = *record.DeviceID
	}
	ipForCheck := ""
	if record.IP != nil {
		ipForCheck = *record.IP
	}

	deviceHasHistory, err := repository.AdConversionRecordRepo.CheckDeviceHasHistory(record.AdID, deviceID, ipForCheck)
	if err != nil {
		return nil, err
	}

	// 使用repository创建记录（包含重复检查）
	err = repository.AdConversionRecordRepo.CreateWithRepeatCheck(record)
	if err != nil {
		return nil, err
	}

	// 构建响应对象
	resp := &response.AdConversionResponse{
		ID:                 record.ID,
		AdID:               record.AdID,
		ChannelCode:        record.ChannelCode,
		ChannelName:        record.GetChannelName(),
		UserID:             record.UserID,
		ConversionType:     record.ConversionType,
		ConversionTypeName: record.GetConversionTypeName(),
		IsRepeatConvert:    record.IsRepeatConvert,
		Platform:           record.Platform,
		DeviceID:           record.DeviceID,
		IP:                 record.IP,
		ConversionTime:     record.ConversionTime,
		Medium:             record.Medium,
		CreatedAt:          record.CreatedAt,
		DeviceHasHistory:   deviceHasHistory, // 返回设备历史状态
	}

	return resp, nil
}
