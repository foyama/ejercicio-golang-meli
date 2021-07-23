package controller

import (
	"ejercicio-golang-meli/ejercicio_uno/internal/models"
	"ejercicio-golang-meli/ejercicio_uno/internal/server/http"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

func MyApi() {
	r := gin.Default()
	r.GET("/myapi", func(c *gin.Context) {
		body, err := http.DoGet("https://api.coingecko.com/api/v3/coins/bitcoin")
		if err != nil {
			log.Fatalln(err)
		}
		var crypto models.CryptoResponse
		err = json.Unmarshal(body, &crypto)
		if err != nil {
			log.Fatalln(err)
		}
		myApiResponse := parseCryptoToMyApi(&crypto)
		if myApiResponse.Partial {
			c.JSON(200, myApiResponse)
		} else {
			c.JSON(206, myApiResponse)
		}
	})
	r.Run(":8080")
}

func parseCryptoToMyApi(cryptoResponse *models.CryptoResponse) *models.MyApiResponse {
	obj := &models.MyApiResponse{
		ID: cryptoResponse.ID,
		Partial: true,
	}
	if (cryptoResponse.MarketData != models.MarketData{} && cryptoResponse.MarketData.CurrentPrice != models.CurrentPrice{}) {
		obj.Partial = false
		obj.Content = models.Content{
			Price: cryptoResponse.MarketData.CurrentPrice.Usd,
			Currency: "usd",
		}
	}
	return obj
}
