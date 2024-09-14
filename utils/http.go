// utils/http.go
package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// SendRequest 发送 HTTP 请求并返回响应
func SendRequest(httpClient *http.Client, method, url string, body io.Reader, contentType string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close() // 确保在非200响应时关闭响应体
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}
	return resp, nil
}

// ParseResponse 解析响应体到目标结构体
func ParseResponse(resp *http.Response, result interface{}) error {
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}
	return nil
}

// CreateMultipartBody 创建 multipart/form-data 请求体
func CreateMultipartBody(filePath string) (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return nil, "", fmt.Errorf("error creating form file: %w", err)
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()
	if _, err := io.Copy(part, file); err != nil {
		return nil, "", fmt.Errorf("error copying file: %w", err)
	}
	if err := writer.Close(); err != nil {
		return nil, "", fmt.Errorf("error closing writer: %w", err)
	}
	return body, writer.FormDataContentType(), nil
}
