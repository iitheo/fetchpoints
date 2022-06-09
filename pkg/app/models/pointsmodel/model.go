package pointsmodel

import "time"

type TransactionDB struct {
	ID        string    `json:"ID"`
	Payer     string    `json:"payer"`
	Points    int       `json:"points"`
	Timestamp time.Time `json:"timestamp"`
}
