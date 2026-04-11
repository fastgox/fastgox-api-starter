package handle

import (
	"strings"

	"github.com/fastgox/fastgox-api-starter/src/core/i18n"
	"github.com/fastgox/fastgox-api-starter/src/core/tcp"
	"github.com/fastgox/fastgox-api-starter/src/models/dto/request"
	"github.com/fastgox/fastgox-api-starter/src/models/dto/response"
	"github.com/fastgox/fastgox-api-starter/src/router"
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
// @Success 200 {object} response.Response "发送成功"
// @Failure 400 {object} response.Response "参数错误"
// @Failure 500 {object} response.Response "发送失败"
// @Router /auth/send-login-sms [post]
func SendLoginSms(c *gin.Context) {
	var req request.SendLoginSmsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, i18n.BadRequest)
		return
	}

	req.Phone = strings.TrimSpace(req.Phone)
	if len(req.Phone) != 11 {
		response.BadRequest(c, i18n.PhoneFormatError)
		return
	}

	result, err := services.UserSvc.SendLoginSms(req.Phone)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	if !result.Success {
		response.BadRequest(c, result.Message)
		return
	}

	response.OKMsg(c, i18n.SmsSendSuccess, result)
}

// LoginWithSms 使用短信验证码登录
// @Summary 短信验证码登录
// @Description 使用手机号和验证码登录，不存在账户时自动注册
// @Tags 用户认证
// @Accept json
// @Produce json
// @Param request body request.LoginWithSmsRequest true "登录请求参数"
// @Success 200 {object} response.Response "登录成功"
// @Failure 400 {object} response.Response "参数错误"
// @Failure 500 {object} response.Response "登录失败"
// @Router /auth/login-with-sms [post]
func LoginWithSms(c *gin.Context) {
	var req request.LoginWithSmsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, i18n.BadRequest)
		return
	}

	req.Phone = strings.TrimSpace(req.Phone)
	req.Code = strings.TrimSpace(req.Code)

	if len(req.Phone) != 11 {
		response.BadRequest(c, i18n.PhoneFormatError)
		return
	}

	if len(req.Code) != 4 {
		response.BadRequest(c, i18n.CodeFormatError)
		return
	}

	result, err := services.UserSvc.LoginWithSms(req.Phone, req.Code)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	if !result.Success {
		response.BadRequest(c, result.Message)
		return
	}

	if result.Token != "" {
		c.Header("Authorization", "Bearer "+result.Token)
	}

	response.OKMsg(c, i18n.LoginSuccess, result)
}

func init() {
	// 注册用户认证路由（公开接口）
	router.PublicRouter.POST("/auth/send-login-sms", SendLoginSms)
	router.PublicRouter.POST("/auth/login-with-sms", LoginWithSms)

	// TCP 路由
	router.TCPRouter.Handle("echo", TCPEcho)
}

// ===== TCP Handlers =====

// TCPEcho echo测试
func TCPEcho(ctx *tcp.Context) error {
	response.OK(ctx, map[string]any{
		"echo": string(ctx.Msg.Data),
	})
	return nil
}
