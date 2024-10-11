package anythingllm

import (
	"bytes"
	"doollm/clients/anythingllm/documents"
	"doollm/clients/anythingllm/system"
	"doollm/utils"
	"encoding/json"
	"fmt"
)

func (c *Client) QueryDocument() (*documents.QueryResponse, error) {
	url := GetRequestUrl("/v1/documents")
	resp, err := utils.SendRequest(c.httpClient, "GET", url, nil, "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data documents.QueryResponse
	if err := utils.ParseResponse(resp, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) DocumentUpload(filePath, extension string) (*documents.UploadResponse, error) {
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

func (c *Client) DocumentUploadFormString(content, fileName, extension string) (*documents.UploadResponse, error) {
	url := GetRequestUrl("/v1/document/upload")
	body, contentType, err := utils.CreateMultipartBodyFromString(content, fileName, extension)
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

func (c *Client) DocumentUploadRowText(params documents.RawTextParams) (*documents.UploadResponse, error) {
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
