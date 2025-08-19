package database

import "errors"

// 数据库相关错误定义
var (
	ErrInvalidHost   = errors.New("数据库主机地址不能为空")
	ErrInvalidPort   = errors.New("数据库端口必须在1-65535之间")
	ErrInvalidUser   = errors.New("数据库用户名不能为空")
	ErrInvalidDBName = errors.New("数据库名称不能为空")

	ErrConnectionFailed  = errors.New("数据库连接失败")
	ErrNotInitialized    = errors.New("数据库管理器未初始化")
	ErrClientNotFound    = errors.New("数据库客户端未找到")
	ErrMigrationFailed   = errors.New("数据库迁移失败")
	ErrTransactionFailed = errors.New("数据库事务失败")
)
