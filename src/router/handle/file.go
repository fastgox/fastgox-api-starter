package handle

import (
	"github.com/fastgox/fastgox-api-starter/src/core/i18n"
	"github.com/fastgox/fastgox-api-starter/src/models/dto/response"
	"github.com/fastgox/fastgox-api-starter/src/router"
	"github.com/fastgox/fastgox-api-starter/src/services"
	"github.com/fastgox/utils/logger"
	"github.com/gin-gonic/gin"
)

// UploadSingleFile 单文件上传
// @Summary 单文件上传
// @Description 上传单个文件
// @Tags 文件上传
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param file formData file true "上传的文件"
// @Success 200 {object} response.Response{data=response.FileUploadResult} "上传成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 401 {object} response.Response "未授权"
// @Failure 500 {object} response.Response "服务器错误"
// @Router /file/upload/single [post]
func UploadSingleFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		logger.Warn("获取上传文件失败: %v", err)
		response.BadRequest(c, i18n.FileRequired)
		return
	}

	result, err := services.FileSvc.UploadSingleFile(file)
	if err != nil {
		logger.Error("文件上传失败: %v", err)
		response.InternalError(c, err.Error())
		return
	}

	logger.Info("文件上传成功 fileName=%s", file.Filename)
	response.OKMsg(c, i18n.FileUploadSuccess, result)
}

func init() {
	// 注册文件路由（需要认证）
	router.AuthRouter.POST("/file/upload/single", UploadSingleFile)
}
