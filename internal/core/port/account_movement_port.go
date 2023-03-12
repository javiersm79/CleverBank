package port

type AccountMovementPort interface {
	Deposit(accountId string) (string, error)
}
