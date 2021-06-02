package users

import (
	"banking-app/backend/database"
	"log"
	"time"
	"banking-app/backend/helpers"
	"banking-app/backend/interfaces"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func Register(username string, email string, pass string) map[string]interface{} {
	// Add validation to registration
	valid := helpers.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: email, Valid: "email"},
			{Value: pass, Valid: "password"},
		})
	if valid {
		generatedPassword := helpers.HashAndSalt([]byte(pass))
		user := &interfaces.UserLoginInfo{Accountname: username, Email: email, Accountpass: generatedPassword}
		database.DBconnection.Create(&user)
		var response = prepareResponse(user, false)
		return response
	} else {
		return map[string]interface{}{
			"message": "Cannot register, please check username & email & password",
			"result": false,
		}
	}
}

func Login(username string, pass string) map[string]interface{} {
	// Add validation to login function
	log.Printf("accountName : %s", username)
	log.Printf("accountPass : %s", pass)
	isValid := helpers.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: pass, Valid: "password"},
		})
	if isValid {
		user := &interfaces.UserLoginInfo{}
		// First check user exist or not
		database.DBconnection.Raw("SELECT id, accountname, accountpass, email, token, expire FROM users WHERE accountName = ?", username).Scan(&user)
		if user.Accountname == "" {
			return map[string]interface{}{
				"message": "User not found",
				"result": false,
			}
		}
		// Password check
		passErr := bcrypt.CompareHashAndPassword([]byte(user.Accountpass), []byte(pass))
		if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
			return map[string]interface{}{
				"message": "Wrong password",
				"result": false,
			}
		}
		// get user account info from db
		withToken := helpers.ValidateTokenOnly(user.Token)
		if time.Now().Before(user.Expire) {
			withToken = true
		}
		return prepareResponse(user, withToken)
	} else {
		return 	map[string]interface{}{"message": "Validation step: invalid values"}
	}
}

func prepareToken(user *interfaces.UserLoginInfo) string {
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry": time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleErr(err)

	return token
}

func prepareResponse(user *interfaces.UserLoginInfo, withToken bool) map[string]interface{} {
	responseUser := &interfaces.ResponseUser {
		ID: user.ID,
		Username: user.Accountpass,
		Email: user.Email,
	}
	var response = map[string]interface{}{
		"message": "OK",
		"result": true,
	}
	if withToken {
		var token = prepareToken(user)
		response["jwt"] = token
	}
	response["data"] = responseUser

	return response
}

func GetUser(jwt string) map[string]interface{} {
	isValid := helpers.ValidateTokenOnly(jwt)
	if isValid {
		user := &interfaces.UserLoginInfo{}
		database.DBconnection.Raw("SELECT id, accountName, accountPass, email, token, expire FROM users WHERE token = ?", jwt).Scan(&user)
		if user.Accountpass == "" {
			return map[string]interface{}{
				"message": "User not found",
				"result": false,
			}
		}
		var response = prepareResponse(user, false)
		return response
	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}
}

