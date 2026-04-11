package i18n

// 消息ID常量，与 locales/*.yaml 中的 id 一一对应

// 通用
const (
	Success     = "success"
	ServerError = "server_error"
)

// 参数校验
const (
	BadRequest       = "bad_request"
	ParamError       = "param_error"
	PhoneFormatError = "phone_format_error"
	CodeFormatError  = "code_format_error"
)

// 认证
const (
	Unauthorized  = "unauthorized"
	Forbidden     = "forbidden"
	TokenMissing  = "token_missing"
	TokenFmtError = "token_format_error"
	TokenEmpty    = "token_empty"
	TokenInvalid  = "token_invalid"
)

// 资源
const (
	NotFound        = "not_found"
	TooManyRequests = "too_many_requests"
)

// 用户 / 短信
const (
	SmsSendSuccess = "sms_send_success"
	SmsSendFailed  = "sms_send_failed"
	LoginSuccess   = "login_success"
	LoginFailed    = "login_failed"
)

// 文件
const (
	FileRequired      = "file_required"
	FileUploadSuccess = "file_upload_success"
	FileUploadFailed  = "file_upload_failed"
)
