package api

import (
	"banking-app/backend/configs"
	"encoding/json"
	"fmt"
	"banking-app/backend/interfaces"
	"io/ioutil"
	"log"
	"net/http"
	"banking-app/backend/helpers"
	"banking-app/backend/users"
	"github.com/gorilla/mux"
)

type Login struct {
	Username string
	Password string
}

type ErrResponse struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)
	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)
	// Check the response == all is fine or not
	if login["result"] == true {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := ErrResponse{Message: fmt.Sprintf("%v", login["message"])}
		json.NewEncoder(w).Encode(resp)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	// Read body
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)
	// Handle registration
	var formattedBody interfaces.Register
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)
	// Prepare response
	if register["result"] == true {
		resp := register
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := ErrResponse{Message: fmt.Sprintf("%v", register["message"])}
		json.NewEncoder(w).Encode(resp)
	}
}

// Login API
func StartApi() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	log.Printf("DB info: %s", configs.DBInfo)
	fmt.Println("App is working on port :8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
