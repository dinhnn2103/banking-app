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
	AccountName string
	AccountPass string
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello I am from login api")
	var formattedBody Login
	body := readBody(r)
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	login := users.Login(formattedBody.AccountName, formattedBody.AccountPass)
	// Check the response == all is fine or not
	apiResponse(login, w)
}

func register(w http.ResponseWriter, r *http.Request) {
	// Handle registration
	var formattedBody interfaces.Register
	body := readBody(r)
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)
	// Prepare response
	apiResponse(register, w)
}

// Login API
func StartApi() {
	router := mux.NewRouter()
	router.Use(helpers.PanicHandler)
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	log.Printf("DB info: %s", configs.DBInfo)
	fmt.Println("App is working on port :8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}

func readBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)
	return body
}

func apiResponse(call map[string]interface{}, w http.ResponseWriter) {
	if call["result"] == true {
		resp := call
		json.NewEncoder(w).Encode(resp)
		// Handle error in else
	} else {
		resp := interfaces.ErrResponse{Message: fmt.Sprintf("%v", call["message"])}
		json.NewEncoder(w).Encode(resp)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	user := users.GetUser(auth)
	apiResponse(user, w)
}
