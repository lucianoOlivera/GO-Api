package orm

import (
	"fmt"

	"../config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	CreatConnection()

}

func CreatConnection() {
	if GetConnection() == nil {
		return
	}
	dsn := config.GetUrl()
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("conecion exitosa ")
		db = connection
	}
}
func GetConnection() *gorm.DB {
	return db
}
func CreateTables() {
	db.AutoMigrate(&User{})
}

func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()
	return user
}
