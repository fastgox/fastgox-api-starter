package database

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/fastgox/utils/logger"
)

// MigrationRecord 迁移记录
type MigrationRecord struct {
	ID         int64     `json:"id"`
	Version    string    `json:"version"`
	Name       string    `json:"name"`
	ExecutedAt time.Time `json:"executed_at"`
	Checksum   string    `json:"checksum"`
	Success    bool      `json:"success"`
}

// Migrator 数据库迁移管理器
type Migrator struct {
	db            *sql.DB
	driver        string
	migrationsDir string
}

// NewMigrator 创建迁移管理器
func NewMigrator(db *sql.DB, driver string, migrationsDir string) *Migrator {
	return &Migrator{
		db:            db,
		driver:        driver,
		migrationsDir: migrationsDir,
	}
}

// RunMigrations 执行所有待迁移的 SQL 文件
func (m *Migrator) RunMigrations() error {
	// 1. 确保 schema_migrations 表存在
	if err := m.ensureMigrationsTable(); err != nil {
		return fmt.Errorf("创建迁移记录表失败: %w", err)
	}

	// 2. 扫描迁移目录
	files, err := m.scanMigrationFiles()
	if err != nil {
		return fmt.Errorf("扫描迁移文件失败: %w", err)
	}

	if len(files) == 0 {
		logger.Info("[Migration] 没有找到迁移文件")
		return nil
	}

	// 3. 获取已执行的迁移记录
	executed, err := m.getExecutedMigrations()
	if err != nil {
		return fmt.Errorf("获取已执行迁移记录失败: %w", err)
	}

	// 4. 找出待执行的迁移
	pending := m.findPendingMigrations(files, executed)

	logger.Info("[Migration] 发现 %d 个迁移文件，已执行 %d 个，待执行 %d 个",
		len(files), len(executed), len(pending))

	if len(pending) == 0 {
		logger.Info("[Migration] 所有迁移已是最新")
		return nil
	}

	// 5. 按顺序执行待迁移的 SQL
	for _, file := range pending {
		if err := m.executeMigration(file); err != nil {
			return fmt.Errorf("执行迁移 %s 失败: %w", file, err)
		}
	}

	logger.Info("[Migration] ✅ 所有迁移执行完毕")
	return nil
}

// ensureMigrationsTable 确保迁移记录表存在
func (m *Migrator) ensureMigrationsTable() error {
	var createSQL string

	switch m.driver {
	case "mysql":
		createSQL = `CREATE TABLE IF NOT EXISTS schema_migrations (
			id          BIGINT AUTO_INCREMENT PRIMARY KEY,
			version     VARCHAR(255) NOT NULL UNIQUE,
			name        VARCHAR(255) NOT NULL,
			executed_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			checksum    VARCHAR(64) NOT NULL,
			success     TINYINT(1) NOT NULL DEFAULT 1
		)`
	case "postgres", "postgresql":
		createSQL = `CREATE TABLE IF NOT EXISTS schema_migrations (
			id          BIGSERIAL PRIMARY KEY,
			version     VARCHAR(255) NOT NULL UNIQUE,
			name        VARCHAR(255) NOT NULL,
			executed_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			checksum    VARCHAR(64) NOT NULL,
			success     BOOLEAN NOT NULL DEFAULT TRUE
		)`
	default:
		return fmt.Errorf("不支持的数据库驱动: %s", m.driver)
	}

	_, err := m.db.Exec(createSQL)
	return err
}

// scanMigrationFiles 扫描迁移文件目录，返回按文件名排序的文件列表
func (m *Migrator) scanMigrationFiles() ([]string, error) {
	// 检查目录是否存在
	if _, err := os.Stat(m.migrationsDir); os.IsNotExist(err) {
		logger.Info("[Migration] 迁移目录不存在: %s，跳过迁移", m.migrationsDir)
		return nil, nil
	}

	entries, err := os.ReadDir(m.migrationsDir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.HasSuffix(entry.Name(), ".sql") {
			files = append(files, entry.Name())
		}
	}

	// 按文件名排序
	sort.Strings(files)
	return files, nil
}

// getExecutedMigrations 获取已执行的迁移记录
func (m *Migrator) getExecutedMigrations() (map[string]MigrationRecord, error) {
	records := make(map[string]MigrationRecord)

	rows, err := m.db.Query("SELECT version, name, checksum, success FROM schema_migrations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var record MigrationRecord
		if err := rows.Scan(&record.Version, &record.Name, &record.Checksum, &record.Success); err != nil {
			return nil, err
		}
		records[record.Name] = record
	}

	return records, rows.Err()
}

// findPendingMigrations 找出未执行的迁移文件
func (m *Migrator) findPendingMigrations(files []string, executed map[string]MigrationRecord) []string {
	var pending []string
	for _, file := range files {
		if _, ok := executed[file]; !ok {
			pending = append(pending, file)
		}
	}
	return pending
}

// executeMigration 执行单个迁移文件
func (m *Migrator) executeMigration(filename string) error {
	filePath := filepath.Join(m.migrationsDir, filename)

	// 读取 SQL 文件
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("读取文件失败: %w", err)
	}

	// 计算 checksum
	checksum := fmt.Sprintf("%x", sha256.Sum256(content))

	// 提取版本号（文件名下划线前的部分）
	version := extractVersion(filename)

	logger.Info("[Migration] 执行迁移: %s ...", filename)

	// 分割并执行 SQL 语句
	statements := splitSQL(string(content))
	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		if _, err := m.db.Exec(stmt); err != nil {
			// 记录失败
			m.recordMigration(version, filename, checksum, false)
			return fmt.Errorf("执行 SQL 失败: %w\nSQL: %s", err, stmt)
		}
	}

	// 记录为成功
	if err := m.recordMigration(version, filename, checksum, true); err != nil {
		return fmt.Errorf("记录迁移状态失败: %w", err)
	}

	logger.Info("[Migration] ✅ %s 执行成功", filename)
	return nil
}

// recordMigration 记录迁移执行结果
func (m *Migrator) recordMigration(version, name, checksum string, success bool) error {
	insertSQL := "INSERT INTO schema_migrations (version, name, checksum, success) VALUES (?, ?, ?, ?)"
	if m.driver == "postgres" || m.driver == "postgresql" {
		insertSQL = "INSERT INTO schema_migrations (version, name, checksum, success) VALUES ($1, $2, $3, $4)"
	}

	_, err := m.db.Exec(insertSQL, version, name, checksum, success)
	return err
}

// extractVersion 从文件名中提取版本号
// "000001_create_users.sql" → "000001"
func extractVersion(filename string) string {
	parts := strings.SplitN(filename, "_", 2)
	if len(parts) > 0 {
		return parts[0]
	}
	return filename
}

// splitSQL 将 SQL 内容按分号分割为多条语句
// 会忽略注释中的分号
func splitSQL(content string) []string {
	var statements []string
	var current strings.Builder
	inSingleLineComment := false
	inMultiLineComment := false

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// 跳过空行
		if trimmed == "" {
			continue
		}

		// 跳过单行注释
		if strings.HasPrefix(trimmed, "--") {
			continue
		}

		for i := 0; i < len(line); i++ {
			ch := line[i]

			// 处理多行注释
			if !inSingleLineComment && i+1 < len(line) && ch == '/' && line[i+1] == '*' {
				inMultiLineComment = true
				i++ // 跳过 *
				continue
			}
			if inMultiLineComment && i+1 < len(line) && ch == '*' && line[i+1] == '/' {
				inMultiLineComment = false
				i++ // 跳过 /
				continue
			}
			if inMultiLineComment {
				continue
			}

			// 处理行内 -- 注释
			if i+1 < len(line) && ch == '-' && line[i+1] == '-' {
				inSingleLineComment = true
				continue
			}

			if ch == ';' {
				stmt := strings.TrimSpace(current.String())
				if stmt != "" {
					statements = append(statements, stmt)
				}
				current.Reset()
			} else {
				current.WriteByte(ch)
			}
		}

		inSingleLineComment = false
		if current.Len() > 0 {
			current.WriteByte('\n')
		}
	}

	// 处理最后一条语句（可能没有分号结尾）
	stmt := strings.TrimSpace(current.String())
	if stmt != "" {
		statements = append(statements, stmt)
	}

	return statements
}
