package pkg_test

import (
	"fmt"
	"testing"

	"github.com/fastgox/fastgox-api-starter/src/pkg"
)

// MockProvider 用于测试的模拟 Provider
type MockProvider struct {
	name string
}

func (p *MockProvider) GetName() string {
	return p.name
}

func (p *MockProvider) Call(input string) (string, error) {
	return "result_from_" + p.name, nil
}

func TestManagerRegisterAndGet(t *testing.T) {
	manager := pkg.NewManager[pkg.Provider[string, string]]("default")

	provider1 := &MockProvider{name: "provider1"}
	provider2 := &MockProvider{name: "provider2"}

	manager.Register("provider1", provider1)
	manager.Register("provider2", provider2)

	// 通过名称获取
	p, err := manager.Get("provider1")
	if err != nil {
		t.Fatalf("获取 provider1 失败: %v", err)
	}
	if p.GetName() != "provider1" {
		t.Errorf("name = %q, want %q", p.GetName(), "provider1")
	}
}

func TestManagerGetDefault(t *testing.T) {
	manager := pkg.NewManager[pkg.Provider[string, string]]("default_provider")

	provider := &MockProvider{name: "default_provider"}
	manager.Register("default_provider", provider)

	// 空名称应返回默认
	p, err := manager.Get("")
	if err != nil {
		t.Fatalf("获取默认 provider 失败: %v", err)
	}
	if p.GetName() != "default_provider" {
		t.Errorf("默认 provider 名称不匹配: got=%q", p.GetName())
	}
}

func TestManagerGetFallback(t *testing.T) {
	manager := pkg.NewManager[pkg.Provider[string, string]]("default_provider")

	provider := &MockProvider{name: "default_provider"}
	manager.Register("default_provider", provider)

	// 请求不存在的 provider，应回退到默认
	p, err := manager.Get("nonexistent")
	if err != nil {
		t.Fatalf("回退到默认 provider 失败: %v", err)
	}
	if p.GetName() != "default_provider" {
		t.Errorf("应回退到默认 provider: got=%q", p.GetName())
	}
}

func TestManagerGetEmpty(t *testing.T) {
	manager := pkg.NewManager[pkg.Provider[string, string]]("default")

	// 没有注册任何 provider
	_, err := manager.Get("")
	if err == nil {
		t.Fatal("空管理器应返回错误")
	}
}

func TestManagerGetExact(t *testing.T) {
	manager := pkg.NewManager[pkg.Provider[string, string]]("default")

	provider := &MockProvider{name: "exact"}
	manager.Register("exact", provider)

	p, ok := manager.GetExact("exact")
	if !ok {
		t.Fatal("GetExact 应找到已注册的 provider")
	}
	if p.GetName() != "exact" {
		t.Errorf("name = %q, want %q", p.GetName(), "exact")
	}

	_, ok = manager.GetExact("nonexistent")
	if ok {
		t.Error("GetExact 对不存在的 key 应返回 false")
	}
}

func TestManagerGetNames(t *testing.T) {
	manager := pkg.NewManager[pkg.Provider[string, string]]("default")

	manager.Register("a", &MockProvider{name: "a"})
	manager.Register("b", &MockProvider{name: "b"})
	manager.Register("c", &MockProvider{name: "c"})

	names := manager.GetNames()
	if len(names) != 3 {
		t.Errorf("应有 3 个 provider: got=%d", len(names))
	}

	nameSet := make(map[string]bool)
	for _, n := range names {
		nameSet[n] = true
	}
	for _, expected := range []string{"a", "b", "c"} {
		if !nameSet[expected] {
			t.Errorf("缺少 provider: %q", expected)
		}
	}
}

func TestManagerCount(t *testing.T) {
	manager := pkg.NewManager[pkg.Provider[string, string]]("default")

	if manager.Count() != 0 {
		t.Errorf("初始 Count 应为 0: got=%d", manager.Count())
	}

	manager.Register("p1", &MockProvider{name: "p1"})
	if manager.Count() != 1 {
		t.Errorf("注册后 Count 应为 1: got=%d", manager.Count())
	}

	manager.Register("p2", &MockProvider{name: "p2"})
	if manager.Count() != 2 {
		t.Errorf("注册后 Count 应为 2: got=%d", manager.Count())
	}
}

func TestProviderCall(t *testing.T) {
	provider := &MockProvider{name: "test"}
	result, err := provider.Call("input")
	if err != nil {
		t.Fatalf("Call 失败: %v", err)
	}
	if result != "result_from_test" {
		t.Errorf("result = %q, want %q", result, "result_from_test")
	}
}

// BenchmarkManagerGet 基准测试
func BenchmarkManagerGet(b *testing.B) {
	manager := pkg.NewManager[pkg.Provider[string, string]]("default")
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("provider_%d", i)
		manager.Register(name, &MockProvider{name: name})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		manager.Get("provider_5")
	}
}
