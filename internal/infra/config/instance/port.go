package instance

import "cleverbank/internal/core/port"

func GetAccountMovementPort() port.AccountMovementPort {
	return AccountMovementRepository{}
}

func GetAccountInfoPort() port.AccountInfoPort {
	return AccountInfotRepository{}
}
