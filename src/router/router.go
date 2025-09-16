package router

import (
	"github.com/lonng/nano/component"
)

var (
	// Components TCP组件管理器
	Components *component.Components
)

// init 包初始化时创建组件管理器
func init() {
	Components = &component.Components{}
}

// GetComponents 获取组件管理器（供外部使用）
func GetComponents() *component.Components {
	return Components
}

// RegisterComponent 注册TCP组件
func Register(comp component.Component) {
	Components.Register(comp)
}
