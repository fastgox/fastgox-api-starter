package response

// AuthVerifyResponse 身份认证响应
type AuthVerifyResponse struct {
	Success      bool   `json:"success"`      // 是否认证成功
	VerifyStatus string `json:"verifyStatus"` // 认证状态
	FlowNo       string `json:"flowNo"`       // 认证流水号
	ServiceId    string `json:"serviceId"`    // 服务流程ID
	Message      string `json:"message"`      // 响应消息
	Provider     string `json:"provider"`     // 服务提供商
	CostTime     int64  `json:"costTime"`     // 耗时（毫秒）
}
