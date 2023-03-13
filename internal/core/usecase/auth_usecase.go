package usecase

import (
	"cleverbank/internal/core/domain/auth"
)

type LoginUseCase interface {
	Handle(loginReq auth.LoginReq) (auth.LoginResponse, error)
}
