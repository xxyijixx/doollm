package anythingllm

import (
	"bytes"
	workspacethread "doollm/clients/anythingllm/workspace_thread"
	"doollm/utils"
	"encoding/json"
	"fmt"
)

func (c *Client) CreateWorkspaceThread(slug string) (*workspacethread.CreateResponse, error) {
	url := GetRequestUrl(fmt.Sprintf("/v1/workspace/%s/thread/new", slug))
	params := map[string]string{}
	params["user_id"] = "1"
	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %w", err)
	}
	fmt.Printf("url: %s, JSON: %s", url, string(jsonData))
	resp, err := utils.SendRequest(c.httpClient, "POST", url, bytes.NewBuffer(jsonData), "application/json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data workspacethread.CreateResponse
	if err := utils.ParseResponse(resp, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetChatsForWorkspaceThread(workspaceSlug, threadSlug string) (*workspacethread.HistoryChatsResponse, error) {
	url := GetRequestUrl(fmt.Sprintf("/v1/workspace/%s/thread/%s/chats", workspaceSlug, threadSlug))
	resp, err := utils.SendRequest(c.httpClient, "GET", url, nil, "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data workspacethread.HistoryChatsResponse
	if err := utils.ParseResponse(resp, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
