package implementation

import (
	"cleverbank/internal/core/domain/account"
	"cleverbank/internal/core/port"
	"fmt"
)

type AccountMovementService struct {
	accountMovementPort port.AccountMovementPort
}

func (c AccountMovementService) Handle(movementRequest account.AccountMovementRequest) (account.AccountMovementResponse, error) {
	response, err := c.accountMovementPort.Deposit(movementRequest)

	if err != nil {
		fmt.Printf("error AccountMovementService in Handle [request: %v], error: %s", movementRequest, err.Error())
		return account.AccountMovementResponse{}, err
	}

	return response, nil
}
