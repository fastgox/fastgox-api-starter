package auth

import (
	"github.com/fastgox/fastgox-api-starter/src/pkg"
)

// 提供商名称常量
const (
	ProviderNameQiandun = "qiandun" // 钱盾提供商
	// 可以继续添加其他提供商...
)

// 默认提供商
const DefaultProvider = ProviderNameQiandun

// 全局身份认证服务提供商管理器
var AuthManager = pkg.NewManager[pkg.Provider[*ThreeElementsInput, *AuthOutput]](DefaultProvider)

// ThreeElementsInput 三要素验证输入
type ThreeElementsInput struct {
	pkg.BaseInput
	Name   string `json:"name" binding:"required"`   // 姓名
	IdCard string `json:"idCard" binding:"required"` // 身份证号码
	Mobile string `json:"mobile" binding:"required"` // 手机号
}

// AuthOutput 身份认证响应
type AuthOutput struct {
	pkg.BaseOutput
	VerifyStatus string `json:"verifyStatus"` // 认证状态：passed-通过，failed-未通过
	FlowNo       string `json:"flowNo"`       // 认证流水号
	ServiceId    string `json:"serviceId"`    // 本次服务流程ID
	LogKey       string `json:"logKey"`       // 日志键
}

// AuthResult 认证结果详情
type AuthResult struct {
	// 基础信息
	Success   bool   `json:"success"`   // 认证是否成功
	Message   string `json:"message"`   // 结果消息
	RequestId string `json:"requestId"` // 请求ID

	// 认证详情
	NameMatch   *bool `json:"nameMatch,omitempty"`   // 姓名是否匹配
	IdCardMatch *bool `json:"idCardMatch,omitempty"` // 身份证是否匹配
	MobileMatch *bool `json:"mobileMatch,omitempty"` // 手机号是否匹配
	BankMatch   *bool `json:"bankMatch,omitempty"`   // 银行卡是否匹配

	// 扩展信息
	Gender   string `json:"gender,omitempty"`   // 性别
	Birthday string `json:"birthday,omitempty"` // 生日
	Address  string `json:"address,omitempty"`  // 地址
	BankName string `json:"bankName,omitempty"` // 银行名称
	CardType string `json:"cardType,omitempty"` // 卡类型
}
