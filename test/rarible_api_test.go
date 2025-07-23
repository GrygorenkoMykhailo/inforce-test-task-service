package test

import (
	"encoding/json"
	"inforce-test-task-service/internal/apis"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRaribleApi_TestGetNftOwnershipById(t *testing.T) {
	expected := apis.GetNftOwnershipsByIdResponse{
		ID:         "test-id",
		Blockchain: "ETHEREUM",
		ItemID:     "ETHEREUM:0x1234:1",
		Contract:   "0x1234",
		Collection: "0x1234",
		TokenID:    "1",
		Owner:      "0xowner",
		Value:      "1",
		CreatedAt:  time.Now(),
		LastUpdatedAt: time.Now(),
		Creators:   []apis.GetNftOwnershipsByIdCreator{},
		LazyValue:  "0",
		Pending:    []interface{}{},
		OriginOrders: []interface{}{},
		Version:    1,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/ownerships/test-id", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(expected)
	}))
	defer server.Close()

	client := apis.RaribleApiImpl{ 
		ApiKey: "test-key",
		ApiBaseUrl: server.URL,
		Client: &http.Client{},
	}

	actual, err := client.GetNftOwnershipsById("test-id")
	assert.NoError(t, err)
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Owner, actual.Owner)
	assert.Equal(t, expected.Value, actual.Value)
}

func TestRaribleApi_TestGetNftTraitsRarity(t *testing.T) {
	expected := apis.GetNftTraitsRarityResponse{
		Traits: []apis.NtfTrait{
			{
				Key: "foo",
				Value: "bar",
				Rarity: "rare",
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/items/traits/rarity", r.URL.Path)

		var payload apis.GetNftTraitsRarityPayload
		err := json.NewDecoder(r.Body).Decode(&payload)
		assert.NoError(t, err)
		assert.Equal(t, "foo", payload.Properties[0].Key)
		assert.Equal(t, "bar", payload.Properties[0].Value)

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(expected)
		assert.NoError(t, err)
	}))
	defer server.Close()

	api := &apis.RaribleApiImpl{
		ApiKey: "test-key",
		ApiBaseUrl: server.URL,
		Client:  server.Client(),
	}

	payload := &apis.GetNftTraitsRarityPayload{
		CollectionId: "",
		Properties: []apis.GetNftTraitsRarityProperties{
			{
				Key: "foo",
				Value: "bar",
			},
		},
	}

	resp, err := api.GetNftTraitsRarity(payload)
	assert.NoError(t, err)
	assert.Equal(t, &expected, resp)
}
