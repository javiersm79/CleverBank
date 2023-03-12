package usecase

type AccountMovementUseCase interface {
	Handle(movement string) (string, error)
}
