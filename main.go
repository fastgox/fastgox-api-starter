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
	case "run", "dev":
		runDev()
	case "build":
		build()
	case "clean":
		clean()
	case "fmt":
		formatCode()
	case "tidy":
		tidyDeps()
	case "help":
		showHelp()
	default:
		fmt.Printf("未知命令: %s\n", command)
		showHelp()
	}
}

func runDev() {
	fmt.Println("启动开发服务器...")
	cmd := exec.Command("go", "run", "cmd/server/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Printf("启动失败: %v\n", err)
		os.Exit(1)
	}
}

func build() {
	fmt.Println("🔨 构建项目 (Linux x86)...")
	cmd := exec.Command("go", "build", "-o", "bin/fastgox-api-starter", "cmd/server/main.go")
	cmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=386")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("构建失败: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("构建完成: bin/fastgox-api-starter (Linux x86)")
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

func showHelp() {
	fmt.Println("fastgox-api-starter API 开发工具")
	fmt.Println()
	fmt.Println("用法: go run main.go <命令>")
	fmt.Println()
	fmt.Println("可用命令:")
	fmt.Println("  run/dev  - 启动开发服务器")
	fmt.Println("  build    - 构建项目到 bin/fastgox-api-starter")
	fmt.Println("  clean    - 清理构建文件")
	fmt.Println("  fmt      - 格式化代码")
	fmt.Println("  tidy     - 整理Go依赖")
	fmt.Println("  help     - 显示帮助")
	fmt.Println()
	fmt.Println("示例:")
	fmt.Println("  go run main.go run    # 启动开发服务器")
	fmt.Println("  go run main.go build  # 构建项目")
}
