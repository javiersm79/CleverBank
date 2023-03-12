package instance

import "cleverbank/internal/core/port"

func GetAccountMovementPort() port.AccountMovementPort {
	return AccountMovementRepository{}
}
