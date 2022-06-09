package pointsdto

import "time"

type AddPointsDTO struct {
	Payer     string `json:"payer"`
	Points    int    `json:"points"`
	Timestamp string `json:"timestamp"`
}

type SpendPointsRequestDTO struct {
	Points int `json:"points"`
}

type SpendPointsResponseDTO struct {
	Payer  string `json:"payer"`
	Points int    `json:"points"`
}

type GetPointsBalanceDTO struct {
	Payer  string `json:"payer"`
	Points int    `json:"points"`
}

type GetAllPointsDTO struct {
	Payer     string    `json:"payer"`
	Points    int       `json:"points"`
	Timestamp time.Time `json:"timestamp"`
}
