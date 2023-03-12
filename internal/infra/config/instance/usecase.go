package instance

import (
	"cleverbank/internal/core/implementation"
	"cleverbank/internal/core/usecase"
	"github.com/jinzhu/gorm"
)

func GetAccountMovementUseCase() usecase.AccountMovementUseCase {
	return implementation.GetAccountMovementService(GetAccountMovementPort())
}

func GetAccountInfoUseCase() usecase.AccountMovementUseCase {
	return implementation.GetAccountInfoService(GetAccountInfoPort())
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
type AccountMovementRepository struct {
}

func (amr AccountMovementRepository) Deposit(accountId string) (string, error) {
	return "movementID-Generated", nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
type AccountInfotRepository struct {
	dbconex *gorm.DB
}

func (air AccountInfotRepository) GetBalance(accountNumber string) (string, error) {
	var dbuser User
	air.dbconex.Where("email = ?", "user@test.com").First(&dbuser)
	return dbuser.Password, nil
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
