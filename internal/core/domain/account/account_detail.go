package account

import "time"

type AccountDetails struct {
	Id        int64      `json:"id,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	ClientId  int64      `json:"client_id,omitempty"`
	Type      string     `json:"accountType,omitempty"`
	Number    string     `json:"number,omitempty"`
	Balance   int64      `json:"balance,omitempty"`
}
