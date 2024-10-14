package anythingllm

import (
	"doollm/clients/anythingllm/authentication"
	"doollm/config"
	"doollm/utils"
	"net/http"
)

// CustomTransport 是自定义的 RoundTripper，用于为每个请求添加 Headers
type CustomTransport struct {
	Transport http.RoundTripper
	Headers   map[string]string
}

// RoundTrip 实现 http.RoundTripper 接口，自动为每个请求添加自定义的 Headers
func (c *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}
	return c.Transport.RoundTrip(req)
}

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	headers := map[string]string{
		"Accept":        "application/json",
		"Authorization": "Bearer " + config.EnvConfig.LLM_TOKEN,
	}
	transport := &CustomTransport{
		Transport: http.DefaultTransport,
		Headers:   headers,
	}
	return &Client{
		httpClient: &http.Client{
			Transport: transport,
		},
	}
}

// ValidToken 验证Toekn
func (c *Client) ValidToken() (*authentication.AuthResponse, error) {
	url := GetRequestUrl("/v1/auth")
	resp, err := utils.SendRequest(c.httpClient, "GET", url, nil, "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data authentication.AuthResponse
	if err := utils.ParseResponse(resp, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetRequestUrl 获取请求地址
func GetRequestUrl(uri string) string {
	return config.EnvConfig.LLM_SERVER_URL + uri
}
