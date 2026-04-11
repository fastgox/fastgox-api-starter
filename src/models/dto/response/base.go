package response

import (
	"fmt"
	"net/http"

	"github.com/fastgox/fastgox-api-starter/src/core/i18n"
	"github.com/fastgox/fastgox-api-starter/src/core/tcp"
	"github.com/gin-gonic/gin"
)

// Response 基础响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// PageResponse 分页响应结构
type PageResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Total   int64       `json:"total"`
	Page    int         `json:"page"`
	Size    int         `json:"size"`
}

// BizError 业务错误，携带错误码和HTTP状态码
type BizError struct {
	HTTPCode int    // HTTP状态码
	Code     int    // 业务错误码
	MsgID    string // i18n 消息ID（优先使用）
	Message  string // 降级文案（MsgID 未命中时使用）
}

func (e *BizError) Error() string {
	if e.MsgID != "" {
		return i18n.T(e.MsgID)
	}
	return e.Message
}

// 预定义业务错误码（使用 i18n 消息ID）
var (
	ErrBadRequest     = &BizError{HTTPCode: http.StatusBadRequest, Code: 400, MsgID: i18n.BadRequest}
	ErrUnauthorized   = &BizError{HTTPCode: http.StatusUnauthorized, Code: 401, MsgID: i18n.Unauthorized}
	ErrForbidden      = &BizError{HTTPCode: http.StatusForbidden, Code: 403, MsgID: i18n.Forbidden}
	ErrNotFound       = &BizError{HTTPCode: http.StatusNotFound, Code: 404, MsgID: i18n.NotFound}
	ErrInternal       = &BizError{HTTPCode: http.StatusInternalServerError, Code: 500, MsgID: i18n.ServerError}
	ErrTooManyRequest = &BizError{HTTPCode: http.StatusTooManyRequests, Code: 429, MsgID: i18n.TooManyRequests}
)

// NewBizError 创建自定义业务错误（纯文本降级）
func NewBizError(httpCode, code int, message string) *BizError {
	return &BizError{HTTPCode: httpCode, Code: code, Message: message}
}

// BadRequest 快捷400错误响应（支持 i18n 消息ID 或纯文本）
func BadRequest(c interface{}, msgID string) {
	Fail(c, &BizError{HTTPCode: http.StatusBadRequest, Code: 400, MsgID: msgID})
}

// InternalError 快捷500错误响应（支持 i18n 消息ID 或纯文本）
func InternalError(c interface{}, msgID string) {
	Fail(c, &BizError{HTTPCode: http.StatusInternalServerError, Code: 500, MsgID: msgID})
}

// Unauthorized 快捷401错误响应（支持 i18n 消息ID 或纯文本）
func Unauthorized(c interface{}, msgID string) {
	Fail(c, &BizError{HTTPCode: http.StatusUnauthorized, Code: 401, MsgID: msgID})
}

// ===== 从 context 提取语言 =====

func getLang(c interface{}) string {
	switch ctx := c.(type) {
	case *gin.Context:
		if lang, exists := ctx.Get("lang"); exists {
			return lang.(string)
		}
	case *tcp.Context:
		return ctx.GetString("lang")
	}
	return "zh"
}

// resolveMessage 解析 BizError 消息（i18n 优先，降级到 Message 字段）
func resolveMessage(c interface{}, err *BizError) string {
	if err.MsgID != "" {
		return i18n.TWithLang(getLang(c), err.MsgID)
	}
	return err.Message
}

// ===== 统一响应方法（自动适配 HTTP / TCP + i18n） =====

// OK 成功响应
func OK(c interface{}, data interface{}) {
	msg := i18n.TWithLang(getLang(c), i18n.Success)
	switch ctx := c.(type) {
	case *gin.Context:
		ctx.JSON(http.StatusOK, Response{Code: 200, Message: msg, Data: data})
	case *tcp.Context:
		ctx.ReplyMsg(200, msg, data)
	}
}

// OKMsg 成功响应（消息支持 i18n 消息ID，未命中则原文返回）
func OKMsg(c interface{}, msgID string, data interface{}) {
	msg := i18n.TWithLang(getLang(c), msgID)
	switch ctx := c.(type) {
	case *gin.Context:
		ctx.JSON(http.StatusOK, Response{Code: 200, Message: msg, Data: data})
	case *tcp.Context:
		ctx.ReplyMsg(200, msg, data)
	}
}

// Fail 失败响应，接受 error 或 *BizError
func Fail(c interface{}, err error) {
	bizErr, ok := err.(*BizError)
	if !ok {
		bizErr = &BizError{HTTPCode: http.StatusInternalServerError, Code: 500, Message: fmt.Sprintf("%v", err)}
	}
	msg := resolveMessage(c, bizErr)
	switch ctx := c.(type) {
	case *gin.Context:
		ctx.JSON(bizErr.HTTPCode, Response{Code: bizErr.Code, Message: msg})
	case *tcp.Context:
		ctx.Error(bizErr.Code, msg)
	}
}

// FailMsg 快捷失败响应（消息支持 i18n 消息ID）
func FailMsg(c interface{}, code int, msgID string) {
	msg := i18n.TWithLang(getLang(c), msgID)
	switch ctx := c.(type) {
	case *gin.Context:
		ctx.JSON(code, Response{Code: code, Message: msg})
	case *tcp.Context:
		ctx.Error(code, msg)
	}
}
