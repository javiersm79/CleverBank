package accountType

import (
	"cleverbank/internal/infra/utils"
	"strconv"
)

type Corriente struct {
	Account
}

func newCorriente() IAccount {
	return &Corriente{
		Account: Account{
			Type:   "corriente",
			Number: "002-" + strconv.Itoa(utils.GenerateNumber()) + "-" + strconv.Itoa(utils.GenerateNumber2()),
		},
	}
}
