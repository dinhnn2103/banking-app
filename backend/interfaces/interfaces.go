package interfaces

import (
	"github.com/jinzhu/gorm"
)

// we separate user and account tables because 1 user can have many accounts :D
type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

type ResponseAccount struct {
	ID      uint
	Name    string
	Balance int
}

type ResponseUser struct {
	ID       uint
	Username string
	Email    string
	Accounts []ResponseAccount
}

type Validation struct {
	Value string
	Valid string
}

type Register struct {
	Username string
	Email    string
	Password string
}
