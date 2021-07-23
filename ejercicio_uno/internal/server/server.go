package server

import (
	"ejercicio-golang-meli/ejercicio_uno/internal/server/controller/crypto"

	"github.com/gin-gonic/gin"
)

func GetMeliServer(input crypto.CryptoController) *gin.Engine {
	router := gin.Default()
	router.GET("/myapi", input.Request)
	return router
}
