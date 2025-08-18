#!/bin/bash

# fastgox-api-starter API Swagger 文档更新脚本
echo "🔄 更新 fastgox-api-starter API Swagger 文档..."

# 切换到项目根目录
cd "$(dirname "$0")/.."

# 1. 尝试自动生成
echo "📝 尝试自动生成 swagger 文档..."
~/go/bin/swag init -g cmd/server/main.go --parseDependency --parseInternal --output docs

# 2. 检查是否包含知识库API
if grep -q "knowledge" docs/swagger.yaml; then
    echo "✅ 自动生成成功！知识库API已包含。"
else
    echo "⚠️  自动生成缺少知识库API，使用手动文档..."
    
    # 使用手动维护的完整文档
    if [ -f "docs/swagger_manual.yaml" ]; then
        cp docs/swagger_manual.yaml docs/swagger.yaml
        echo "✅ 已使用手动维护的完整文档"
    else
        echo "❌ 手动文档不存在！"
        exit 1
    fi
fi

# 3. 更新 docs.go 文件
echo "🔧 更新 docs.go 文件..."
~/go/bin/swag fmt

echo "🎉 Swagger 文档更新完成！"
echo "📖 访问地址: http://localhost:8080/swagger/index.html" 