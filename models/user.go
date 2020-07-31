package models

import (
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

// Read a user
func ReadUser(email string) User {

	var user User
	Db.Debug().Where("email = ?", email).Find(&user)
	return user
}
