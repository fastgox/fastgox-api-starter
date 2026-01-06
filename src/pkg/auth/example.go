package auth

import "github.com/fastgox/fastgox-api-starter/src/pkg"

// ExampleAuthProvider 示例认证服务提供商
type ExampleAuthProvider struct{}

// NewExampleAuthProvider 创建示例认证服务提供商
func NewExampleAuthProvider() *ExampleAuthProvider {
	return &ExampleAuthProvider{}
}

// GetName 获取提供商名称
func (p *ExampleAuthProvider) GetName() string {
	return ProviderNameExample
}

// Call 调用认证服务
func (p *ExampleAuthProvider) Call(input *AuthInput) (*AuthOutput, error) {
	// 示例实现，实际项目中替换为真实的认证逻辑
	return &AuthOutput{
		BaseOutput: pkg.BaseOutput{
			Success:  true,
			Provider: p.GetName(),
		},
		Verified: true,
		Message:  "认证成功",
	}, nil
}

func init() {
	// 注册示例认证服务提供商
	provider := NewExampleAuthProvider()
	AuthManager.Register(provider.GetName(), provider)
}
