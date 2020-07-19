package main

import (
	"QuizChallenge/Config/db"
	router "QuizChallenge/Router"
	"fmt"
)

func main() {
	fmt.Println("Starting the Quiz Challenge...")

	db.InitialiseDatabase()
	router.InitialiseRouter()
}
