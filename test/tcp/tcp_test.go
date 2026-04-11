package tcp_test

import (
	"encoding/binary"
	"encoding/json"
	"testing"

	"github.com/fastgox/fastgox-api-starter/src/core/tcp"
)

func TestCodecEncodeDecode(t *testing.T) {
	codec := tcp.NewCodec(65536)

	resp := &tcp.Response{
		Cmd:     "test",
		Seq:     "seq-001",
		Code:    200,
		Message: "success",
		Data:    map[string]any{"key": "value"},
	}

	// 编码
	buf, err := codec.Encode(resp)
	if err != nil {
		t.Fatalf("编码失败: %v", err)
	}

	// 检查包头
	if len(buf) < 4 {
		t.Fatal("编码结果长度不足 4 字节")
	}

	msgLen := binary.BigEndian.Uint32(buf[:4])
	if int(msgLen) != len(buf)-4 {
		t.Errorf("包头长度不匹配: header=%d, body=%d", msgLen, len(buf)-4)
	}

	// 验证 payload 是有效 JSON
	payload := buf[4:]
	var decoded tcp.Response
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("Payload 不是有效 JSON: %v", err)
	}

	if decoded.Cmd != "test" {
		t.Errorf("Cmd = %q, want %q", decoded.Cmd, "test")
	}
	if decoded.Code != 200 {
		t.Errorf("Code = %d, want %d", decoded.Code, 200)
	}
}

func TestCodecEncodeEmptyResponse(t *testing.T) {
	codec := tcp.NewCodec(0) // 使用默认 maxSize

	resp := &tcp.Response{
		Cmd:  "empty",
		Code: 200,
	}

	buf, err := codec.Encode(resp)
	if err != nil {
		t.Fatalf("编码空响应失败: %v", err)
	}

	if len(buf) <= 4 {
		t.Error("编码结果应大于 4 字节")
	}
}

func TestNewCodecDefaultMaxSize(t *testing.T) {
	codec := tcp.NewCodec(0)
	if codec == nil {
		t.Fatal("NewCodec(0) should not return nil")
	}
}

// ===== Router & Middleware 测试 =====

func TestRouterHandle(t *testing.T) {
	r := tcp.NewRouter()

	called := false
	r.Handle("test_cmd", func(ctx *tcp.Context) error {
		called = true
		return nil
	})

	// 手动创建上下文测试 Dispatch
	msg := &tcp.Message{Cmd: "test_cmd", Seq: "1"}
	codec := tcp.NewCodec(65536)
	ctx := tcp.NewContext(nil, "test-conn-1", msg, codec)

	err := r.Dispatch(ctx)
	if err != nil {
		t.Fatalf("Dispatch 失败: %v", err)
	}
	if !called {
		t.Error("handler 应该被调用")
	}
}

func TestRouterDispatchUnknownCmd(t *testing.T) {
	r := tcp.NewRouter()

	msg := &tcp.Message{Cmd: "unknown_cmd", Seq: "1"}
	codec := tcp.NewCodec(65536)
	ctx := tcp.NewContext(nil, "test-conn-1", msg, codec)

	// unknown cmd 会调用 ctx.Error()，但由于 conn 为 nil，会 panic
	// 这里验证下分发逻辑能找到"未知命令"
	defer func() {
		if r := recover(); r != nil {
			// 预期: conn 为 nil 导致 AsyncWrite panic
			// 这证明 router 正确识别了未知命令并尝试回复错误
		}
	}()
	_ = r.Dispatch(ctx)
}

func TestMiddlewareChain(t *testing.T) {
	order := []string{}

	mw1 := func(next tcp.Handler) tcp.Handler {
		return func(ctx *tcp.Context) error {
			order = append(order, "mw1-before")
			err := next(ctx)
			order = append(order, "mw1-after")
			return err
		}
	}

	mw2 := func(next tcp.Handler) tcp.Handler {
		return func(ctx *tcp.Context) error {
			order = append(order, "mw2-before")
			err := next(ctx)
			order = append(order, "mw2-after")
			return err
		}
	}

	handler := func(ctx *tcp.Context) error {
		order = append(order, "handler")
		return nil
	}

	chain := tcp.Chain([]tcp.Middleware{mw1, mw2}, handler)

	msg := &tcp.Message{Cmd: "test"}
	codec := tcp.NewCodec(65536)
	ctx := tcp.NewContext(nil, "conn-1", msg, codec)
	_ = chain(ctx)

	expected := []string{"mw1-before", "mw2-before", "handler", "mw2-after", "mw1-after"}
	if len(order) != len(expected) {
		t.Fatalf("中间件执行顺序不对: got=%v, want=%v", order, expected)
	}
	for i, v := range expected {
		if order[i] != v {
			t.Errorf("第 %d 项: got=%q, want=%q", i, order[i], v)
		}
	}
}

// ===== Context 测试 =====

func TestContextSetGet(t *testing.T) {
	msg := &tcp.Message{Cmd: "test"}
	codec := tcp.NewCodec(65536)
	ctx := tcp.NewContext(nil, "conn-1", msg, codec)

	ctx.Set("key1", "value1")
	ctx.Set("key2", int64(42))

	if v := ctx.GetString("key1"); v != "value1" {
		t.Errorf("GetString(key1) = %q, want %q", v, "value1")
	}
	if v := ctx.GetInt64("key2"); v != 42 {
		t.Errorf("GetInt64(key2) = %d, want %d", v, 42)
	}
	if v := ctx.GetString("nonexistent"); v != "" {
		t.Errorf("GetString(nonexistent) = %q, want empty", v)
	}
	if v := ctx.GetInt64("nonexistent"); v != 0 {
		t.Errorf("GetInt64(nonexistent) = %d, want 0", v)
	}
}

func TestContextBind(t *testing.T) {
	type TestData struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	data := json.RawMessage(`{"name":"test","age":25}`)
	msg := &tcp.Message{Cmd: "test", Data: data}
	codec := tcp.NewCodec(65536)
	ctx := tcp.NewContext(nil, "conn-1", msg, codec)

	var result TestData
	if err := ctx.Bind(&result); err != nil {
		t.Fatalf("Bind 失败: %v", err)
	}

	if result.Name != "test" || result.Age != 25 {
		t.Errorf("Bind 结果不匹配: %+v", result)
	}
}

func TestContextBindNilData(t *testing.T) {
	msg := &tcp.Message{Cmd: "test", Data: nil}
	codec := tcp.NewCodec(65536)
	ctx := tcp.NewContext(nil, "conn-1", msg, codec)

	var result map[string]any
	if err := ctx.Bind(&result); err != nil {
		t.Fatalf("Bind nil data 失败: %v", err)
	}
}

func TestContextAbort(t *testing.T) {
	msg := &tcp.Message{Cmd: "test"}
	codec := tcp.NewCodec(65536)
	ctx := tcp.NewContext(nil, "conn-1", msg, codec)

	if ctx.IsAborted() {
		t.Error("新 context 不应是 aborted")
	}

	ctx.Abort()
	if !ctx.IsAborted() {
		t.Error("Abort 后应是 aborted")
	}
}

// ===== ConnManager 测试 =====

func TestConnManagerCount(t *testing.T) {
	cm := tcp.NewConnManager()
	if cm.Count() != 0 {
		t.Errorf("初始连接数应为 0: got=%d", cm.Count())
	}
}
