# HKChat API - 智能聊天后端服务

HKChat API 是一个基于 Go 语言开发的简洁高效的智能聊天后端服务，支持 OpenAI 模型集成和实时聊天功能。

## 📋 快速导航

- [🏗️ 系统架构文档](./docs/architecture/README.md)
- [🤖 LLM 集成架构](./docs/architecture/llm-integration.md)
- [📖 API 文档](./docs/swagger.yaml)
- [🚀 开发脚本说明](./scripts/README.md)

## 🚀 项目概述

HKChat API 是一个功能完整的智能聊天后端服务，集成了 RAG (检索增强生成) 系统和向量数据库。

### 核心特性

- ✅ **OpenAI 集成**: 支持 GPT 模型的聊天完成
- ✅ **RAG 系统**: 文档上传、向量化、智能检索
- ✅ **Milvus 向量数据库**: 高性能向量存储和相似度搜索
- ✅ **知识库管理**: 文档解析、分块、向量化存储
- ✅ **简洁架构**: 分层架构设计，易于理解和维护
- ✅ **配置管理**: 环境变量和 YAML 配置支持
- ✅ **数据持久化**: PostgreSQL + Redis + Milvus
- ✅ **API 文档**: Swagger 自动生成的 API 文档
- ✅ **一键部署**: Windows 批处理脚本快速启动

## 🏗️ 架构概览

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   客户端层      │───▶│   API 路由层     │───▶│   中间件层      │
│ Web/Mobile App  │    │  Gin Router      │    │  Auth/CORS      │
└─────────────────┘    └──────────────────┘    └─────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                        业务服务层                                 │
│           ChatService │ LLM Gateway │ OpenAI Client              │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                      数据访问层 (Repository)                      │
│              ChatRepo │ UserRepo │ ModelRepo                    │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                        数据存储层                                 │
│                       PostgreSQL                                │
└─────────────────────────────────────────────────────────────────┘
```

## 📁 项目结构

```
hkchat_api/
├── api/                    # API 路由层
│   ├── chat_api_routes.go  # 聊天 API 路由
│   ├── middleware/         # 认证中间件
│   └── registry.go         # 路由注册器
├── cmd/server/             # 服务器启动入口
├── config/                 # 配置文件
│   ├── dev.yaml           # 开发环境配置
│   └── prod.yaml          # 生产环境配置
├── docs/                   # 文档和API规范
├── internal/              # 内部模块
│   ├── config/            # 配置管理
│   ├── core/llm/          # LLM 网关和集成
│   ├── models/            # 数据模型
│   │   ├── entities/      # 数据库实体
│   │   ├── dto/           # 数据传输对象
│   │   └── domain/        # 领域模型
│   ├── repository/        # 数据访问层
│   └── services/          # 业务服务层
├── pkg/                   # 公共包
│   ├── auth/              # JWT认证
│   ├── clients/           # 外部客户端 (OpenAI, 数据库)
│   ├── prompts/           # 提示词模板
│   └── utils/             # 工具函数
├── scripts/               # 开发脚本
├── go.mod                 # 依赖管理
└── main.go               # 开发工具入口
```

## 🛠️ 技术栈

- **Web 框架**: [Gin](https://github.com/gin-gonic/gin) - 高性能 HTTP Web 框架
- **ORM**: [GORM](https://gorm.io/) - Go 语言 ORM 库  
- **关系数据库**: [PostgreSQL](https://www.postgresql.org/) - 主数据存储
- **向量数据库**: [Milvus](https://milvus.io/) - 高性能向量存储和检索
- **缓存**: [Redis](https://redis.io/) - 内存缓存数据库
- **LLM 集成**: [OpenAI API](https://openai.com/api/) - GPT 模型和 Embeddings
- **RAG 系统**: 文档解析、文本分块、向量化、相似度搜索
- **文档**: [Swagger](https://swagger.io/) - API 文档生成
- **容器化**: [Docker](https://www.docker.com/) - 服务容器化部署

## 🚀 快速开始

### 环境要求

- Go 1.21+
- Docker Desktop (用于运行数据库和向量数据库)
- OpenAI API Key

### Windows 一键启动 (推荐)

1. **克隆项目**
```bash
git clone <repository-url>
cd hkchat_api_go
```

2. **一键启动所有服务**
```bash
# 运行启动脚本 (会自动启动 Milvus + PostgreSQL + Redis)
scripts\start-hkchat.bat
```

3. **启动主应用**
```bash
# 开发模式
go run main.go

# 或构建后运行
go build -o hkchat_api.exe
hkchat_api.exe
```

4. **访问服务**
- API 服务: http://localhost:8080
- Swagger 文档: http://localhost:8080/swagger/index.html
- 健康检查: http://localhost:8080/health

### 手动启动 (Linux/Mac)

1. **克隆项目**
```bash
git clone <repository-url>
cd hkchat_api_go
```

2. **配置环境变量**
```bash
cp env.example .env
# 编辑 .env 文件，设置 OPENAI_API_KEY
```

3. **启动服务依赖**
```bash
# 启动 Milvus 向量数据库
curl -sfL https://raw.githubusercontent.com/milvus-io/milvus/master/scripts/standalone_embed.sh -o standalone_embed.sh
bash standalone_embed.sh start

# 启动 PostgreSQL 和 Redis
cd deploy/docker
docker compose -f docker-compose-services.yml up -d
cd ../..
```

4. **运行应用**
```bash
go run main.go
```

### 停止服务

```bash
# Windows
scripts\stop-hkchat.bat

# Linux/Mac
bash standalone_embed.sh stop
cd deploy/docker && docker compose -f docker-compose-services.yml down
```

## ⚙️ 配置管理

### 环境变量

| 变量名 | 说明 | 默认值 |
|--------|------|--------|
| `DB_HOST` | 数据库主机 | localhost |
| `DB_PORT` | 数据库端口 | 5432 |
| `DB_USER` | 数据库用户名 | postgres |
| `DB_PASSWORD` | 数据库密码 | - |
| `DB_NAME` | 数据库名称 | hkchat |
| `OPENAI_API_KEY` | OpenAI API 密钥 | - |
| `APP_PORT` | 服务端口 | 8080 |

### 配置文件

项目支持多环境配置，配置文件位于 `config/` 目录：

- `dev.yaml`: 开发环境配置
- `prod.yaml`: 生产环境配置

## 📖 开发指南

### 代码结构

项目采用简洁的三层架构：

1. **API 层** (`api/`): HTTP 路由和中间件
2. **服务层** (`internal/services/`): 业务逻辑处理  
3. **仓储层** (`internal/repository/`): 数据访问层
4. **模型层** (`internal/models/`): 数据模型定义

### 添加新功能

1. 在 `internal/models/entities/` 定义数据实体
2. 在 `internal/repository/` 实现数据访问逻辑
3. 在 `internal/services/` 实现业务逻辑
4. 在 `api/` 添加 HTTP 路由

### 开发命令

```bash
# 启动开发服务器
go run main.go run

# 格式化代码
go run main.go fmt

# 构建项目
go run main.go build

# 清理构建文件
go run main.go clean
```

## 🔗 相关链接

- [系统架构文档](./docs/architecture/README.md)
- [LLM 集成文档](./docs/architecture/llm-integration.md)
- [API 文档](./docs/swagger.yaml)
- [开发脚本说明](./scripts/README.md)

## 📄 许可证

本项目采用 MIT 许可证。

## 🤝 贡献

欢迎提交 Issue 和 Pull Request 来改进项目！
