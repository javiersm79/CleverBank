package implementation

import (
	"cleverbank/internal/core/port"
	"fmt"
)

type AccountMovementService struct {
	accountMovementPort port.AccountMovementPort
}

func (c AccountMovementService) Handle(accountId string) (string, error) {
	response, err := c.accountMovementPort.Deposit(accountId)

	if err != nil {
		fmt.Printf("error AccountMovementService in Handle [account: %v], error: %s", accountId, err.Error())
		return "", err
	}

	return response, nil
}
