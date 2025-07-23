package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"inforce-test-task-service/internal/apis"
	"inforce-test-task-service/internal/config"
	"inforce-test-task-service/internal/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockRaribleApi struct {
	GetOwnershipsFunc func(id string) (*apis.GetNftOwnershipsByIdResponse, error)
	GetTraitsFunc     func(payload *apis.GetNftTraitsRarityPayload) (*apis.GetNftTraitsRarityResponse, error)
}

func (m *mockRaribleApi) GetNftOwnershipsById(id string) (*apis.GetNftOwnershipsByIdResponse, error) {
	return m.GetOwnershipsFunc(id)
}

func (m *mockRaribleApi) GetNftTraitsRarity(payload *apis.GetNftTraitsRarityPayload) (*apis.GetNftTraitsRarityResponse, error) {
	return m.GetTraitsFunc(payload)
}

func TestRaribleController_GetNftOwnershipsById(t *testing.T) {
	mockApi := &mockRaribleApi{
		GetOwnershipsFunc: func(id string) (*apis.GetNftOwnershipsByIdResponse, error) {
			if id == "exists" {
				return &apis.GetNftOwnershipsByIdResponse{
					ID:    "exists",
					Owner: "0xowner",
				}, nil
			}
			return nil, fmt.Errorf("not found")
		},
	}


	logger := config.NewLogger()
	deps := &config.AppDependencies{
		RaribleApi: mockApi,
		Logger:     logger,
	}

	controller := controllers.NewRaribleController(deps)

	r := gin.New()
	r.GET("/rarible/nft-ownership/:collection_id", controller.GetNftOwnershipsById)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/rarible/nft-ownership/exists", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"id":"exists"`)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/rarible/nft-ownership/", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "404 page not found")

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/rarible/nft-ownership/notfound", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "something went wrong")
}

func TestRaribleController_GetNftTraitsRarity(t *testing.T) {
	mockApi := &mockRaribleApi{
		GetTraitsFunc: func(payload *apis.GetNftTraitsRarityPayload) (*apis.GetNftTraitsRarityResponse, error) {
			if payload.CollectionId == "exists" && len(payload.Properties) > 0 &&
				payload.Properties[0].Key == "foo" && payload.Properties[0].Value == "bar" {
				return &apis.GetNftTraitsRarityResponse{
					Traits: []apis.NtfTrait{
						{
							Key:    "foo",
							Value:  "bar",
							Rarity: 1,
						},
					},
				}, nil
			}
			return nil, fmt.Errorf("not found")
		},
	}

	logger := config.NewLogger()
	deps := &config.AppDependencies{
		RaribleApi: mockApi,
		Logger:     logger,
	}

	controller := controllers.NewRaribleController(deps)

	r := gin.New()
	r.POST("/nft-rarity", controller.GetNftTraitsRarity)

	validBody, _ := json.Marshal(apis.GetNftTraitsRarityPayload{
		CollectionId: "exists",
		Properties: []apis.GetNftTraitsRarityProperties{
			{ Key: "foo", Value: "bar" },
		},
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/nft-rarity", bytes.NewReader(validBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp apis.GetNftTraitsRarityResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Len(t, resp.Traits, 1)
	assert.Equal(t, "foo", resp.Traits[0].Key)
	assert.Equal(t, "bar", resp.Traits[0].Value)
	assert.Equal(t, 1, resp.Traits[0].Rarity)

	w2 := httptest.NewRecorder()
	badReq, _ := http.NewRequest(http.MethodPost, "/nft-rarity", bytes.NewReader([]byte("{bad json")))
	badReq.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w2, badReq)

	assert.Equal(t, http.StatusBadRequest, w2.Code)
	assert.Contains(t, w2.Body.String(), "invalid request body")

	badPayloadBody, _ := json.Marshal(apis.GetNftTraitsRarityPayload{
		CollectionId: "not-exist",
		Properties: []apis.GetNftTraitsRarityProperties{
			{ Key: "foo", Value: "bar" },
		},
	})

	w3 := httptest.NewRecorder()
	badReq2, _ := http.NewRequest(http.MethodPost, "/nft-rarity", bytes.NewReader(badPayloadBody))
	badReq2.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w3, badReq2)

	assert.Equal(t, http.StatusInternalServerError, w3.Code)
	assert.Contains(t, w3.Body.String(), "something went wrong")
}
