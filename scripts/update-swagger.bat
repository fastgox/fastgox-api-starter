@echo off
chcp 65001 >nul

rem fastgox-api-starter API Swagger 文档更新脚本
echo 🔄 更新 fastgox-api-starter API Swagger 文档...

rem 切换到项目根目录
cd /d "%~dp0\.."

rem 1. 尝试自动生成
echo 📝 尝试自动生成 swagger 文档...
%USERPROFILE%\go\bin\swag.exe init -g cmd/server/main.go --parseDependency --parseInternal --output docs

rem 2. 检查是否包含知识库API
findstr /c:"knowledge" docs\swagger.yaml >nul 2>&1
if %errorlevel% equ 0 (
    echo ✅ 自动生成成功！知识库API已包含。
) else (
    echo ⚠️  自动生成缺少知识库API，使用手动文档...
    
    rem 使用手动维护的完整文档
    if exist "docs\swagger_manual.yaml" (
        copy "docs\swagger_manual.yaml" "docs\swagger.yaml" >nul
        echo ✅ 已使用手动维护的完整文档
    ) else (
        echo ❌ 手动文档不存在！
        pause
        exit /b 1
    )
)

rem 3. 更新 docs.go 文件
echo 🔧 更新 docs.go 文件...
%USERPROFILE%\go\bin\swag.exe fmt

echo 🎉 Swagger 文档更新完成！
echo 📖 访问地址: http://localhost:8080/swagger/index.html
pause 