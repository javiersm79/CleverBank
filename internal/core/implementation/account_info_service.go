package implementation

import (
	Account "cleverbank/internal/core/domain/account"
	"cleverbank/internal/core/port"
	"fmt"
)

type AccountInfoService struct {
	accountInfoPort port.AccountInfoPort
}

func (c AccountInfoService) Handle(accountNumber string) (Account.AccountDetails, error) {
	response, err := c.accountInfoPort.GetBalance(accountNumber)

	if err != nil {
		fmt.Printf("error AccountMovementService in GetBalance [account: %v], error: %s", accountNumber, err.Error())
		return Account.AccountDetails{}, err
	}

	return response, nil
}
