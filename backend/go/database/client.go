package database

import (
	"log"

	"github.com/dimo92/devops-challenge-01/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Connector variable used for CRUD operation's
var Connector *gorm.DB

//Connect creates MySQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}

//Migrate create/updates database table
// func Migrate(table ...*gorm.Model) {
// 	Connector.AutoMigrate(&table)
// 	log.Println("Table migrated")
// }

func Migrate(product *entity.Product, user *entity.User) {
	Connector.AutoMigrate(&product, &user)
	log.Println("Table migrated")
}
