package handle

import (
	"net/http"
	"strings"

	"github.com/fastgox/fastgox-api-starter/src/models/dto"
	"github.com/fastgox/fastgox-api-starter/src/models/dto/request"
	"github.com/fastgox/fastgox-api-starter/src/services"
	"github.com/gin-gonic/gin"
)

// SendLoginSms 发送登录短信验证码
// @Summary 发送登录短信验证码
// @Description 向指定手机号发送登录验证码
// @Tags 用户认证
// @Accept json
// @Produce json
// @Param request body request.SendLoginSmsRequest true "发送短信请求参数"
// @Success 200 {object} dto.Response "发送成功"
// @Failure 400 {object} dto.Response "参数错误"
// @Failure 500 {object} dto.Response "发送失败"
// @Router /auth/send-login-sms [post]
func SendLoginSms(c *gin.Context) {
	var req request.SendLoginSmsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code:    400,
			Message: "参数错误: " + err.Error(),
		})
		return
	}

	// 参数验证
	req.Phone = strings.TrimSpace(req.Phone)
	if len(req.Phone) != 11 {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code:    400,
			Message: "手机号格式不正确",
		})
		return
	}

	// 调用服务发送短信
	result, err := services.UserSvc.SendLoginSms(req.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Code:    500,
			Message: "发送短信失败: " + err.Error(),
		})
		return
	}

	if !result.Success {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code:    400,
			Message: result.Message,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Code:    200,
		Message: result.Message,
		Data:    result,
	})
}

// LoginWithSms 使用短信验证码登录
// @Summary 短信验证码登录
// @Description 使用手机号和验证码登录，不存在账户时自动注册
// @Tags 用户认证
// @Accept json
// @Produce json
// @Param request body request.LoginWithSmsRequest true "登录请求参数"
// @Success 200 {object} dto.Response "登录成功"
// @Failure 400 {object} dto.Response "参数错误"
// @Failure 500 {object} dto.Response "登录失败"
// @Router /auth/login-with-sms [post]
func LoginWithSms(c *gin.Context) {
	var req request.LoginWithSmsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code:    400,
			Message: "参数错误: " + err.Error(),
		})
		return
	}

	// 参数验证
	req.Phone = strings.TrimSpace(req.Phone)
	req.Code = strings.TrimSpace(req.Code)

	if len(req.Phone) != 11 {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code:    400,
			Message: "手机号格式不正确",
		})
		return
	}

	if len(req.Code) != 4 {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code:    400,
			Message: "验证码格式不正确",
		})
		return
	}

	// 调用服务进行登录
	result, err := services.UserSvc.LoginWithSms(req.Phone, req.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Code:    500,
			Message: "登录失败: " + err.Error(),
		})
		return
	}

	if !result.Success {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code:    400,
			Message: result.Message,
		})
		return
	}

	// 设置响应头中的令牌
	if result.Token != "" {
		c.Header("Authorization", "Bearer "+result.Token)
	}

	c.JSON(http.StatusOK, dto.Response{
		Code:    200,
		Message: result.Message,
		Data:    result,
	})
}
