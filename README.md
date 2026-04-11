# FastGox API Starter

[English](README_EN.md) | 中文

一个简洁的 Go API 启动模板，分层架构，开箱即用。

## 项目结构

```
main.go                    # 入口
config/                    # 配置文件（YAML，按环境区分）
deploy/docker/             # Dockerfile + docker-compose
docs/
  db/                      # SQL 迁移文件
  swagger/                 # Swagger 文档
src/
  core/
    config/                # 配置加载
    database/              # 数据库连接、迁移、工具
    i18n/                  # 国际化（中/英）
    session/               # 会话管理
    tcp/                   # TCP 长连接服务
  models/
    config/                # 配置结构体
    dto/                   # 数据传输对象（请求/响应/错误码）
    entity/                # 数据库实体
  repository/              # 数据访问层
  services/                # 业务逻辑层
  router/
    handle/                # HTTP 路由处理器
    middleware/             # 中间件（认证、CORS、i18n、Recovery）
  utils/                   # 工具函数（JWT、AES、Hash 等）
  pkg/                     # 外部服务封装
test/                      # 单元测试
templates/                 # HTML 模板
```

## 特性

- **分层架构**：Repository → Service → Router，职责清晰
- **HTTP + TCP 双协议**：HTTP API 与 TCP 长连接服务并行运行
- **国际化（i18n）**：内置中/英双语支持，按请求头自动切换
- **统一响应格式**：标准化的成功/失败/分页响应结构
- **SQL 迁移版本管理**：基于 `schema_migrations` 表自动执行
- **Graceful Shutdown**：HTTP 与 TCP 服务均支持优雅关闭
- **Swagger 文档自动生成**
- **JWT 认证 + CORS + Recovery 中间件**
- **多环境配置**：dev / test / prod 配置隔离
- **Docker 一键部署**
- **Task 任务管理**（开发、测试、构建、部署一体化）
- **单元测试覆盖**：config、database、entity、i18n、tcp、utils 等

## 快速开始

### 环境要求

- Go 1.24+
- [Task](https://taskfile.dev)（推荐）
- Docker / Docker Compose（部署用）

### 开发

```bash
# 初始化项目（依赖 + Swagger）
task setup

# 启动开发服务器
task dev

# 格式化代码
task fmt

# 生成 Swagger 文档
task swagger
```

### 测试

```bash
# 运行所有单元测试
task test

# 测试并生成覆盖率报告
task test:cover

# 快速测试（跳过集成测试）
task test:short
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
APP_ENV=dev    # 加载 config/dev.yaml
APP_ENV=test   # 加载 config/test.yaml
APP_ENV=prod   # 加载 config/prod.yaml
```

主要配置项：

| 模块 | 说明 |
| --- | --- |
| `app` | 应用名称、端口、调试模式、Swagger、关闭超时 |
| `database` | 数据库驱动、连接参数、自动迁移 |
| `jwt` | JWT 密钥 |
| `sms-code` | 短信验证码长度、过期、频率限制、白名单 |
| `tcp` | TCP 服务端口、最大连接数、超时、心跳 |
| `file` | 文件上传大小限制、数量限制、允许类型 |

## 数据库迁移

在 `docs/db/` 下按编号添加 SQL 文件：

```
docs/db/000001_create_users.sql
docs/db/000002_create_sms_codes.sql
docs/db/000003_your_migration.sql
```

启动时自动执行未运行的迁移（需配置 `auto_migrate: true`）。

## 许可证

Apache 2.0 - 查看 [LICENSE](LICENSE)
