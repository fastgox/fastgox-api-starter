package database

import "fmt"

// formatDSN 格式化数据库连接字符串
func formatDSN(host string, port int, user, password, dbname, sslmode, timezone string) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone,
	)
}
