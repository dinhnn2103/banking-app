package interfaces

import (
	"time"
)

// we separate user and account tables because 1 user can have many accounts :D
type UserLoginInfo struct {
	ID          uint
	Accountname    string
	Accountpass    string
	Email 		string
	Token       string
	Expire time.Time
}

type LoginResponse struct {
	Token  string
	Expire time.Time
}

type Account struct {
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

type Transaction struct {
	From   uint
	To     uint
	Amount int
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
}

type ErrResponse struct {
	Message string
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
