package auth

import "github.com/fastgox/fastgox-api-starter/src/pkg"

// 提供商名称常量
const (
	ProviderNameExample = "example"
)

// 默认提供商
const DefaultProvider = ProviderNameExample

// 全局身份认证服务提供商管理器
var AuthManager = pkg.NewManager[pkg.Provider[*AuthInput, *AuthOutput]](DefaultProvider)

// AuthInput 认证输入
type AuthInput struct {
	pkg.BaseInput
	Name   string `json:"name" binding:"required"`
	IdCard string `json:"id_card" binding:"required"`
}

// AuthOutput 认证输出
type AuthOutput struct {
	pkg.BaseOutput
	Verified bool   `json:"verified"`
	Message  string `json:"message"`
}
