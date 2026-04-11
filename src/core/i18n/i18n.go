package i18n

import (
	"embed"

	"github.com/fastgox/fastgox-api-starter/src/core/tcp"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

//go:embed locales/*.yaml
var localeFS embed.FS

var (
	bundle      *i18n.Bundle
	defaultLang = language.Chinese
)

// Init 初始化 i18n，加载语言文件
func Init() {
	bundle = i18n.NewBundle(defaultLang)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	// 加载嵌入的语言文件
	entries, err := localeFS.ReadDir("locales")
	if err != nil {
		panic("i18n: 无法读取 locales 目录: " + err.Error())
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			bundle.LoadMessageFileFS(localeFS, "locales/"+entry.Name())
		}
	}
}

// T 根据消息ID翻译，使用默认语言
func T(msgID string) string {
	localizer := i18n.NewLocalizer(bundle, defaultLang.String())
	msg, err := localizer.Localize(&i18n.LocalizeConfig{MessageID: msgID})
	if err != nil {
		return msgID
	}
	return msg
}

// TWithLang 根据消息ID和语言翻译
func TWithLang(lang string, msgID string) string {
	localizer := i18n.NewLocalizer(bundle, lang)
	msg, err := localizer.Localize(&i18n.LocalizeConfig{MessageID: msgID})
	if err != nil {
		return msgID
	}
	return msg
}

// TWithData 带模板数据的翻译
func TWithData(lang string, msgID string, data map[string]interface{}) string {
	localizer := i18n.NewLocalizer(bundle, lang)
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    msgID,
		TemplateData: data,
	})
	if err != nil {
		return msgID
	}
	return msg
}

// LangFromGin 从 gin.Context 的 Accept-Language header 获取语言
func LangFromGin(c *gin.Context) string {
	lang := c.GetHeader("Accept-Language")
	if lang == "" {
		return defaultLang.String()
	}
	// 解析 Accept-Language，取最佳匹配
	tags, _, err := language.ParseAcceptLanguage(lang)
	if err != nil || len(tags) == 0 {
		return defaultLang.String()
	}
	matcher := language.NewMatcher(bundle.LanguageTags())
	_, idx, _ := matcher.Match(tags...)
	matched := bundle.LanguageTags()
	if idx < len(matched) {
		return matched[idx].String()
	}
	return defaultLang.String()
}

// LangFromString 解析语言字符串，返回标准化语言标签
func LangFromString(lang string) string {
	if lang == "" {
		return defaultLang.String()
	}
	tag, err := language.Parse(lang)
	if err != nil {
		return defaultLang.String()
	}
	return tag.String()
}

// LangFromCtx 从 context（gin.Context / tcp.Context）提取语言
func LangFromCtx(c interface{}) string {
	switch ctx := c.(type) {
	case *gin.Context:
		if lang, exists := ctx.Get("lang"); exists {
			return lang.(string)
		}
	case *tcp.Context:
		return ctx.GetString("lang")
	}
	return defaultLang.String()
}

// TFromCtx 从 context 提取语言并翻译（未命中则原文返回）
func TFromCtx(c interface{}, msgID string) string {
	return TWithLang(LangFromCtx(c), msgID)
}
