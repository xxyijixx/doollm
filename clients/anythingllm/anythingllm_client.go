package anythingllm

import (
	"bytes"
	"doollm/clients/anythingllm/authentication"
	"doollm/clients/anythingllm/documents"
	"doollm/clients/anythingllm/system"
	"doollm/config"
	"doollm/utils"
	"encoding/json"
	"fmt"
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

func (c *Client) ValidToken() (*authentication.AuthResponse, error) {
	url := GetRequestUrl("/v1/auth")
	utils.SendRequest(c.httpClient, "GET", url, nil, "")
	resp, err := c.httpClient.Get(url)
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

func (c *Client) UploadFile(filePath, extension string) (*documents.UploadResponse, error) {
	url := GetRequestUrl("/v1/document/upload")
	body, contentType, err := utils.CreateMultipartBody(filePath, extension)
	if err != nil {
		return nil, err
	}

	resp, err := utils.SendRequest(c.httpClient, "POST", url, body, contentType)
	if err != nil {
		return nil, err
	}

	var data documents.UploadResponse
	if err := utils.ParseResponse(resp, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) UploadRowText(params documents.RawTextParams) (*documents.UploadResponse, error) {
	url := GetRequestUrl("/v1/document/raw-text")
	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %w", err)
	}

	resp, err := utils.SendRequest(c.httpClient, "POST", url, bytes.NewBuffer(jsonData), "application/json")
	if err != nil {
		return nil, err
	}

	var data documents.UploadResponse
	if err := utils.ParseResponse(resp, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) RemoveDocument(params system.RemoveDocumentParams) error {
	url := GetRequestUrl("/v1/system/remove-documents")
	jsonData, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	resp, err := utils.SendRequest(c.httpClient, "DELETE", url, bytes.NewBuffer(jsonData), "application/json")
	if err != nil {
		return err
	}

	var data map[string]string
	if err := utils.ParseResponse(resp, &data); err != nil {
		return err
	}
	return nil
}

func GetRequestUrl(uri string) string {
	return config.EnvConfig.LLM_SERVER_URL + uri
}
