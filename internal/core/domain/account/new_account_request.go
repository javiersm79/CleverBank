package account

type NewAccountReq struct {
	ClientId int64  `json:"client_id,omitempty"`
	Type     string `json:"account_type,omitempty"`
}
