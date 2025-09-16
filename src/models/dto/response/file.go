package response

import "github.com/fastgox/fastgox-api-starter/src/pkg/file"

// FileUploadResult 单文件上传结果
type FileUploadResult struct {
	Success bool             `json:"success"` // 上传是否成功
	Message string           `json:"message"` // 响应消息
	File    *file.FileOutput `json:"file"`    // 文件信息
}
