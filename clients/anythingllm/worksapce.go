package anythingllm

import (
	"bytes"
	"doollm/clients/anythingllm/workspace"
	"doollm/utils"
	"encoding/json"
	"fmt"
)

func (c *Client) CreateWorkspace(params workspace.CreateParams) (*workspace.CreateResponse, error) {
	url := GetRequestUrl("/v1/workspace/new")
	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %w", err)
	}
	resp, err := utils.SendRequest(c.httpClient, "POST", url, bytes.NewBuffer(jsonData), "application/json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data workspace.CreateResponse
	if err := utils.ParseResponse(resp, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) QueryWorkspaces() (*workspace.ListResponse, error) {
	url := GetRequestUrl("/v1/workspaces")
	resp, err := utils.SendRequest(c.httpClient, "GET", url, nil, "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data workspace.ListResponse
	if err := utils.ParseResponse(resp, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// QueryWorkspace Get a workspace by its unique slug.
func (c *Client) QueryWorkspace(slug string) (*workspace.GetResponse, error) {
	url := GetRequestUrl("/v1/workspaces/" + slug)
	resp, err := utils.SendRequest(c.httpClient, "GET", url, nil, "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data workspace.GetResponse
	if err := utils.ParseResponse(resp, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// DeleteWorkspace Deletes a workspace by its slug.
func (c *Client) DeleteWorkspace(slug string) error {
	url := GetRequestUrl("/v1/workspaces/" + slug)
	resp, err := utils.SendRequest(c.httpClient, "DELETE", url, nil, "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// Add or remove documents from a workspace by its unique slug.
func (c *Client) UpdateEmbeddings(slug string, params workspace.UpdateEmbeddingsParams) (*workspace.UpdateEmbeddingsResponse, error) {
	url := GetRequestUrl("/v1/workspace/" + slug + "/update-embeddings")
	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %w", err)
	}
	resp, err := utils.SendRequest(c.httpClient, "POST", url, bytes.NewBuffer(jsonData), "application/json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data workspace.UpdateEmbeddingsResponse
	if err := utils.ParseResponse(resp, &data); err != nil {
		return nil, err
	}
	return &data, nil

}
