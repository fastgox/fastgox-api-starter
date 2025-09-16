package handle

import (
	"net/http"

	"github.com/fastgox/fastgox-api-starter/src/models/dto"
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
// @Success 200 {object} dto.Response{data=response.FileUploadResult} "上传成功"
// @Failure 400 {object} dto.Response "请求参数错误"
// @Failure 401 {object} dto.Response "未授权"
// @Failure 500 {object} dto.Response "服务器错误"
// @Router /file/upload/single [post]
func UploadSingleFile(c *gin.Context) {
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		logger.Warn("获取上传文件失败: %v", err)
		c.JSON(http.StatusBadRequest, dto.Response{
			Code:    400,
			Message: "请选择要上传的文件",
		})
		return
	}

	// 使用文件服务上传文件
	result, err := services.FileSvc.UploadSingleFile(file)
	if err != nil {
		logger.Error("文件上传失败: %v", err)
		c.JSON(http.StatusBadRequest, dto.Response{
			Code:    400,
			Message: "文件上传失败: " + err.Error(),
		})
		return
	}

	logger.Info("文件上传成功 fileName=%s", file.Filename)

	c.JSON(http.StatusOK, dto.Response{
		Code:    200,
		Message: "文件上传成功",
		Data:    result,
	})
}

func init() {
	// 注册文件路由（需要认证）
	router.AuthRouter.POST("/file/upload/single", UploadSingleFile)
}
