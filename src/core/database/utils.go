package database

import (
	"fmt"
	"net/url"
)

// formatPostgresDSN 格式化PostgreSQL数据库连接字符串
func formatPostgresDSN(host string, port int, user, password, dbname, sslmode, timezone string) string {
	escapedPassword := url.QueryEscape(password)
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		host, user, escapedPassword, dbname, port, sslmode, timezone,
	)
}

func formatMySQLDSN(host string, port int, user, password, dbname, timezone string) string {
	escapedTimezone := url.QueryEscape(timezone)
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s&timeout=30s&allowNativePasswords=true&tls=false",
		user, password, host, port, dbname, escapedTimezone,
	)
}
