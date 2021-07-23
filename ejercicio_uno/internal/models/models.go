package models

type CryptoResponse struct {
	ID         string     `json:"id"`
	MarketData MarketData `json:"market_data"`
}

type MarketData struct {
	CurrentPrice CurrentPrice `json:"current_price"`
}

type CurrentPrice struct {
	Usd      float64 `json:"usd"`
}

type MyApiResponse struct {
	ID      string  `json:"id"`
	Content Content `json:"content"`
	Partial bool    `json:"partial"`
}

type Content struct {
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}
