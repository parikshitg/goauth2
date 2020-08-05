package models

import (
	"encoding/json"
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

	m := make(map[string]interface{})
	m["Github"] = ""
	m["Linkedin"] = ""
	m["Twitter"] = ""

	v, err := json.Marshal(m)
	if err != nil {
		log.Println("Marshal error: ", err)
		return
	}

	user := &User{Name: name, Email: email, Password: password, Meta: string(v)}

	Db.Debug().Create(&user)
}

// Check if user exists in database
func ExistingUser(email string) (User, bool) {

	var user User
	Db.Debug().Where("email = ?", email).Find(&user)
	if user == (User{}) {
		return User{}, false
	}
	return user, true
}

// Read All Users from the database
func UsersTable() []User {

	var users []User
	Db.Debug().Select("id, name, email, meta").Find(&users)

	return users
}

// Set Users New Password in Database
func SetNewPass(email, pass string) {
	Db.Debug().Table("users").Where("email = ?", email).Update("password", pass)
}

// Check if user details by id
func UserDetails(id string) (User, bool) {

	var user User
	Db.Debug().Where("id = ?", id).Find(&user)
	if user == (User{}) {
		return User{}, false
	}
	return user, true
}
