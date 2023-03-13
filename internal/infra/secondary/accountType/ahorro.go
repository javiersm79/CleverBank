package accountType

import (
	"cleverbank/internal/infra/utils"
	"strconv"
)

type Ahorro struct {
	Account
}

func newAhorro() IAccount {
	return &Ahorro{
		Account: Account{
			Type:   "ahorro",
			Number: "001-" + strconv.Itoa(utils.GenerateNumber()) + "-" + strconv.Itoa(utils.GenerateNumber2()),
		},
	}
}
