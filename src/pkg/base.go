package pkg

import (
	"fmt"
	"math/rand"
	"sync"
)

// Provider 通用服务提供商接口 - 支持泛型
type Provider[TInput any, TOutput any] interface {
	GetName() string                    // 获取提供商名称
	Call(input TInput) (TOutput, error) // 类型安全的调用方法
}

// Manager 服务提供商管理器
type Manager[T any] struct {
	providers   map[string]T
	defaultName string
	mutex       sync.RWMutex
}

// NewManager 创建新的服务提供商管理器
func NewManager[T any](defaultName string) *Manager[T] {
	return &Manager[T]{
		providers:   make(map[string]T),
		defaultName: defaultName,
	}
}

// Register 注册服务提供商
func (m *Manager[T]) Register(name string, provider T) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.providers[name] = provider
}

// Get 根据名称获取服务提供商，支持智能回退
func (m *Manager[T]) Get(name string) (T, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	var zero T

	// 如果没有指定名称，使用默认名称
	if name == "" {
		name = m.defaultName
	}

	// 尝试获取指定的提供商
	if provider, exists := m.providers[name]; exists {
		return provider, nil
	}

	// 如果指定的提供商不存在，且不是默认提供商，尝试获取默认提供商
	if name != m.defaultName {
		if provider, exists := m.providers[m.defaultName]; exists {
			return provider, nil
		}
	}

	// 如果默认提供商也不存在，尝试随机获取一个
	if len(m.providers) == 0 {
		return zero, fmt.Errorf("没有可用的服务提供商")
	}

	keys := make([]string, 0, len(m.providers))
	for k := range m.providers {
		keys = append(keys, k)
	}

	randomKey := keys[rand.Intn(len(keys))]
	return m.providers[randomKey], nil
}

// GetExact 精确获取指定名称的提供商
func (m *Manager[T]) GetExact(name string) (T, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	provider, exists := m.providers[name]
	return provider, exists
}

// GetNames 获取所有服务提供商名称
func (m *Manager[T]) GetNames() []string {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	names := make([]string, 0, len(m.providers))
	for name := range m.providers {
		names = append(names, name)
	}
	return names
}

// Count 获取已注册的服务提供商数量
func (m *Manager[T]) Count() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.providers)
}
