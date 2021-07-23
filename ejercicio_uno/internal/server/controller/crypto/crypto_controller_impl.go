package crypto

import (
	"ejercicio-golang-meli/ejercicio_uno/internal/api/service"
	"ejercicio-golang-meli/ejercicio_uno/internal/server/controller/crypto/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CryptoController struct {
	CryptoService service.CryptoService
}

func NewCryptoController(service service.CryptoService) *CryptoController {
	return &CryptoController{
		CryptoService: service,
	}
}

func (controller *CryptoController) Request(c *gin.Context) {
	data := dto.Input{
		Id: c.Query("id"),
		Currency: c.Query("currency"),
	}
	response, err := controller.CryptoService.GetPrice(data.Id, data.Currency)

	if err != nil {
		c.JSON(http.StatusPartialContent, dto.BuildPartialResponse(data.Id))
		return
	}

	c.JSON(http.StatusOK, response)
}
