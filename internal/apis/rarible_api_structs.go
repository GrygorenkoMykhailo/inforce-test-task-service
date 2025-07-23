package apis

import "time"

type RaribleApiErrorResponse struct {
	Code string
	Message string
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

type NtfTrait struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Rarity string `json:"rarity"`
}

type GetNftTraitsRarityResponse struct {
	Traits []NtfTrait `json:"traits"`
}

type GetNftTraitsRarityProperties struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type GetNftTraitsRarityPayload struct {
	CollectionId string `json:"collectionId"`
	Properties []GetNftTraitsRarityProperties `json:"properties"`
}
