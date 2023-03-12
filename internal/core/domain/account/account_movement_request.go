package account

type AccountMovementRequest struct {
	Type                 string `json:"type"`
	SourceAccountNumber  string `json:"source_account_number"`
	DestinyAccountNumber string `json:"destiny_account_number"`
	Amount               int64  `json:"amount"`
}
