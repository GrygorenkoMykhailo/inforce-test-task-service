package apis

import "time"

type RaribleApiErrorResponse struct {
	Code string
	Message string
}

type GetNftOwnershipsByIdResponse struct {
	ID            string    `json:"id"`
	Blockchain    string    `json:"blockchain"`
	ItemID        string    `json:"itemId"`
	Contract      string    `json:"contract"`
	Collection    string    `json:"collection"`
	TokenID       int       `json:"tokenId"`
	Owner         string    `json:"owner"`
	Value         int       `json:"value"`
	Source        string    `json:"source"`
	CreatedAt     time.Time `json:"createdAt"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
	LazyValue     int       `json:"lazyValue"`
	Pending       []struct {
		Type      string `json:"@type"`
		Royalties []struct {
			Account string `json:"account"`
			Value   int    `json:"value"`
		} `json:"royalties,omitempty"`
		From string `json:"from,omitempty"`
	} `json:"pending"`
	BestSellOrder struct {
		ID                string    `json:"id"`
		Fill              float64   `json:"fill"`
		Platform          string    `json:"platform"`
		Status            string    `json:"status"`
		StartedAt         time.Time `json:"startedAt"`
		EndedAt           time.Time `json:"endedAt"`
		MakeStock         float64   `json:"makeStock"`
		Cancelled         bool      `json:"cancelled"`
		OptionalRoyalties bool      `json:"optionalRoyalties"`
		CreatedAt         time.Time `json:"createdAt"`
		LastUpdatedAt     time.Time `json:"lastUpdatedAt"`
		DbUpdatedAt       time.Time `json:"dbUpdatedAt"`
		MakePrice         float64   `json:"makePrice"`
		TakePrice         float64   `json:"takePrice"`
		MakePriceUsd      float64   `json:"makePriceUsd"`
		TakePriceUsd      float64   `json:"takePriceUsd"`
		Maker             string    `json:"maker"`
		Taker             string    `json:"taker"`
		Make              struct {
			Type struct {
				Blockchain string `json:"blockchain"`
				Contract   string `json:"contract"`
				Type       string `json:"@type"`
			} `json:"type"`
			Value float64 `json:"value"`
		} `json:"make"`
		Take struct {
			Type struct {
				Blockchain string `json:"blockchain"`
				Contract   string `json:"contract"`
				Type       string `json:"@type"`
			} `json:"type"`
			Value float64 `json:"value"`
		} `json:"take"`
		Salt      string   `json:"salt"`
		Signature string   `json:"signature"`
		FeeTakers []string `json:"feeTakers"`
		Data      struct {
			Type string `json:"@type"`
			Data struct {
			} `json:"data"`
		} `json:"data"`
		Version int `json:"version"`
	} `json:"bestSellOrder"`
	OriginOrders []struct {
		Origin        string `json:"origin"`
		BestSellOrder struct {
			ID                string    `json:"id"`
			Fill              float64   `json:"fill"`
			Platform          string    `json:"platform"`
			Status            string    `json:"status"`
			StartedAt         time.Time `json:"startedAt"`
			EndedAt           time.Time `json:"endedAt"`
			MakeStock         float64   `json:"makeStock"`
			Cancelled         bool      `json:"cancelled"`
			OptionalRoyalties bool      `json:"optionalRoyalties"`
			CreatedAt         time.Time `json:"createdAt"`
			LastUpdatedAt     time.Time `json:"lastUpdatedAt"`
			DbUpdatedAt       time.Time `json:"dbUpdatedAt"`
			MakePrice         float64   `json:"makePrice"`
			TakePrice         float64   `json:"takePrice"`
			MakePriceUsd      float64   `json:"makePriceUsd"`
			TakePriceUsd      float64   `json:"takePriceUsd"`
			Maker             string    `json:"maker"`
			Taker             string    `json:"taker"`
			Make              struct {
				Type struct {
					Blockchain string `json:"blockchain"`
					Contract   string `json:"contract"`
					Type       string `json:"@type"`
				} `json:"type"`
				Value float64 `json:"value"`
			} `json:"make"`
			Take struct {
				Type struct {
					Blockchain string `json:"blockchain"`
					Contract   string `json:"contract"`
					Type       string `json:"@type"`
				} `json:"type"`
				Value float64 `json:"value"`
			} `json:"take"`
			Salt      string   `json:"salt"`
			Signature string   `json:"signature"`
			FeeTakers []string `json:"feeTakers"`
			Data      struct {
				Type string `json:"@type"`
				Data struct {
				} `json:"data"`
			} `json:"data"`
			Version int `json:"version"`
		} `json:"bestSellOrder"`
		BestBidOrder struct {
			ID                string    `json:"id"`
			Fill              float64   `json:"fill"`
			Platform          string    `json:"platform"`
			Status            string    `json:"status"`
			StartedAt         time.Time `json:"startedAt"`
			EndedAt           time.Time `json:"endedAt"`
			MakeStock         float64   `json:"makeStock"`
			Cancelled         bool      `json:"cancelled"`
			OptionalRoyalties bool      `json:"optionalRoyalties"`
			CreatedAt         time.Time `json:"createdAt"`
			LastUpdatedAt     time.Time `json:"lastUpdatedAt"`
			DbUpdatedAt       time.Time `json:"dbUpdatedAt"`
			MakePrice         float64   `json:"makePrice"`
			TakePrice         float64   `json:"takePrice"`
			MakePriceUsd      float64   `json:"makePriceUsd"`
			TakePriceUsd      float64   `json:"takePriceUsd"`
			Maker             string    `json:"maker"`
			Taker             string    `json:"taker"`
			Make              struct {
				Type struct {
					Blockchain string `json:"blockchain"`
					Contract   string `json:"contract"`
					Type       string `json:"@type"`
				} `json:"type"`
				Value float64 `json:"value"`
			} `json:"make"`
			Take struct {
				Type struct {
					Blockchain string `json:"blockchain"`
					Contract   string `json:"contract"`
					Type       string `json:"@type"`
				} `json:"type"`
				Value float64 `json:"value"`
			} `json:"take"`
			Salt      string   `json:"salt"`
			Signature string   `json:"signature"`
			FeeTakers []string `json:"feeTakers"`
			Data      struct {
				Type string `json:"@type"`
				Data struct {
				} `json:"data"`
			} `json:"data"`
			Version int `json:"version"`
		} `json:"bestBidOrder"`
	} `json:"originOrders"`
	Version int `json:"version"`
}

type NtfTrait struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Rarity int `json:"rarity"`
}

type GetNftTraitsRarityResponse struct {
	Continuation string `json:"continuation"`
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
