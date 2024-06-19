package entity

type Current struct {
	TempC float64 `json:"temp_C"`
}

type WeatherApiResponse struct {
	Current `json:"current"`
}
