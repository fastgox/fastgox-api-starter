package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "run", "dev", "start":
		runGame() // 直接启动游戏服务器
	case "build":
		buildGame() // 直接构建游戏服务器
	case "clean":
		clean()
	case "fmt":
		formatCode()
	case "tidy":
		tidyDeps()
	case "swagger":
		generateSwagger()
	case "help":
		showHelp()
	default:
		fmt.Printf("未知命令: %s\n", command)
		showHelp()
	}
}

func runGame() {
	fmt.Println("🎮 启动TCP游戏服务器...")
	cmd := exec.Command("go", "run", "cmd/server/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Printf("游戏服务器启动失败: %v\n", err)
		os.Exit(1)
	}
}

func buildGame() {
	fmt.Println("🔨 构建TCP游戏服务器...")
	cmd := exec.Command("go", "build", "-o", "bin/fastgox-tcp-game-server", "cmd/server/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("游戏服务器构建失败: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("🎮 TCP游戏服务器构建完成: bin/fastgox-tcp-game-server")
}

func clean() {
	fmt.Println("🧹 清理构建文件...")
	targets := []string{"bin/", "coverage.out", "coverage.html"}
	for _, target := range targets {
		if err := os.RemoveAll(target); err == nil {
			fmt.Printf("删除 %s\n", target)
		}
	}
	exec.Command("go", "clean").Run()
	fmt.Println("清理完成!")
}

func formatCode() {
	fmt.Println("格式化代码...")
	cmd := exec.Command("go", "fmt", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("格式化失败: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("代码格式化完成")
}

func tidyDeps() {
	fmt.Println("📦 整理依赖...")
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("整理依赖失败: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("依赖整理完成!")
}

func generateSwagger() {
	fmt.Println("生成 Swagger 文档...")
	cmd := exec.Command("swag", "init", "-g", "cmd/server/main.go", "-o", "docs")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Swagger 文档生成失败: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Swagger 文档已生成到 docs 目录")
}

func showHelp() {
	fmt.Println("🎮 FastGox TCP游戏服务器 开发工具")
	fmt.Println()
	fmt.Println("用法: go run main.go <命令>")
	fmt.Println()
	fmt.Println("可用命令:")
	fmt.Println("  run/dev/start - 启动TCP游戏服务器")
	fmt.Println("  build         - 构建TCP游戏服务器")
	fmt.Println("  clean         - 清理构建文件")
	fmt.Println("  fmt           - 格式化代码")
	fmt.Println("  tidy          - 整理Go依赖")
	fmt.Println("  swagger       - 生成 Swagger 文档")
	fmt.Println("  help          - 显示帮助")
	fmt.Println()
	fmt.Println("示例:")
	fmt.Println("  go run main.go run     # 启动TCP游戏服务器")
	fmt.Println("  go run main.go build   # 构建TCP游戏服务器")
	fmt.Println("  go run main.go clean   # 清理构建文件")
	fmt.Println()
	fmt.Println("🚀 纯TCP游戏服务器，基于nano框架")
}
