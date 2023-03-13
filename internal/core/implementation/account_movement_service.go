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
	var response account.AccountMovementResponse
	var err error
	switch movementRequest.Type {
	case "deposit":
		response, err = c.accountMovementPort.Deposit(movementRequest)
	case "withdrawal":
		response, err = c.accountMovementPort.Withdrawal(movementRequest)
	case "transfer":
		response, err = c.accountMovementPort.Transfer(movementRequest)
	default:
		//descriptionInfo = "Operation not found"
	}

	if err != nil {
		fmt.Printf("error AccountMovementService in Handle [request: %v], error: %s", movementRequest, err.Error())
		return account.AccountMovementResponse{}, err
	}

	return response, nil
}
