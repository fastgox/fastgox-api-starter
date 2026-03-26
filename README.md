# FastGox API Starter

[English](README_EN.md) | 中文

一个简洁的 Go API 启动模板，分层架构，开箱即用。

## 项目结构

```
main.go                    # 入口
config/                    # 配置文件（YAML）
deploy/docker/             # Dockerfile + docker-compose
docs/
  db/                      # SQL 迁移文件
  swagger/                 # Swagger 文档
src/
  core/                    # 核心：配置加载、数据库、Session
  models/                  # 数据模型：entity、DTO、config
  repository/              # 数据访问层
  services/                # 业务逻辑层
  router/                  # 路由 + 中间件
  utils/                   # 工具函数
  pkg/                     # 外部服务封装
```

## 特性

- 分层架构：Repository -> Service -> Router
- 自动路由注册（基于 init()）
- SQL 迁移版本管理（schema_migrations）
- Graceful Shutdown（生产可用）
- Swagger 文档自动生成
- JWT 认证 + CORS 中间件
- Docker 一键部署
- Task 任务管理

## 快速开始

### 环境要求

- Go 1.24+
- [Task](https://taskfile.dev)（推荐）
- Docker / Docker Compose（部署用）

### 开发

```bash
# 初始化项目
task setup

# 启动开发服务器
task dev

# 格式化代码
task fmt

# 生成 Swagger 文档
task swagger
```

### 构建与发布

```bash
# 本地构建
task build

# 构建发布版本（clean + fmt + tidy + swagger + build）
task release
```

### Docker 部署

```bash
# 构建镜像
task docker-build

# 构建并启动（含 PostgreSQL）
task deploy

# 查看日志
task docker-logs

# 停止服务
task docker-down
```

### 所有可用任务

```bash
task --list
```

## 配置

配置文件位于 `config/` 目录，通过 `APP_ENV` 环境变量选择：

```bash
APP_ENV=dev   # 加载 config/dev.yaml
APP_ENV=prod  # 加载 config/prod.yaml
```

## 数据库迁移

在 `docs/db/` 下按编号添加 SQL 文件：

```
docs/db/000001_create_users.sql
docs/db/000002_create_sms_codes.sql
docs/db/000003_your_migration.sql
```

启动时自动执行未运行的迁移（需配置 `auto_migrate: true`）。

## 许可证

MIT - 查看 [LICENSE](LICENSE)
