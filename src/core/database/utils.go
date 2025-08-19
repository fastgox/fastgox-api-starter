package database

import "fmt"

// formatPostgresDSN 格式化PostgreSQL数据库连接字符串
func formatPostgresDSN(host string, port int, user, password, dbname, sslmode, timezone string) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone,
	)
}

// formatMySQLDSN 格式化MySQL数据库连接字符串
func formatMySQLDSN(host string, port int, user, password, dbname, timezone string) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		user, password, host, port, dbname, timezone,
	)
}
