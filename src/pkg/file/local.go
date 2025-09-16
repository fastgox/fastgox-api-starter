package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fastgox/fastgox-api-starter/src/core/config"
	"github.com/fastgox/fastgox-api-starter/src/pkg"
	"github.com/google/uuid"
)

// LocalFileProvider 本地文件存储提供商
type LocalFileProvider struct{}

// GetName 获取提供商名称
func (p *LocalFileProvider) GetName() string {
	return ProviderNameLocal
}

// Call 调用文件服务
func (p *LocalFileProvider) Call(input FileInput) (*FileOutput, error) {
	startTime := time.Now()

	switch req := input.(type) {
	case *FileUploadInput:
		return p.uploadFile(req, startTime)
	default:
		return &FileOutput{
			BaseOutput: pkg.BaseOutput{
				Success:   false,
				Code:      "INVALID_INPUT",
				Message:   "不支持的输入类型",
				Provider:  p.GetName(),
				CostTime:  time.Since(startTime).Milliseconds(),
				Timestamp: time.Now(),
			},
		}, fmt.Errorf("不支持的输入类型")
	}
}

// uploadFile 上传文件
func (p *LocalFileProvider) uploadFile(input *FileUploadInput, startTime time.Time) (*FileOutput, error) {
	if config.GlobalConfig == nil {
		return &FileOutput{
			BaseOutput: pkg.BaseOutput{
				Success:   false,
				Code:      "CONFIG_ERROR",
				Message:   "配置未初始化",
				RequestID: input.RequestID,
				Provider:  p.GetName(),
				CostTime:  time.Since(startTime).Milliseconds(),
				Timestamp: time.Now(),
			},
		}, fmt.Errorf("配置未初始化")
	}

	fileConfig := config.GlobalConfig.File

	// 验证文件大小
	if input.FileSize > fileConfig.MaxSize {
		return &FileOutput{
			BaseOutput: pkg.BaseOutput{
				Success:   false,
				Code:      "FILE_TOO_LARGE",
				Message:   fmt.Sprintf("文件大小超过限制，最大允许 %d 字节", fileConfig.MaxSize),
				RequestID: input.RequestID,
				Provider:  p.GetName(),
				CostTime:  time.Since(startTime).Milliseconds(),
				Timestamp: time.Now(),
			},
		}, fmt.Errorf("文件大小超过限制")
	}

	// 验证文件扩展名
	ext := strings.ToLower(filepath.Ext(input.FileName))
	if !isAllowedExtension(ext, fileConfig.AllowedExtensions) {
		return &FileOutput{
			BaseOutput: pkg.BaseOutput{
				Success:   false,
				Code:      "INVALID_FILE_TYPE",
				Message:   fmt.Sprintf("不支持的文件类型: %s", ext),
				RequestID: input.RequestID,
				Provider:  p.GetName(),
				CostTime:  time.Since(startTime).Milliseconds(),
				Timestamp: time.Now(),
			},
		}, fmt.Errorf("不支持的文件类型")
	}

	// 创建按日期分组的目录
	now := time.Now()
	dateDir := now.Format("2006/01/02")
	uploadDir := filepath.Join(fileConfig.UploadPath, dateDir)

	// 确保目录存在
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return &FileOutput{
			BaseOutput: pkg.BaseOutput{
				Success:   false,
				Code:      "CREATE_DIR_FAILED",
				Message:   fmt.Sprintf("创建目录失败: %v", err),
				RequestID: input.RequestID,
				Provider:  p.GetName(),
				CostTime:  time.Since(startTime).Milliseconds(),
				Timestamp: time.Now(),
			},
		}, err
	}

	// 生成唯一文件名
	savedName := generateUniqueFileName(input.FileName)
	filePath := filepath.Join(uploadDir, savedName)
	// 相对路径包含日期目录
	relativePath := filepath.Join(dateDir, savedName)

	// 创建并写入文件
	file, err := os.Create(filePath)
	if err != nil {
		return &FileOutput{
			BaseOutput: pkg.BaseOutput{
				Success:   false,
				Code:      "CREATE_FILE_FAILED",
				Message:   fmt.Sprintf("创建文件失败: %v", err),
				RequestID: input.RequestID,
				Provider:  p.GetName(),
				CostTime:  time.Since(startTime).Milliseconds(),
				Timestamp: time.Now(),
			},
		}, err
	}
	defer file.Close()

	// 写入文件数据
	if _, err := file.Write(input.FileData); err != nil {
		return &FileOutput{
			BaseOutput: pkg.BaseOutput{
				Success:   false,
				Code:      "WRITE_FILE_FAILED",
				Message:   fmt.Sprintf("写入文件失败: %v", err),
				RequestID: input.RequestID,
				Provider:  p.GetName(),
				CostTime:  time.Since(startTime).Milliseconds(),
				Timestamp: time.Now(),
			},
		}, err
	}

	// 构建文件访问URL，使用相对路径
	fileURL := fmt.Sprintf("%s/%s", fileConfig.URLPrefix, strings.ReplaceAll(relativePath, "\\", "/"))

	return &FileOutput{
		BaseOutput: pkg.BaseOutput{
			Success:   true,
			Code:      "SUCCESS",
			Message:   "文件上传成功",
			RequestID: input.RequestID,
			Provider:  p.GetName(),
			CostTime:  time.Since(startTime).Milliseconds(),
			Timestamp: time.Now(),
		},
		FileName:   input.FileName,
		SavedName:  savedName,
		FilePath:   relativePath, // 返回相对路径，便于前端使用
		FileURL:    fileURL,
		FileSize:   input.FileSize,
		FileExt:    ext,
		UploadTime: time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}

// isAllowedExtension 检查文件扩展名是否被允许
func isAllowedExtension(ext string, allowedExts []string) bool {
	for _, allowedExt := range allowedExts {
		if strings.EqualFold(ext, allowedExt) {
			return true
		}
	}
	return false
}

// generateUniqueFileName 生成唯一的文件名
func generateUniqueFileName(originalName string) string {
	ext := filepath.Ext(originalName)
	nameWithoutExt := strings.TrimSuffix(originalName, ext)

	// 生成时间戳和UUID
	timestamp := time.Now().Format("20060102150405")
	uniqueID := strings.ReplaceAll(uuid.New().String(), "-", "")[:8]

	return fmt.Sprintf("%s_%s_%s%s", nameWithoutExt, timestamp, uniqueID, ext)
}

// InitLocalProvider 初始化本地文件提供商
func InitLocalProvider() {
	FileManager.Register(ProviderNameLocal, &LocalFileProvider{})
}
