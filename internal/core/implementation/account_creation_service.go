package implementation

import (
	Account "cleverbank/internal/core/domain/account"
	"cleverbank/internal/core/port"
)

type AccountCreationService struct {
	accountCreationPort port.AccountCreationPort
}

func (c AccountCreationService) Handle(newAccountReq Account.NewAccountReq) (Account.AccountDetails, error) {
	response, err := c.accountCreationPort.CreateAccount(newAccountReq)

	if err != nil {
		//fmt.Printf("error AccountMovementService in CreateAccount [req: %v], error: %s", newAccountReq, err.Error())
		return Account.AccountDetails{}, err
	}

	return response, nil
}
