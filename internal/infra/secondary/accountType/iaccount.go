package accountType

type IAccount interface {
	SetNumberAccount(number string)
	SetAccountType(accountType string)
	GetNumberAccount() string
	GetAccountType() string
}
