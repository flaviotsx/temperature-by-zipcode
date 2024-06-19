package web

import "github.com/flaviotsx/temperature-by-zipcode/internal/entity"

type GetTemperatureByZipCodeUseCaseInterface interface {
	Execute(zipcode string) (*entity.GetTemperatureByZipcodeResponse, *entity.ErrorResponse)
}
