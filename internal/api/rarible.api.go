package api

import (
	"net/http"
	"os"
	"time"
)

type RaribleApi interface {
	GetNftOwnershipsById(id string) (*GetNftOwnershipsByIdResponse, error)
}

type raribleApiImpl struct {}

func NewRaribleApi() RaribleApi {
	return &raribleApiImpl{}
}

type GetNftOwnershipsByIdCreator struct {
	Account string `json:"account"`
	Value   int    `json:"value"`
}

type GetNftOwnershipsByIdResponse struct {
	ID            string    `json:"id"`
	Blockchain    string    `json:"blockchain"`
	ItemID        string    `json:"itemId"`
	Contract      string    `json:"contract"`
	Collection    string    `json:"collection"`
	TokenID       string    `json:"tokenId"`
	Owner         string    `json:"owner"`
	Value         string    `json:"value"`
	CreatedAt     time.Time `json:"createdAt"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
	Creators      []GetNftOwnershipsByIdCreator `json:"creators"`
	LazyValue    	string        `json:"lazyValue"`
	Pending      	[]interface{} `json:"pending"`
	OriginOrders 	[]interface{} `json:"originOrders"`
	Version      	int           `json:"version"`
}

func (api *raribleApiImpl) GetNftOwnershipsById(id string) (*GetNftOwnershipsByIdResponse, error) {
	raribleApiBaseUrl := os.Getenv("RARIBLE_API_URL")
	http.DefaultClient

	req, err := http.NewRequest("GET", raribleApiBaseUrl + "/ownerships/" + id, nil)
	if err != nil {

	}
}