package usecase

type AccountInfoUseCase interface {
	Handle(accountNumber string) (string, error)
}
