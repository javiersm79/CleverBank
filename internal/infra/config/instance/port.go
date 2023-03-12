package instance

import (
	"cleverbank/internal/core/port"
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
	return connection
}

func GetAccountMovementPort() port.AccountMovementPort {
	return AccountMovementRepository{}
}

func GetAccountInfoPort() port.AccountInfoPort {
	return AccountInfotRepository{GetDatabaseConextion()}
}
