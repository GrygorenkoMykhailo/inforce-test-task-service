package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RaribleApi interface {
	GetNftOwnershipsById(id string) (*GetNftOwnershipsByIdResponse, error)
	GetNftTraitsRarity(payload *GetNftTraitsRarityPayload) (*GetNftTraitsRarityResponse, error)
}

type RaribleApiImpl struct {
	ApiKey string
	ApiBaseUrl string
	Client *http.Client
}

func NewRaribleApi() RaribleApi {
	return &RaribleApiImpl{ 
		ApiKey: os.Getenv("RARIBLE_API_KEY"), 
		ApiBaseUrl: os.Getenv("RARIBLE_API_URL"),
		Client: &http.Client{},
	}
}

func (api *RaribleApiImpl) doRequest(method, path string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, api.ApiBaseUrl + path, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("X-API-KEY", api.ApiKey)
	if method == http.MethodPost {
		req.Header.Set("content-type", "application/json")
	}

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		var apiErr RaribleApiErrorResponse
		err := json.Unmarshal(respBody, &apiErr)
		if err != nil {
			return nil, fmt.Errorf("http %d: failed to parse error body: %w", resp.StatusCode, err)
		}
		return nil, fmt.Errorf("http %d: %s - %s", resp.StatusCode, apiErr.Code, apiErr.Message)
	}

	return respBody, nil
}

func (api *RaribleApiImpl) GetNftOwnershipsById(id string) (*GetNftOwnershipsByIdResponse, error) {
	body, err := api.doRequest(http.MethodGet, "/ownerships/" + id, nil)
	if err != nil {
		return nil, err
	}

	var result GetNftOwnershipsByIdResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ownerships response: %w", err)
	}

	return &result, nil
}

func (api *RaribleApiImpl) GetNftTraitsRarity(payload *GetNftTraitsRarityPayload) (*GetNftTraitsRarityResponse, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal traits payload: %w", err)
	}

	body, err := api.doRequest(http.MethodPost, "/items/traits/rarity", bytes.NewReader(jsonPayload))
	if err != nil {
		return nil, err
	}

	var result GetNftTraitsRarityResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal traits response: %w", err)
	}

	return &result, nil
}