package implementation

import (
	"cleverbank/internal/core/domain/auth"
	"cleverbank/internal/core/port"
	"fmt"
)

type AuthLoginService struct {
	authLoginPort port.AuthLoginPort
}

func (c AuthLoginService) Handle(loginReq auth.LoginReq) (auth.LoginResponse, error) {
	response, err := c.authLoginPort.Login(loginReq)

	if err != nil {
		fmt.Printf("error AuthLoginService in Login - error: %s", err.Error())
		return auth.LoginResponse{}, err
	}

	return response, nil
}
