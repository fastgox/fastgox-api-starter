package config

// import "time"

// // AppConfiguration 应用配置结构（对应数据库中的JSON数据）
// //
// //	@Description	完整的应用配置对象
// type AppConfiguration struct {
// 	Version         int                   `json:"version" example:"1"`
// 	UI              UIConfig              `json:"ui"`
// 	OpenAI          AppOpenAIConfig       `json:"openai"`
// 	RAG             RAGConfig             `json:"rag"`
// 	GoogleDrive     GoogleDriveConfig     `json:"google_drive"`
// 	OneDrive        OneDriveConfig        `json:"onedrive"`
// 	WebhookURL      string                `json:"webhook_url" example:"https://example.com/webhook"`
// 	Auth            AppAuthConfig         `json:"auth"`
// 	WebUI           WebUIConfig           `json:"webui"`
// 	Channels        ChannelsConfig        `json:"channels"`
// 	Notes           NotesConfig           `json:"notes"`
// 	Audio           AudioConfig           `json:"audio"`
// 	LDAP            LDAPConfig            `json:"ldap"`
// 	Task            TaskConfig            `json:"task"`
// 	Ollama          OllamaConfig          `json:"ollama"`
// 	Direct          DirectConfig          `json:"direct"`
// 	User            UserConfig            `json:"user"`
// 	Evaluation      EvaluationConfig      `json:"evaluation"`
// 	Recaptcha       RecaptchaConfig       `json:"recaptcha"`
// 	SMTP            SMTPConfig            `json:"smtp"`
// 	ContentFilter   ContentFilterConfig   `json:"content_filter"`
// 	Langfuse        LangfuseConfig        `json:"langfuse"`
// 	DeepResearch    DeepResearchConfig    `json:"deep_research"`
// 	ToolServer      ToolServerConfig      `json:"tool_server"`
// 	CodeExecution   CodeExecutionConfig   `json:"code_execution"`
// 	CodeInterpreter CodeInterpreterConfig `json:"code_interpreter"`
// 	ImageGeneration ImageGenConfig        `json:"image_generation"`
// 	Anyscale        AnyscaleConfig        `json:"anyscale,omitempty"`    // Anyscale配置
// 	Cohere          CohereConfig          `json:"cohere,omitempty"`      // Cohere配置
// 	Anthropic       AnthropicConfig       `json:"anthropic,omitempty"`   // Anthropic配置
// 	LlamaCloud      LlamaCloudConfig      `json:"llama_cloud,omitempty"` // LlamaCloud配置
// 	Proxy           ProxyConfig           `json:"proxy,omitempty"`       // 代理配置
// }

// // UIConfig UI配置
// //
// //	@Description	UI相关配置
// type UIConfig struct {
// 	EnableSignup              bool               `json:"enable_signup" example:"true"`
// 	DefaultUserRole           string             `json:"default_user_role" example:"user"`
// 	EnableCommunitySharing    bool               `json:"enable_community_sharing" example:"false"`
// 	EnableMessageRating       bool               `json:"enable_message_rating" example:"true"`
// 	EnableUserWebhooks        bool               `json:"enable_user_webhooks" example:"false"`
// 	PendingUserOverlayTitle   string             `json:"pending_user_overlay_title" example:"Account Pending"`
// 	PendingUserOverlayContent string             `json:"pending_user_overlay_content" example:"Your account is pending approval"`
// 	Watermark                 string             `json:"watermark" example:""`
// 	PromptSuggestions         []PromptSuggestion `json:"prompt_suggestions"`
// 	Banners                   []Banner           `json:"banners"`
// 	DefaultModels             string             `json:"default_models,omitempty"`   // 默认模型列表
// 	ModelOrderList            []string           `json:"model_order_list,omitempty"` // 模型显示顺序
// }

// // AppOpenAIConfig OpenAI配置（避免与现有OpenAIConfig冲突）
// //
// //	@Description	OpenAI相关配置
// type AppOpenAIConfig struct {
// 	Enable       bool                          `json:"enable" example:"true"`
// 	APIBaseURLs  []string                      `json:"api_base_urls" example:"[\"https://api.openai.com/v1\"]"`
// 	APIKeys      []string                      `json:"api_keys" example:"[\"sk-...\"]"`
// 	APIConfigs   map[string]AppOpenAIAPIConfig `json:"api_configs"`
// 	ModelMapping map[string]ModelInfo          `json:"model_mapping,omitempty"` // 模型ID -> ModelInfo映射
// 	Models       []ModelInfo                   `json:"models,omitempty"`        // 所有可用模型列表（用于遍历）
// }

// // ModelInfo 模型信息
// type ModelInfo struct {
// 	ID         string `json:"id"`          // 完整的模型ID（带前缀）
// 	OriginalID string `json:"original_id"` // 原始模型ID（用于API调用）
// 	Object     string `json:"object"`
// 	Created    int64  `json:"created"`
// 	OwnedBy    string `json:"owned_by"`
// 	BaseURL    string `json:"base_url"`  // 模型所属的BaseURL
// 	APIKey     string `json:"api_key"`   // 对应的API Key
// 	APIIndex   int    `json:"api_index"` // 对应的API配置索引
// 	Provider   string `json:"provider"`  // 提供商名称（如openai、deepseek等）
// }

// // AppOpenAIAPIConfig OpenAI API具体配置
// type AppOpenAIAPIConfig struct {
// 	Enable         bool     `json:"enable"`
// 	Tags           []string `json:"tags"`
// 	PrefixID       string   `json:"prefix_id"`
// 	ModelIDs       []string `json:"model_ids"`
// 	ConnectionType string   `json:"connection_type"`
// }

// // RAGConfig RAG配置
// //
// //	@Description	RAG（检索增强生成）配置
// type RAGConfig struct {
// 	Template                            string                 `json:"template" example:"default"`
// 	TopK                                int                    `json:"top_k" example:"5"`
// 	BypassEmbeddingAndRetrieval         bool                   `json:"bypass_embedding_and_retrieval" example:"false"`
// 	FullContext                         bool                   `json:"full_context" example:"false"`
// 	EnableHybridSearch                  bool                   `json:"enable_hybrid_search" example:"true"`
// 	TopKReranker                        int                    `json:"top_k_reranker" example:"3"`
// 	RelevanceThreshold                  float64                `json:"relevance_threshold" example:"0.7"`
// 	ChatModel                           string                 `json:"chat_model,omitempty"`                // 对话模型
// 	TaskModel                           string                 `json:"task_model,omitempty"`                // 任务模型
// 	TitleGenerationModel                string                 `json:"title_generation_model,omitempty"`    // 标题生成模型
// 	AzureAISearchThreshold              float64                `json:"azure_ai_search_threshold,omitempty"` // Azure AI搜索阈值
// 	ContentExtractionEngine             string                 `json:"CONTENT_EXTRACTION_ENGINE"`
// 	PDFExtractImages                    bool                   `json:"pdf_extract_images"`
// 	ExternalDocumentLoaderURL           string                 `json:"external_document_loader_url"`
// 	ExternalDocumentLoaderAPIKey        string                 `json:"external_document_loader_api_key"`
// 	TikaServerURL                       string                 `json:"tika_server_url"`
// 	DoclingServerURL                    string                 `json:"docling_server_url"`
// 	DoclingOCREngine                    string                 `json:"docling_ocr_engine"`
// 	DoclingOCRLang                      string                 `json:"docling_ocr_lang"`
// 	DoclingDoPictureDescription         bool                   `json:"docling_do_picture_description"`
// 	DocumentIntelligenceEndpoint        string                 `json:"document_intelligence_endpoint"`
// 	DocumentIntelligenceKey             string                 `json:"document_intelligence_key"`
// 	MistralOCRAPIKey                    string                 `json:"mistral_ocr_api_key"`
// 	RerankingEngine                     string                 `json:"reranking_engine"`
// 	ExternalRerankerURL                 string                 `json:"external_reranker_url"`
// 	ExternalRerankerAPIKey              string                 `json:"external_reranker_api_key"`
// 	RerankingModel                      string                 `json:"reranking_model"`
// 	TextSplitter                        string                 `json:"text_splitter"`
// 	ChunkSize                           int                    `json:"chunk_size"`
// 	ChunkOverlap                        int                    `json:"chunk_overlap"`
// 	File                                FileConfig             `json:"file"`
// 	Web                                 WebSearchConfig        `json:"web"`
// 	YoutubeLoaderLanguage               []string               `json:"youtube_loader_language"`
// 	YoutubeLoaderProxyURL               string                 `json:"youtube_loader_proxy_url"`
// 	EmbeddingEngine                     string                 `json:"embedding_engine"`
// 	EmbeddingModel                      string                 `json:"embedding_model"`
// 	HybridBM25Weight                    float64                `json:"hybrid_bm25_weight"`
// 	DatalabMarkerAPIKey                 string                 `json:"datalab_marker_api_key"`
// 	DatalabMarkerLangs                  string                 `json:"datalab_marker_langs"`
// 	DatalabMarkerSkipCache              bool                   `json:"datalab_marker_skip_cache"`
// 	DatalabMarkerForceOCR               bool                   `json:"datalab_marker_force_ocr"`
// 	DatalabMarkerPaginate               bool                   `json:"datalab_marker_paginate"`
// 	DatalabMarkerStripExistingOCR       bool                   `json:"datalab_marker_strip_existing_ocr"`
// 	DatalabMarkerDisableImageExtraction bool                   `json:"datalab_marker_disable_image_extraction"`
// 	DatalabMarkerOutputFormat           string                 `json:"datalab_marker_output_format"`
// 	DatalabMarkerUseLLM                 bool                   `json:"DATALAB_MARKER_USE_LLM"`
// 	DoclingPictureDescriptionMode       string                 `json:"docling_picture_description_mode"`
// 	DoclingPictureDescriptionLocal      map[string]interface{} `json:"docling_picture_description_local"`
// 	DoclingPictureDescriptionAPI        map[string]interface{} `json:"docling_picture_description_api"`
// 	OpenAIAPIBaseURL                    string                 `json:"openai_api_base_url"`
// 	OpenAIAPIKey                        string                 `json:"openai_api_key"`
// 	Ollama                              OllamaRAGConfig        `json:"ollama"`
// 	AzureOpenAI                         AzureOpenAIConfig      `json:"azure_openai"`
// 	EmbeddingBatchSize                  int                    `json:"embedding_batch_size"`
// 	TopKRerankerWeb                     int                    `json:"top_k_reranker_web"`
// 	RelevanceThresholdWeb               float64                `json:"relevance_threshold_web"`
// 	WebLoaderTopK                       int                    `json:"web_loader_top_k"`
// 	WebLoaderTokenLimit                 int                    `json:"web_loader_token_limit"`
// }

// // OllamaRAGConfig Ollama RAG配置
// type OllamaRAGConfig struct {
// 	URL string `json:"url"`
// 	Key string `json:"key"`
// }

// // AzureOpenAIConfig Azure OpenAI配置
// type AzureOpenAIConfig struct {
// 	BaseURL    string `json:"base_url"`
// 	APIKey     string `json:"api_key"`
// 	APIVersion string `json:"api_version"`
// }

// // WebSearchConfig 网页搜索配置
// type WebSearchConfig struct {
// 	Search SearchConfig `json:"search"`
// 	Loader LoaderConfig `json:"loader"`
// }

// // SearchConfig 搜索配置
// type SearchConfig struct {
// 	Enable                       bool         `json:"enable"`
// 	Engine                       string       `json:"engine"`
// 	TrustEnv                     bool         `json:"trust_env"`
// 	ResultCount                  int          `json:"result_count"`
// 	ConcurrentRequests           int          `json:"concurrent_requests"`
// 	Domain                       DomainConfig `json:"domain"`
// 	BypassEmbeddingAndRetrieval  bool         `json:"bypass_embedding_and_retrieval"`
// 	SearxngQueryURL              string       `json:"searxng_query_url"`
// 	YaCyQueryURL                 string       `json:"yacy_query_url"`
// 	YaCyUsername                 string       `json:"yacy_username"`
// 	YaCyPassword                 string       `json:"yacy_password"`
// 	GooglePSEAPIKey              string       `json:"google_pse_api_key"`
// 	GooglePSEEngineID            string       `json:"google_pse_engine_id"`
// 	BraveSearchAPIKey            string       `json:"brave_search_api_key"`
// 	KagiSearchAPIKey             string       `json:"kagi_search_api_key"`
// 	MojeekSearchAPIKey           string       `json:"mojeek_search_api_key"`
// 	BochaSearchAPIKey            string       `json:"bocha_search_api_key"`
// 	SerpstackAPIKey              string       `json:"serpstack_api_key"`
// 	SerpstackHTTPS               bool         `json:"serpstack_https"`
// 	SerperAPIKey                 string       `json:"serper_api_key"`
// 	SerplyAPIKey                 string       `json:"serply_api_key"`
// 	TavilyAPIKey                 string       `json:"tavily_api_key"`
// 	SearchAPIAPIKey              string       `json:"searchapi_api_key"`
// 	SearchAPIEngine              string       `json:"searchapi_engine"`
// 	SerpAPIAPIKey                string       `json:"serpapi_api_key"`
// 	SerpAPIEngine                string       `json:"serpapi_engine"`
// 	JinaAPIKey                   string       `json:"jina_api_key"`
// 	BingSearchV7Endpoint         string       `json:"bing_search_v7_endpoint"`
// 	BingSearchV7SubscriptionKey  string       `json:"bing_search_v7_subscription_key"`
// 	ExaAPIKey                    string       `json:"exa_api_key"`
// 	PerplexityAPIKey             string       `json:"perplexity_api_key"`
// 	SougouAPISID                 string       `json:"sougou_api_sid"`
// 	SougouAPISK                  string       `json:"sougou_api_sk"`
// 	ExternalWebSearchURL         string       `json:"external_web_search_url"`
// 	ExternalWebSearchAPIKey      string       `json:"external_web_search_api_key"`
// 	TavilyExtractDepth           string       `json:"tavily_extract_depth"`
// 	BypassWebLoader              bool         `json:"bypass_web_loader"`
// 	PerplexityModel              string       `json:"perplexity_model"`
// 	PerplexitySearchContextUsage string       `json:"perplexity_search_context_usage"`
// 	TavilyTimeRange              string       `json:"tavily_time_range"`
// 	TavilySearchDepth            string       `json:"tavily_search_depth"`
// 	TavilyChunksPerSource        string       `json:"tavily_chunks_per_source"`
// }

// // DomainConfig 域名配置
// type DomainConfig struct {
// 	FilterList []string `json:"filter_list"`
// 	WhiteList  []string `json:"white_list"`
// }

// // LoaderConfig 加载器配置
// type LoaderConfig struct {
// 	Engine                  string `json:"engine"`
// 	SSLVerification         bool   `json:"ssl_verification"`
// 	PlaywrightWSURL         string `json:"playwright_ws_url"`
// 	PlaywrightTimeout       int    `json:"playwright_timeout"`
// 	FirecrawlAPIKey         string `json:"firecrawl_api_key"`
// 	FirecrawlAPIURL         string `json:"firecrawl_api_url"`
// 	ExternalWebLoaderURL    string `json:"external_web_loader_url"`
// 	ExternalWebLoaderAPIKey string `json:"external_web_loader_api_key"`
// 	ContentCleaning         bool   `json:"content_cleaning"`
// }

// // FileConfig 文件配置
// type FileConfig struct {
// 	MaxSize           int      `json:"max_size"`
// 	MaxCount          int      `json:"max_count"`
// 	AllowedExtensions []string `json:"allowed_extensions"`
// }

// // GoogleDriveConfig Google Drive配置
// type GoogleDriveConfig struct {
// 	Enable bool `json:"enable"`
// }

// // OneDriveConfig OneDrive配置
// type OneDriveConfig struct {
// 	Enable bool `json:"enable"`
// }

// // AppAuthConfig 认证配置（避免与可能的AuthConfig冲突）
// type AppAuthConfig struct {
// 	Admin     AdminConfig  `json:"admin"`
// 	APIKey    APIKeyConfig `json:"api_key"`
// 	JWTExpiry string       `json:"jwt_expiry"`
// }

// // AdminConfig 管理员配置
// type AdminConfig struct {
// 	Show bool `json:"show"`
// }

// // APIKeyConfig API Key配置
// type APIKeyConfig struct {
// 	Enable               bool   `json:"enable"`
// 	EndpointRestrictions bool   `json:"endpoint_restrictions"`
// 	AllowedEndpoints     string `json:"allowed_endpoints"`
// }

// // WebUIConfig WebUI配置
// type WebUIConfig struct {
// 	URL     string `json:"url,omitempty"`     // WebUI地址
// 	Name    string `json:"name,omitempty"`    // 应用名称
// 	Favicon string `json:"favicon,omitempty"` // 网站图标
// }

// // ChannelsConfig 频道配置
// type ChannelsConfig struct {
// 	Enable bool `json:"enable"`
// }

// // NotesConfig 笔记配置
// type NotesConfig struct {
// 	Enable bool `json:"enable"`
// }

// // AudioConfig 音频配置
// type AudioConfig struct {
// 	TTS TTSConfig `json:"tts"`
// 	STT STTConfig `json:"stt"`
// }

// // TTSConfig 文本转语音配置
// type TTSConfig struct {
// 	OpenAI  AppOpenAIAudioConfig `json:"openai"`
// 	APIKey  string               `json:"api_key"`
// 	Engine  string               `json:"engine"`
// 	Model   string               `json:"model"`
// 	Voice   string               `json:"voice"`
// 	SplitOn string               `json:"split_on"`
// 	Azure   AzureTTSConfig       `json:"azure"`
// }

// // AzureTTSConfig Azure TTS配置
// type AzureTTSConfig struct {
// 	SpeechRegion       string `json:"speech_region"`
// 	SpeechBaseURL      string `json:"speech_base_url"`
// 	SpeechOutputFormat string `json:"speech_output_format"`
// }

// // STTConfig 语音转文本配置
// type STTConfig struct {
// 	OpenAI       AppOpenAIAudioConfig `json:"openai"`
// 	Engine       string               `json:"engine"`
// 	Model        string               `json:"model"`
// 	WhisperModel string               `json:"whisper_model"`
// 	Deepgram     DeepgramConfig       `json:"deepgram"`
// 	Azure        AzureSTTConfig       `json:"azure"`
// }

// // DeepgramConfig Deepgram配置
// type DeepgramConfig struct {
// 	APIKey string `json:"api_key"`
// }

// // AzureSTTConfig Azure STT配置
// type AzureSTTConfig struct {
// 	APIKey      string `json:"api_key"`
// 	Region      string `json:"region"`
// 	Locales     string `json:"locales"`
// 	BaseURL     string `json:"base_url"`
// 	MaxSpeakers string `json:"max_speakers"`
// }

// // AppOpenAIAudioConfig OpenAI音频配置
// type AppOpenAIAudioConfig struct {
// 	APIBaseURL string `json:"api_base_url"`
// 	APIKey     string `json:"api_key"`
// }

// // LDAPConfig LDAP配置
// type LDAPConfig struct {
// 	Enable       bool              `json:"enable"`
// 	URL          string            `json:"url,omitempty"`           // LDAP服务器URL
// 	BindDN       string            `json:"bind_dn,omitempty"`       // 绑定DN
// 	BindPassword string            `json:"bind_password,omitempty"` // 绑定密码
// 	UserBase     string            `json:"user_base,omitempty"`     // 用户搜索基础DN
// 	UserFilter   string            `json:"user_filter,omitempty"`   // 用户过滤器
// 	AttributeMap map[string]string `json:"attribute_map,omitempty"` // 属性映射
// }

// // TaskConfig 任务配置
// type TaskConfig struct {
// 	Model              TaskModelConfig              `json:"model"`
// 	Title              TaskTitleConfig              `json:"title"`
// 	Image              TaskImageConfig              `json:"image"`
// 	Autocomplete       TaskAutocompleteConfig       `json:"autocomplete"`
// 	Tags               TaskTagsConfig               `json:"tags"`
// 	Query              TaskQueryConfig              `json:"query"`
// 	Tools              TaskToolsConfig              `json:"tools"`
// 	FollowUp           TaskFollowUpConfig           `json:"follow_up"`
// 	WebSearchDetection TaskWebSearchDetectionConfig `json:"web_search_detection"`
// }

// // TaskModelConfig 任务模型配置
// type TaskModelConfig struct {
// 	Default  string `json:"default"`
// 	External string `json:"external"`
// }

// // TaskTitleConfig 任务标题配置
// type TaskTitleConfig struct {
// 	Enable         bool   `json:"enable"`
// 	PromptTemplate string `json:"prompt_template"`
// }

// // TaskImageConfig 任务图像配置
// type TaskImageConfig struct {
// 	PromptTemplate string `json:"prompt_template"`
// }

// // TaskAutocompleteConfig 任务自动完成配置
// type TaskAutocompleteConfig struct {
// 	Enable         bool `json:"enable"`
// 	InputMaxLength int  `json:"input_max_length"`
// }

// // TaskTagsConfig 任务标签配置
// type TaskTagsConfig struct {
// 	PromptTemplate string `json:"prompt_template"`
// 	Enable         bool   `json:"enable"`
// }

// // TaskQueryConfig 任务查询配置
// type TaskQueryConfig struct {
// 	Search         TaskQuerySearchConfig    `json:"search"`
// 	Retrieval      TaskQueryRetrievalConfig `json:"retrieval"`
// 	PromptTemplate string                   `json:"prompt_template"`
// }

// // TaskQuerySearchConfig 任务查询搜索配置
// type TaskQuerySearchConfig struct {
// 	Enable bool `json:"enable"`
// }

// // TaskQueryRetrievalConfig 任务查询检索配置
// type TaskQueryRetrievalConfig struct {
// 	Enable bool `json:"enable"`
// }

// // TaskToolsConfig 任务工具配置
// type TaskToolsConfig struct {
// 	PromptTemplate string `json:"prompt_template"`
// }

// // TaskFollowUpConfig 任务跟进配置
// type TaskFollowUpConfig struct {
// 	Enable         bool   `json:"enable"`
// 	PromptTemplate string `json:"prompt_template"`
// }

// // TaskWebSearchDetectionConfig 任务网络搜索检测配置
// type TaskWebSearchDetectionConfig struct {
// 	PromptTemplate string `json:"prompt_template"`
// }

// // OllamaConfig Ollama配置
// type OllamaConfig struct {
// 	Enable     bool                       `json:"enable"`
// 	BaseURLs   []string                   `json:"base_urls"`
// 	APIConfigs map[string]OllamaAPIConfig `json:"api_configs"`
// }

// // OllamaAPIConfig Ollama API配置
// type OllamaAPIConfig struct {
// 	// 根据需要添加具体字段
// }

// // DirectConfig 直连配置
// type DirectConfig struct {
// 	Enable bool `json:"enable"`
// }

// // UserConfig 用户配置
// type UserConfig struct {
// 	Permissions UserPermissions `json:"permissions"`
// }

// // UserPermissions 用户权限
// type UserPermissions struct {
// 	Workspace WorkspacePermissions `json:"workspace"`
// 	Sharing   SharingPermissions   `json:"sharing"`
// 	Chat      ChatPermissions      `json:"chat"`
// 	Features  FeaturePermissions   `json:"features"`
// }

// // WorkspacePermissions 工作区权限
// type WorkspacePermissions struct {
// 	Models    bool `json:"models"`
// 	Knowledge bool `json:"knowledge"`
// 	Prompts   bool `json:"prompts"`
// 	Tools     bool `json:"tools"`
// }

// // SharingPermissions 分享权限
// type SharingPermissions struct {
// 	PublicModels    bool `json:"public_models"`
// 	PublicKnowledge bool `json:"public_knowledge"`
// 	PublicPrompts   bool `json:"public_prompts"`
// 	PublicTools     bool `json:"public_tools"`
// }

// // ChatPermissions 聊天权限
// type ChatPermissions struct {
// 	Controls          bool `json:"controls"`
// 	FileUpload        bool `json:"file_upload"`
// 	Delete            bool `json:"delete"`
// 	Edit              bool `json:"edit"`
// 	Share             bool `json:"share"`
// 	Export            bool `json:"export"`
// 	STT               bool `json:"stt"`
// 	TTS               bool `json:"tts"`
// 	Call              bool `json:"call"`
// 	MultipleModels    bool `json:"multiple_models"`
// 	Temporary         bool `json:"temporary"`
// 	TemporaryEnforced bool `json:"temporary_enforced"`
// }

// // FeaturePermissions 功能权限
// type FeaturePermissions struct {
// 	DirectToolServers bool `json:"direct_tool_servers"`
// 	WebSearch         bool `json:"web_search"`
// 	ImageGeneration   bool `json:"image_generation"`
// 	CodeInterpreter   bool `json:"code_interpreter"`
// 	Notes             bool `json:"notes"`
// }

// // EvaluationConfig 评估配置
// type EvaluationConfig struct {
// 	Arena ArenaConfig `json:"arena"`
// }

// // ArenaConfig Arena配置
// type ArenaConfig struct {
// 	Enable bool     `json:"enable"`
// 	Models []string `json:"models"`
// }

// // RecaptchaConfig Recaptcha配置
// type RecaptchaConfig struct {
// 	Enabled     bool              `json:"enabled"`
// 	SiteKey     string            `json:"site_key"`
// 	SecretKey   string            `json:"secret_key"`
// 	Threshold   float64           `json:"threshold"`
// 	V2SiteKey   string            `json:"v2_site_key"`
// 	V2SecretKey string            `json:"v2_secret_key"`
// 	Mode        string            `json:"mode"`
// 	V2Type      string            `json:"v2_type"`
// 	V3          RecaptchaV3Config `json:"v3"`
// 	V2          RecaptchaV2Config `json:"v2"`
// }

// // RecaptchaV3Config Recaptcha V3配置
// type RecaptchaV3Config struct {
// 	SiteKey   string  `json:"site_key"`
// 	SecretKey string  `json:"secret_key"`
// 	Threshold float64 `json:"threshold"`
// }

// // RecaptchaV2Config Recaptcha V2配置
// type RecaptchaV2Config struct {
// 	SiteKey   string `json:"site_key"`
// 	SecretKey string `json:"secret_key"`
// 	Type      string `json:"type"`
// }

// // SMTPConfig SMTP配置
// type SMTPConfig struct {
// 	ForgotPassword ForgotPasswordConfig `json:"forgot_password"`
// 	Host           string               `json:"host"`
// 	Port           int                  `json:"port"`
// 	UseTLS         bool                 `json:"use_tls"`
// 	Username       string               `json:"username"`
// 	Password       string               `json:"password"`
// 	FromEmail      string               `json:"from_email"`
// }

// // ForgotPasswordConfig 忘记密码配置
// type ForgotPasswordConfig struct {
// 	Enable bool `json:"enable"`
// }

// // ContentFilterConfig 内容过滤配置
// type ContentFilterConfig struct {
// 	RegexCheckingWords string `json:"regex_checking_words,omitempty"` // 正则检查词列表
// 	LLMCheckingPrompts string `json:"llm_checking_prompts,omitempty"`
// 	DynamicPrompt      string `json:"dynamic_prompt,omitempty"` // 动态提示配置
// }

// // LangfuseConfig Langfuse配置
// type LangfuseConfig struct {
// 	Enable    bool   `json:"enable"`
// 	SecretKey string `json:"secret_key"`
// 	PublicKey string `json:"public_key"`
// 	Host      string `json:"host"`
// 	AllowList string `json:"allow_list"`
// }

// // DeepResearchConfig 深度研究配置
// type DeepResearchConfig struct {
// 	Enabled            bool   `json:"enabled"`
// 	OpenAIAPIKey       string `json:"openai_api_key"`
// 	OpenAIAPIBase      string `json:"openai_api_base"`
// 	LLMModel           string `json:"llm_model"`
// 	GoogleSearchKey    string `json:"google_search_key"`
// 	GoogleCSEID        string `json:"google_cse_id"`
// 	HKGAISearchURL     string `json:"hkgai_search_url"`
// 	HKGAISearchEnabled bool   `json:"hkgai_search_enabled"`
// 	JinaAPIKey         string `json:"jina_api_key"`
// 	MaxLLMCalls        int    `json:"max_llm_calls"`
// 	MaxTokenLength     int    `json:"max_token_length"`
// 	SearchTimeout      int    `json:"search_timeout"`
// 	LLMTimeout         int    `json:"llm_timeout"`
// 	SystemPrompt       string `json:"system_prompt,omitempty"`        // 系统提示
// 	TaskPrompt         string `json:"task_prompt,omitempty"`          // 任务提示
// 	SearchQueryPrompt  string `json:"search_query_prompt,omitempty"`  // 搜索查询提示
// 	SearchResultPrompt string `json:"search_result_prompt,omitempty"` // 搜索结果提示
// 	RelatedQueryPrompt string `json:"related_query_prompt,omitempty"` // 相关查询提示
// 	SummarizePrompt    string `json:"summarize_prompt,omitempty"`     // 总结提示
// }

// // ToolServerConfig 工具服务器配置
// type ToolServerConfig struct {
// 	Connections []ToolConnection `json:"connections"`
// }

// // ToolConnection 工具连接配置
// type ToolConnection struct {
// 	URL      string                 `json:"url"`
// 	Path     string                 `json:"path"`
// 	AuthType string                 `json:"auth_type"`
// 	Key      string                 `json:"key"`
// 	Config   ToolConnectionConfig   `json:"config"`
// 	Specs    []ToolSpec             `json:"specs"`
// 	Info     ToolInfo               `json:"info,omitempty"`
// 	Paths    map[string]interface{} `json:"paths,omitempty"`
// }

// // ToolInfo 工具信息
// type ToolInfo struct {
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// }

// // ToolSpec 工具规格
// type ToolSpec struct {
// 	Type        string                 `json:"type"`
// 	Name        string                 `json:"name"`
// 	Description string                 `json:"description"`
// 	Parameters  map[string]interface{} `json:"parameters"`
// }

// // ToolConnectionConfig 工具连接配置
// type ToolConnectionConfig struct {
// 	Enable        bool                   `json:"enable"`
// 	AccessControl map[string]interface{} `json:"access_control"`
// }

// // CodeExecutionConfig 代码执行配置
// type CodeExecutionConfig struct {
// 	Enable  bool          `json:"enable"`
// 	Engine  string        `json:"engine"`
// 	Jupyter JupyterConfig `json:"jupyter"`
// }

// // CodeInterpreterConfig 代码解释器配置
// type CodeInterpreterConfig struct {
// 	Enable         bool          `json:"enable"`
// 	Engine         string        `json:"engine"`
// 	PromptTemplate string        `json:"prompt_template"`
// 	Jupyter        JupyterConfig `json:"jupyter"`
// }

// // JupyterConfig Jupyter配置
// type JupyterConfig struct {
// 	URL          string `json:"url"`
// 	Auth         string `json:"auth"`
// 	AuthToken    string `json:"auth_token"`
// 	AuthPassword string `json:"auth_password"`
// 	Timeout      int    `json:"timeout"`
// }

// // ImageGenConfig 图像生成配置
// type ImageGenConfig struct {
// 	Engine   string                 `json:"engine"`
// 	Enable   bool                   `json:"enable"`
// 	Prompt   ImageGenPromptConfig   `json:"prompt"`
// 	OpenAI   ImageGenOpenAIConfig   `json:"openai"`
// 	Gemini   ImageGenGeminiConfig   `json:"gemini"`
// 	Auto1111 ImageGenAuto1111Config `json:"automatic1111"`
// 	ComfyUI  ImageGenComfyUIConfig  `json:"comfyui"`
// }

// // ImageGenPromptConfig 图像生成提示配置
// type ImageGenPromptConfig struct {
// 	Enable bool `json:"enable"`
// }

// // ImageGenOpenAIConfig 图像生成OpenAI配置
// type ImageGenOpenAIConfig struct {
// 	APIBaseURL string `json:"api_base_url"`
// 	APIKey     string `json:"api_key"`
// }

// // ImageGenGeminiConfig 图像生成Gemini配置
// type ImageGenGeminiConfig struct {
// 	APIBaseURL string `json:"api_base_url"`
// 	APIKey     string `json:"api_key"`
// }

// // ImageGenAuto1111Config 图像生成Automatic1111配置
// type ImageGenAuto1111Config struct {
// 	BaseURL   string      `json:"base_url"`
// 	APIAuth   string      `json:"api_auth"`
// 	CFGScale  interface{} `json:"cfg_scale"`
// 	Sampler   interface{} `json:"sampler"`
// 	Scheduler interface{} `json:"scheduler"`
// }

// // ImageGenComfyUIConfig 图像生成ComfyUI配置
// type ImageGenComfyUIConfig struct {
// 	BaseURL  string        `json:"base_url"`
// 	APIKey   string        `json:"api_key"`
// 	Workflow string        `json:"workflow"`
// 	Nodes    []interface{} `json:"nodes"`
// }

// // Banner 横幅配置
// type Banner struct {
// 	ID          string `json:"id"`
// 	Type        string `json:"type"`
// 	Title       string `json:"title"`
// 	Content     string `json:"content"`
// 	Dismissible bool   `json:"dismissible"`
// 	Timestamp   int64  `json:"timestamp"`
// }

// // PromptSuggestion 提示建议
// type PromptSuggestion struct {
// 	Title   map[string]PromptSuggestionText `json:"title"`
// 	Content map[string]string               `json:"content"`
// }

// // PromptSuggestionText 提示建议文本
// type PromptSuggestionText struct {
// 	Main string `json:"main"`
// 	Sub  string `json:"sub"`
// }

// // ConfigSnapshot 配置快照（用于版本管理）
// type ConfigSnapshot struct {
// 	ID        int64             `json:"id"`
// 	Config    *AppConfiguration `json:"config"`
// 	Version   int               `json:"version"`
// 	CreatedAt time.Time         `json:"created_at"`
// 	CreatedBy string            `json:"created_by"`
// 	Comment   string            `json:"comment"`
// }

// // AnyscaleConfig Anyscale配置
// type AnyscaleConfig struct {
// 	Enable      bool                         `json:"enable"`
// 	APIBaseURLs []string                     `json:"api_base_urls,omitempty"`
// 	APIKeys     []string                     `json:"api_keys,omitempty"`
// 	APIConfigs  map[string]AnyscaleAPIConfig `json:"api_configs,omitempty"`
// }

// // AnyscaleAPIConfig Anyscale API配置
// type AnyscaleAPIConfig struct {
// 	Enable         bool     `json:"enable"`
// 	Tags           []string `json:"tags,omitempty"`
// 	PrefixID       string   `json:"prefix_id,omitempty"`
// 	ModelIDs       []string `json:"model_ids,omitempty"`
// 	ConnectionType string   `json:"connection_type,omitempty"`
// }

// // CohereConfig Cohere配置
// type CohereConfig struct {
// 	Enable      bool                       `json:"enable"`
// 	APIBaseURLs []string                   `json:"api_base_urls,omitempty"`
// 	APIKeys     []string                   `json:"api_keys,omitempty"`
// 	APIConfigs  map[string]CohereAPIConfig `json:"api_configs,omitempty"`
// }

// // CohereAPIConfig Cohere API配置
// type CohereAPIConfig struct {
// 	Enable         bool     `json:"enable"`
// 	Tags           []string `json:"tags,omitempty"`
// 	PrefixID       string   `json:"prefix_id,omitempty"`
// 	ModelIDs       []string `json:"model_ids,omitempty"`
// 	ConnectionType string   `json:"connection_type,omitempty"`
// }

// // AnthropicConfig Anthropic配置
// type AnthropicConfig struct {
// 	Enable      bool                          `json:"enable"`
// 	APIBaseURLs []string                      `json:"api_base_urls,omitempty"`
// 	APIKeys     []string                      `json:"api_keys,omitempty"`
// 	APIConfigs  map[string]AnthropicAPIConfig `json:"api_configs,omitempty"`
// }

// // AnthropicAPIConfig Anthropic API配置
// type AnthropicAPIConfig struct {
// 	Enable         bool     `json:"enable"`
// 	Tags           []string `json:"tags,omitempty"`
// 	PrefixID       string   `json:"prefix_id,omitempty"`
// 	ModelIDs       []string `json:"model_ids,omitempty"`
// 	ConnectionType string   `json:"connection_type,omitempty"`
// }

// // LlamaCloudConfig LlamaCloud配置
// type LlamaCloudConfig struct {
// 	Enable     bool                           `json:"enable"`
// 	APIBaseURL string                         `json:"api_base_url,omitempty"`
// 	APIKey     string                         `json:"api_key,omitempty"`
// 	Model      string                         `json:"model,omitempty"`
// 	APIConfigs map[string]LlamaCloudAPIConfig `json:"api_configs,omitempty"`
// }

// // LlamaCloudAPIConfig LlamaCloud API配置
// type LlamaCloudAPIConfig struct {
// 	Enable         bool     `json:"enable"`
// 	Tags           []string `json:"tags,omitempty"`
// 	PrefixID       string   `json:"prefix_id,omitempty"`
// 	ModelIDs       []string `json:"model_ids,omitempty"`
// 	ConnectionType string   `json:"connection_type,omitempty"`
// }

// // ProxyConfig 代理配置
// type ProxyConfig struct {
// 	Enable   bool   `json:"enable"`
// 	URL      string `json:"url,omitempty"`
// 	Username string `json:"username,omitempty"`
// 	Password string `json:"password,omitempty"`
// }
