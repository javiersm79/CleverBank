package instance

import (
	"cleverbank/internal/core/port"
	"cleverbank/internal/infra/secondary/persistence"
	"cleverbank/internal/infra/secondary/persistence/dto"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func GetDatabaseConextion() *gorm.DB {
	databasename := "cleverbank"
	database := "postgres"
	databasepassword := "postgres"
	databaseurl := "postgres://postgres:" + databasepassword + "@localhost/" + databasename + "?sslmode=disable"

	connection, err := gorm.Open(database, databaseurl)
	if err != nil {
		log.Fatalln("Invalid database url")
	}
	sqldb := connection.DB()

	err = sqldb.Ping()
	if err != nil {
		log.Fatal("Database connected")
	}
	fmt.Println("Database connection successful.")
	InitialMigration(connection)
	return connection
}

var connectionDB = GetDatabaseConextion()

func GetAccountMovementPort() port.AccountMovementPort {
	return persistence.AccountMovementRepository{connectionDB}
}

func GetAccountInfoPort() port.AccountInfoPort {
	return persistence.AccountInfoRepository{connectionDB}
}

func InitialMigration(connection *gorm.DB) {
	connection.AutoMigrate(dto.Client{})
	connection.AutoMigrate(dto.User{})
	connection.AutoMigrate(dto.Account{})
	connection.AutoMigrate(dto.MovementAccount{})
}
