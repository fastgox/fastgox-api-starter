package response

// OpenIdCardResult 开放接口身份证识别结果
type OpenIdCardResult struct {
	Name         string `json:"name"`         // 姓名 - 人像面
	Sex          string `json:"sex"`          // 性别 - 人像面
	Nation       string `json:"nation"`       // 民族 - 人像面
	Birth        string `json:"birth"`        // 出生日期 - 人像面
	Address      string `json:"address"`      // 地址 - 人像面
	PsnIdCardNum string `json:"psnIdCardNum"` // 身份证号码 - 人像面
	Authority    string `json:"authority"`    // 发证机关 - 国徽面
	ValidDate    string `json:"validDate"`    // 证件有效期 - 国徽面
}

// OpenIdCardResponse 开放接口身份证识别响应
type OpenIdCardResponse struct {
	Result  OpenIdCardResult `json:"result"`  // 识别结果
	Code    string           `json:"code"`    // 状态码
	Msg     string           `json:"msg"`     // 消息
	LogKey  string           `json:"logKey"`  // 日志键
	Ts      string           `json:"ts"`      // 时间戳
	Success bool             `json:"success"` // 是否成功
}
