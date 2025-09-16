package services

import (
	"fmt"
	"io"
	"mime/multipart"

	"github.com/fastgox/fastgox-api-starter/src/models/dto/response"
	"github.com/fastgox/fastgox-api-starter/src/pkg/file"
)

// FileService 文件服务
type FileService struct {
}

var FileSvc = &FileService{}

// UploadSingleFile 上传单个文件
func (s *FileService) UploadSingleFile(fileHeader *multipart.FileHeader) (*response.FileUploadResult, error) {
	// 获取文件提供商
	provider, err := file.FileManager.Get("")
	if err != nil {
		return &response.FileUploadResult{
			Success: false,
			Message: "没有可用的文件服务提供商: " + err.Error(),
			File:    nil,
		}, err
	}

	// 读取文件数据
	fileData, err := readFileData(fileHeader)
	if err != nil {
		return &response.FileUploadResult{
			Success: false,
			Message: "读取文件数据失败: " + err.Error(),
			File:    nil,
		}, err
	}

	// 构造输入参数
	input := &file.FileUploadInput{
		FileName: fileHeader.Filename,
		FileData: fileData,
		FileSize: fileHeader.Size,
	}

	// 调用文件上传
	result, err := provider.Call(input)
	if err != nil {
		return &response.FileUploadResult{
			Success: false,
			Message: "文件上传失败: " + err.Error(),
			File:    nil,
		}, err
	}

	return &response.FileUploadResult{
		Success: result.Success,
		Message: result.Message,
		File:    result,
	}, nil
}

// readFileData 读取文件数据
func readFileData(fileHeader *multipart.FileHeader) ([]byte, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("打开文件失败: %v", err)
	}
	defer file.Close()

	// 读取文件内容
	fileData, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("读取文件内容失败: %v", err)
	}

	return fileData, nil
}
