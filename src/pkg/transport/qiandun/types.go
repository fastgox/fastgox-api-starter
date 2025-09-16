package qiandun

// QiandunResponse 千盾通用响应结构
type QiandunResponse struct {
	Result   interface{} `json:"result"`
	Code     string      `json:"code"`
	Msg      string      `json:"msg"`
	LogKey   string      `json:"logKey"`
	Ts       string      `json:"ts"`
	Success  bool        `json:"success"`
	CostTime int64       `json:"costTime"` // 耗时（毫秒）
}

// Config 千盾配置
type Config struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	BaseURL   string `json:"base_url"`
	IsProd    bool   `json:"is_prod"`
}
