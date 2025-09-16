package file

import (
	"github.com/fastgox/fastgox-api-starter/src/pkg"
)

// 提供商名称常量
const (
	ProviderNameLocal = "local" // 本地文件存储提供商
)

// 全局文件服务提供商管理器
var FileManager = pkg.NewManager[pkg.Provider[FileInput, *FileOutput]](ProviderNameLocal)

// FileInput 文件输入联合类型
type FileInput interface {
	isFileInput()
}

// FileUploadInput 文件上传输入
type FileUploadInput struct {
	pkg.BaseInput
	FileName string `json:"file_name"` // 文件名
	FileData []byte `json:"file_data"` // 文件数据
	FileSize int64  `json:"file_size"` // 文件大小
}

// 实现 FileInput 接口
func (f *FileUploadInput) isFileInput() {}

// FileOutput 文件输出
type FileOutput struct {
	pkg.BaseOutput
	FileName   string `json:"file_name"`   // 原始文件名
	SavedName  string `json:"saved_name"`  // 保存的文件名
	FilePath   string `json:"file_path"`   // 文件相对路径
	FileURL    string `json:"file_url"`    // 文件访问URL
	FileSize   int64  `json:"file_size"`   // 文件大小
	FileExt    string `json:"file_ext"`    // 文件扩展名
	UploadTime string `json:"upload_time"` // 上传时间
}
