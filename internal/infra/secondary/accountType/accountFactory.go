package accountType

import "fmt"

func GetAccount(accountType string) (IAccount, error) {
	if accountType == "ahorro" {
		return newAhorro(), nil
	}
	if accountType == "corriente" {
		return newCorriente(), nil
	}
	return nil, fmt.Errorf("Wrong account type passed")
}
