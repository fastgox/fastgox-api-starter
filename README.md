# FastGox API Starter

一个简洁的 Go API 启动模板，采用分层架构设计，专为快速开发而优化。

## 项目结构

```
├── main.go                # 开发工具脚本
├── cmd/
│   └── server/
│       └── main.go        # 服务器入口（包含 Swagger 注释）
├── src/
│   ├── server.go          # 服务器实现
│   ├── config/            # 配置管理
│   ├── core/              # 核心功能
│   ├── models/            # 数据模型
│   ├── repository/        # 数据访问层
│   ├── router/            # 路由层
│   │   ├── handle/        # 路由处理器
│   │   └── middleware/    # 中间件
│   └── services/          # 业务服务层
├── scripts/               # 构建脚本
└── docs/                  # 文档
```

## 特性

- ✅ **自动路由注册** - 使用 init() 函数自动注册路由
- ✅ **分层架构** - Repository → Service → Router
- ✅ **中间件支持** - CORS、认证等中间件
- ✅ **配置管理** - YAML 配置文件支持
- ✅ **简洁设计** - 遵循 Go 语言简洁哲学
- ✅ **易于扩展** - 添加新功能只需几步

## 快速开始

### 1. 安装依赖
```bash
go mod tidy
```

### 2. 安装go-task（可选，推荐）
```bash
go install github.com/go-task/task/v3/cmd/task@latest
```

### 3. 启动开发服务器
```bash
# 使用go-task（推荐）
task dev

# 或使用传统方式
go run main.go dev
```

### 4. 构建项目
```bash
# 使用go-task
task build

# 或使用传统方式
go run main.go build
```

### 5. 运行API测试
```bash
# 启动服务器
task dev

# 在另一个终端运行测试
cd test
task test-simple
```

## 设计理念

- **显式优于隐式** - 所有依赖关系清晰可见
- **简单直接** - 不使用复杂的自动注入或反射
- **自动化注册** - 利用 Go 的 init() 机制自动注册组件
- **易于理解** - 新手也能快速上手

## 路由初始化顺序

项目采用基于包导入的自动初始化机制：

1. `router` 包的 `init()` 创建路由组
2. `router/handle` 包的 `init()` 注册具体路由
3. 服务器启动时使用已配置好的路由引擎

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。