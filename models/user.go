package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB

type User struct {
	gorm.Model
	Name     string
	Email    string
	Meta     string
	Password string
}

// Create User in Database
func CreateUser(name, email, password string) {

	user := &User{Name: name, Email: email, Password: password}

	Db.Debug().Create(user)
}
