package coingecko_service

import (
	"ejercicio-golang-meli/ejercicio_uno/internal/api/service/dto"
	"ejercicio-golang-meli/ejercicio_uno/pkg/http"
	"encoding/json"
	coinGeckoDto "ejercicio-golang-meli/ejercicio_uno/internal/api/service/coingecko_service/dto"
	"strings"
)

type CoinGeckoService struct {
	BaseUrl string
}

func NewCoinGeckoService(url string) *CoinGeckoService {
	return &CoinGeckoService{url}
}

func (s *CoinGeckoService) GetPrice(id string, currency string) (*dto.CryptoResponse, error) {
	url := s.BaseUrl + id
	bodyByte, err := http.DoGet(url)
	if err != nil {
		return nil, err
	}

	response, err := buildResponse(bodyByte, currency)
	return response, err
}

func buildResponse(bodyByte []byte, currency string) (*dto.CryptoResponse, error) {
	var cryptoResponse dto.CryptoResponse
	var out coinGeckoDto.CoinGeckoResponse
	err := json.Unmarshal(bodyByte, &out)

	if err != nil {
		return nil, err
	}

	cryptoResponse.Id = out.Id
	cryptoResponse.Content.Price = out.MarketData.CurrentPrice[currency]
	cryptoResponse.Content.Currency = strings.ToUpper(currency)
	cryptoResponse.Partial = false

	return &cryptoResponse, nil
}
