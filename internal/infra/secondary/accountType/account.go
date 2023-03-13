package accountType

type Account struct {
	//ClientId int64
	Type   string
	Number string
}

func (a *Account) SetNumberAccount(number string) {
	a.Number = number
}

func (a *Account) GetNumberAccount() string {
	return a.Number
}

func (a *Account) SetAccountType(accountype string) {
	a.Type = accountype
}

func (a *Account) GetAccountType() string {
	return a.Type
}
