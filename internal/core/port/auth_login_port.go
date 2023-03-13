package port

import (
	"cleverbank/internal/core/domain/auth"
)

type AuthLoginPort interface {
	Login(req auth.LoginReq) (auth.LoginResponse, error)
}
