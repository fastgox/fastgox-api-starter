package main

import (
	"fmt"

	"github.com/fastgox/fastgox-api-starter/src/config"
	"github.com/fastgox/fastgox-api-starter/src/repository"
	"github.com/fastgox/utils/logger"
)

func main() {
	logger.InitWithPath("data/logs")
	logger.Info("测试数据库自动初始化...")

	// 初始化配置
	if err := config.InitConfig(); err != nil {
		logger.Error("配置初始化失败: ", err)
		return
	}

	// 直接获取仓储，这应该会触发自动初始化
	repo := repository.GetRepository[interface{}]()
	
	if repo != nil {
		logger.Info("✅ 数据库自动初始化成功！")
		fmt.Println("✅ 数据库自动初始化测试通过")
	} else {
		logger.Error("❌ 数据库自动初始化失败")
		fmt.Println("❌ 数据库自动初始化测试失败")
	}
}
