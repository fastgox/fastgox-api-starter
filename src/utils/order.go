package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

var (
	// 全局序列号计数器，确保在同一纳秒内也不会重复
	sequenceCounter int64
)

// generateOrderNo 生成订单号 - 基本不会重复的算法
func GenerateOrderNo() string {
	now := time.Now()

	// 1. 使用纳秒时间戳（更精确）
	nanoTime := now.UnixNano()

	// 2. 原子递增序列号（防止同一纳秒内重复）
	sequence := atomic.AddInt64(&sequenceCounter, 1)

	// 3. 添加进程ID增加唯一性
	pid := os.Getpid()

	// 4. 添加4字节加密随机数
	randomBytes := make([]byte, 2)
	rand.Read(randomBytes)
	randomHex := hex.EncodeToString(randomBytes)

	// 格式: ORD + 纳秒时间戳后10位 + 进程ID后4位 + 序列号后4位 + 随机数4位
	// 示例: ORD1234567890123412340001ABCD
	return fmt.Sprintf("ORD%010d%04d%04d%s",
		nanoTime%10000000000, // 纳秒时间戳后10位
		pid%10000,            // 进程ID后4位
		sequence%10000,       // 序列号后4位
		randomHex)            // 4位随机十六进制
}
