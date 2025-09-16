package handle

import (
	"encoding/json"
	"net/http"

	"github.com/fastgox/fastgox-api-starter/src/core/config"
	"github.com/fastgox/fastgox-api-starter/src/models/dto"
	"github.com/fastgox/fastgox-api-starter/src/models/dto/request"
	"github.com/fastgox/fastgox-api-starter/src/pkg/ocr"
	"github.com/fastgox/fastgox-api-starter/src/router"
	"github.com/gin-gonic/gin"
)

// RecognizeIdCard 身份证识别
// @Summary 身份证识别
// @Description 使用阿里云OCR识别身份证信息
// @Tags OCR
// @Accept json
// @Produce json
// @Param request body request.RecognizeIdCardRequest true "识别请求参数"
// @Success 200 {object} dto.Response "识别成功"
// @Failure 400 {object} dto.Response "参数错误"
// @Failure 500 {object} dto.Response "识别失败"
// @Router /ocr/idcard [post]
func RecognizeIdCard(c *gin.Context) {
	var req request.RecognizeIdCardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code:    400,
			Message: "参数错误: " + err.Error(),
		})
		return
	}

	// 验证参数
	if req.Url == "" && req.Body == "" {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code:    400,
			Message: "URL和Body参数不能同时为空，请提供其中一种方式",
		})
		return
	}

	if req.Url != "" && req.Body != "" {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code:    400,
			Message: "URL和Body参数不能同时提供，请选择其中一种方式",
		})
		return
	}

	// 获取OCR提供商
	provider, err := ocr.OcrManager.Get(config.GlobalConfig.OCR.Engine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Code:    500,
			Message: "没有可用的OCR服务提供商: " + err.Error(),
		})
		return
	}

	// 构造OCR输入参数
	input := &ocr.RecognizeIdCardInput{
		Url:      req.Url,
		Body:     req.Body,
		CardSide: req.CardSide,
	}

	// 调用OCR识别
	result, err := provider.Call(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Code:    500,
			Message: "识别失败: " + err.Error(),
		})
		return
	}

	// 直接返回OCR结果中的Data字段，避免嵌套的data.data结构
	var responseData interface{}
	if result.Data != "" {
		// 尝试解析JSON字符串
		if err := json.Unmarshal([]byte(result.Data), &responseData); err != nil {
			// 如果解析失败，直接返回字符串
			responseData = result.Data
		}
	} else {
		responseData = result
	}

	c.JSON(http.StatusOK, dto.Response{
		Code:    200,
		Message: result.Message,
		Data:    responseData,
	})
}

func init() {
	// 注册OCR路由
	router.PublicRouter.POST("/ocr/idcard", RecognizeIdCard)
}
