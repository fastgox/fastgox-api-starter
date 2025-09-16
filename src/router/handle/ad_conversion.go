package handle

import (
	"net/http"

	"github.com/fastgox/fastgox-api-starter/src/models/dto"
	"github.com/fastgox/fastgox-api-starter/src/models/dto/request"
	"github.com/fastgox/fastgox-api-starter/src/router"
	"github.com/fastgox/fastgox-api-starter/src/services"
	"github.com/gin-gonic/gin"
)

// CreateAdConversion 创建广告转化记录
// @Summary 创建广告转化记录
// @Description 记录用户的广告转化行为，包括首次打开app、注册、留资等
// @Tags 广告转化
// @Accept json
// @Produce json
// @Param request body request.CreateAdConversionRequest true "广告转化记录请求参数"
// @Success 200 {object} dto.Response{data=response.AdConversionResponse} "创建成功"
// @Failure 400 {object} dto.Response "参数错误"
// @Failure 500 {object} dto.Response "创建失败"
// @Router /ad-conversion/create [post]
func CreateAdConversion(c *gin.Context) {
	var req request.CreateAdConversionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code:    400,
			Message: "参数错误: " + err.Error(),
		})
		return
	}
	// 调用服务创建转化记录
	result, err := services.AdConversionSvc.CreateAdConversion(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Code:    500,
			Message: "创建广告转化记录失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Code:    200,
		Message: "创建成功",
		Data:    result,
	})
}

func init() {

	router.PublicRouter.POST("/ad-conversion/create", CreateAdConversion)
}
