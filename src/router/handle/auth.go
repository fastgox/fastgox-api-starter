package handle

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fastgox/fastgox-api-starter/src/core/session"
	"github.com/fastgox/fastgox-api-starter/src/models/dto"
	"github.com/fastgox/fastgox-api-starter/src/models/dto/request"
	"github.com/fastgox/fastgox-api-starter/src/models/dto/response"
	"github.com/fastgox/fastgox-api-starter/src/router"
	"github.com/fastgox/fastgox-api-starter/src/services"
	"github.com/gin-gonic/gin"
)

// VerifyThreeElements 三要素认证
// @Summary 三要素认证
// @Description 通过姓名和身份证号，结合token中的手机号进行三要素认证
// @Tags 身份认证
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body request.ThreeElementsVerifyRequest true "认证请求参数，包含姓名和身份证号"
// @Success 200 {object} dto.Response "认证成功"
// @Failure 400 {object} dto.Response "参数错误"
// @Failure 401 {object} dto.Response "未授权"
// @Failure 500 {object} dto.Response "认证失败"
// @Router /auth/identity-verification [post]
func VerifyThreeElements(c *gin.Context) {
	var req request.ThreeElementsVerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code:    400,
			Message: "参数错误: " + err.Error(),
		})
		return
	}

	// 参数验证
	if err := validateThreeElementsRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	// 获取当前用户信息
	userSession, err := session.Manager.GetUserSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Code:    401,
			Message: "获取用户信息失败: " + err.Error(),
		})
		return
	}

	// 调用三要素认证服务，直接使用传入的姓名和身份证号
	result, err := services.AuthSvc.VerifyThreeElements(userSession.UserID, req.Name, req.IdCard, userSession.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Code:    500,
			Message: "认证失败: " + err.Error(),
		})
		return
	}

	// 构造响应
	response := response.AuthVerifyResponse{
		Success:      result.BaseOutput.Success,
		VerifyStatus: result.VerifyStatus,
		FlowNo:       result.FlowNo,
		ServiceId:    result.ServiceId,
		Message:      result.BaseOutput.Message,
		Provider:     result.BaseOutput.Provider,
		CostTime:     result.BaseOutput.CostTime,
	}

	c.JSON(http.StatusOK, dto.Response{
		Code:    200,
		Message: "认证完成",
		Data:    response,
	})
}

// validateThreeElementsRequest 验证三要素请求参数
func validateThreeElementsRequest(req *request.ThreeElementsVerifyRequest) error {
	// 姓名验证
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return fmt.Errorf("姓名不能为空")
	}

	// 身份证号验证
	req.IdCard = strings.TrimSpace(req.IdCard)
	if req.IdCard == "" {
		return fmt.Errorf("身份证号不能为空")
	}

	return nil
}

func init() {
	// 注册身份认证路由（需要用户认证）
	router.AuthRouter.POST("/auth/identity-verification", VerifyThreeElements)
}
