package port

type AccountInfoPort interface {
	GetBalance(accountNumber string) (string, error)
}
