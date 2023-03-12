package mapper

import (
	"cleverbank/internal/core/domain/account"
	"cleverbank/internal/infra/secondary/persistence/dto"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func MapAccountMovementDtoToDomain(accountDto dto.Account, movementRequest account.AccountMovementRequest) account.AccountMovementResponse {

	descriptionInfo := GenerateDescription(movementRequest)

	accountDetails := account.AccountMovementResponse{
		IdMovement:           uuid.NewString(),
		Secuence:             int64(accountDto.ID),
		Date:                 time.Now(),
		SourceAccountNumber:  movementRequest.SourceAccountNumber,
		DestinyAccountNumber: movementRequest.DestinyAccountNumber,
		Description:          descriptionInfo,
		Amount:               movementRequest.Amount,
	}
	return accountDetails
}

func GenerateDescription(movementRequest account.AccountMovementRequest) string {
	var descriptionInfo string
	switch movementRequest.Type {
	case "deposit":
		descriptionInfo = fmt.Sprintf("Deposit - Account number %s by an amount of %d", movementRequest.DestinyAccountNumber, movementRequest.Amount)
	default:
		descriptionInfo = "Operation not found"
	}
	return descriptionInfo
}
