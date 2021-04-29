package helpers

import (
	"banking-app/backend/configs"
	"banking-app/backend/interfaces"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
)

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err)

	return string(hashed)
}

func ConnectDB() *gorm.DB {
	dbInfo := "host=" + configs.Config.DbHost +" port=" + configs.Config.DbPort + " user=" + configs.Config.DbUser +" dbname=" + configs.Config.DbName + " password=" + configs.Config.DbPassword +" sslmode=disable"
	log.Printf("DB info: %s", dbInfo)
	db, err := gorm.Open("postgres", dbInfo)
	HandleErr(err)
	return db
}

// func to validate the inputs is valid for registration
func Validation(values []interfaces.Validation) bool{
	username := regexp.MustCompile("^([A-Za-z0-9]{5,})+$")
	email := regexp.MustCompile("^[A-Za-z0-9]+[@]+[A-Za-z0-9]+[.]+[A-Za-z]+$")
	for i := 0; i < len(values); i++ {
		switch values[i].Valid {
		case "username":
			if !username.MatchString(values[i].Value) {
				return false
			}
		case "email":
			if !email.MatchString(values[i].Value) {
				return false
			}
		case "password":
			if len(values[i].Value) < 5 {
				return false
			}
		}
	}

	return true
}



