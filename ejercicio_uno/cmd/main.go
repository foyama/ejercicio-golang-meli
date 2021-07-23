package main

import (
	"ejercicio-golang-meli/ejercicio_uno/internal/api/service/coingecko_service"
	"ejercicio-golang-meli/ejercicio_uno/internal/server"
	"ejercicio-golang-meli/ejercicio_uno/internal/server/controller/crypto"
)


func main() {
	url := "https://api.coingecko.com/api/v3/coins/"
	service := coingecko_service.NewCoinGeckoService(url)
	controller := crypto.NewCryptoController(service)
	router := server.GetMeliServer(*controller)
	router.Run(":8080")
}