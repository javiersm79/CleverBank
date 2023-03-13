package mapper

import (
	"cleverbank/internal/core/domain/account"
	"cleverbank/internal/infra/secondary/persistence/dto"
	"fmt"
	"time"
)

func MapAccountMovementDtoToDomain(movementAccount dto.MovementAccount) account.AccountMovementResponse {

	descriptionInfo := GenerateDescription(movementAccount)

	accountDetails := account.AccountMovementResponse{
		IdMovement:           movementAccount.MovementId,
		Secuence:             int64(movementAccount.ID),
		Date:                 time.Now(),
		SourceAccountNumber:  movementAccount.SourceAccountNumber,
		DestinyAccountNumber: movementAccount.DestinyAccountNumber,
		Description:          descriptionInfo,
		Amount:               movementAccount.Amount,
	}
	return accountDetails
}

func GenerateDescription(movementAccount dto.MovementAccount) string {
	var descriptionInfo string
	switch movementAccount.Action {
	case "deposit":
		descriptionInfo = fmt.Sprintf("Deposit - Account number %s by an amount of %d", movementAccount.DestinyAccountNumber, movementAccount.Amount)
	case "withdrawal":
		descriptionInfo = fmt.Sprintf("Withdrawal - Account number %s by an amount of %d", movementAccount.SourceAccountNumber, movementAccount.Amount)
	case "transfer":
		descriptionInfo = fmt.Sprintf("Transfer - From Account number %s To Account number %s by an amount of %d", movementAccount.SourceAccountNumber, movementAccount.DestinyAccountNumber, movementAccount.Amount)

	default:
		descriptionInfo = "Operation not found"
	}
	return descriptionInfo
}
