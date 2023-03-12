package implementation

import (
	"cleverbank/internal/core/port"
	"fmt"
)

type AccountInfoService struct {
	accountInfoPort port.AccountInfoPort
}

func (c AccountInfoService) Handle(accountNumber string) (string, error) {
	response, err := c.accountInfoPort.GetBalance(accountNumber)

	if err != nil {
		fmt.Printf("error AccountMovementService in GetBalance [account: %v], error: %s", accountNumber, err.Error())
		return "", err
	}

	return response, nil
}
