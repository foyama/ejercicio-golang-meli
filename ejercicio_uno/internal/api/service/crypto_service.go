package service

import "ejercicio-golang-meli/ejercicio_uno/internal/api/service/dto"

type CryptoService interface {
	GetPrice(id string, currency string) (*dto.CryptoResponse, error)
}