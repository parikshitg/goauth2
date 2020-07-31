package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB

type User struct {
	ID       int `gorm:"primary_key";"AUTO_INCREMENT"`
	Name     string
	Email    string
	Meta     string
	Password string
}

// Create User in Database
func CreateUser(name, email, password string) {

	user := &User{Name: name, Email: email, Password: password}

	Db.Debug().Create(&user)
}

// Check if user exists in database
func ExistingUser(email string) (string, bool) {

	var user User
	Db.Debug().Where("email = ?", email).Find(&user)
	if user == (User{}) {
		return "", false
	}
	return user.Password, true
}

// Read All Users from the database
func UsersTable() []User {

	var users []User
	Db.Debug().Select("id, name, email").Find(&users)

	return users
}
