package router

import (
	controller "QuizChallenge/Controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitialiseRouter() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controller.HomeLink)
	router.HandleFunc("/signup", controller.CreateUser).Methods("POST")
	router.HandleFunc("/login", controller.LoginUser).Methods("POST")
	router.HandleFunc("/instructions", controller.InstructionHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
