package qiandun

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Client 千盾通讯客户端
type Client struct {
	config *Config
}

// NewClient 创建千盾客户端
func NewClient(config *Config) *Client {
	return &Client{
		config: config,
	}
}

// getBaseURL 获取基础URL
func (c *Client) getBaseURL() string {
	if c.config.BaseURL != "" {
		return c.config.BaseURL
	}
	if c.config.IsProd {
		return "https://api.qiandun365.com/api"
	}
	return "https://api.pre-qiandun365.com/api"
}

// getHeadersString 获取参与签名的请求头字符串
func (c *Client) getHeadersString(req *http.Request) string {
	signatureHeaders := req.Header.Get("X-Ca-Signature-Headers")
	if signatureHeaders == "" {
		return ""
	}

	headers := strings.Split(signatureHeaders, ",")
	var headerPairs []string
	for _, header := range headers {
		header = strings.TrimSpace(header)
		value := req.Header.Get(header)
		headerPairs = append(headerPairs, fmt.Sprintf("%s:%s", header, value))
	}
	return strings.Join(headerPairs, "\n")
}

// getUrlString 获取参与签名的URL字符串
func (c *Client) getUrlString(req *http.Request) string {
	// 获取查询参数
	query := req.URL.Query()

	// 按字典序排序
	keys := make([]string, 0, len(query))
	for k := range query {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 拼接字符串
	var builder strings.Builder
	builder.WriteString(req.URL.Path)

	if len(keys) > 0 {
		builder.WriteString("?")
		for i, k := range keys {
			if i > 0 {
				builder.WriteString("&")
			}
			builder.WriteString(k)
			builder.WriteString("=")
			builder.WriteString(query.Get(k))
		}
	}

	return builder.String()
}

// calculateSignature 计算签名 - 使用阿里云API网关标准签名算法
func (c *Client) calculateSignature(req *http.Request) string {
	// 1. 准备签名字符串
	stringToSign := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s",
		req.Method,
		req.Header.Get("Accept"),
		req.Header.Get("Content-MD5"),
		req.Header.Get("Content-Type"),
		req.Header.Get("Date"),
		c.getHeadersString(req),
		c.getUrlString(req),
	)

	// 2. 使用HMAC-SHA256计算签名
	h := hmac.New(sha256.New, []byte(c.config.AppSecret))
	h.Write([]byte(stringToSign))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return signature
}

// SendRequest 发送HTTP请求
func (c *Client) SendRequest(path string, requestData interface{}) (*QiandunResponse, error) {
	// 构建请求体
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %v", err)
	}

	// 构建完整URL
	url := c.getBaseURL() + path

	// 创建HTTP请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("创建HTTP请求失败: %v", err)
	}

	// 设置通用请求头
	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())
	nonce := uuid.New().String()

	req.Header.Set("X-Ca-Key", c.config.AppKey)
	req.Header.Set("X-Ca-Timestamp", timestamp)
	req.Header.Set("X-Ca-Nonce", nonce)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// 设置签名头
	signatureHeaders := []string{"X-Ca-Key", "X-Ca-Nonce", "X-Ca-Timestamp"}
	req.Header.Set("X-Ca-Signature-Headers", strings.Join(signatureHeaders, ","))

	// 计算Content-MD5
	md5Hash := md5.Sum(requestBody)
	contentMD5 := base64.StdEncoding.EncodeToString(md5Hash[:])
	req.Header.Set("Content-MD5", contentMD5)

	// 计算签名
	signature := c.calculateSignature(req)
	req.Header.Set("X-Ca-Signature", signature)

	// 打印调试信息
	log.Printf("=== 千盾请求开始 ===")
	log.Printf("URL: %s", url)
	log.Printf("Method: POST")
	log.Printf("Headers:")
	for key, values := range req.Header {
		log.Printf("  %s: %v", key, values)
	}
	log.Printf("Body: %s", string(requestBody))
	log.Printf("=== 千盾请求结束 ===")

	// 发送请求
	startTime := time.Now()

	// 配置HTTP传输，跳过TLS证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送HTTP请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	// 打印响应调试信息
	log.Printf("=== 千盾响应开始 ===")
	log.Printf("Status: %s", resp.Status)
	log.Printf("Headers:")
	for key, values := range resp.Header {
		log.Printf("  %s: %v", key, values)
	}
	log.Printf("Body: %s", string(respBody))
	log.Printf("Cost Time: %v", time.Since(startTime))
	log.Printf("=== 千盾响应结束 ===")

	// 解析响应
	var qiandunResp QiandunResponse
	if err := json.Unmarshal(respBody, &qiandunResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	// 添加耗时信息
	qiandunResp.CostTime = time.Since(startTime).Milliseconds()

	return &qiandunResp, nil
}
