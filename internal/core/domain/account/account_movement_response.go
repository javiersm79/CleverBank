package account

import "time"

type AccountMovementResponse struct {
	IdMovement           string    `json:"id_movement"`
	Secuence             int64     `json:"secuence"`
	Date                 time.Time `json:"date"`
	SourceAccountNumber  string    `json:"source_account_number,omitempty"`
	DestinyAccountNumber string    `json:"destiny_account_number,omitempty"`
	Description          string    `json:"description"`
	Amount               int64     `json:"amount"`
}
