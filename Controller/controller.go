package controller

import (
	"QuizChallenge/Config/db"
	model "QuizChallenge/Model"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HomeLink handles request for Home
func HomeLink(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Welcome to QuizChallenge!")
}

// CreateUser handles request for Sign up
func CreateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var user model.User
	body, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(body, &user)
	var res model.UserResponseModel

	if err != nil {
		res.Message = "Error While Creating User, Try Again"
		res.Success = false
		json.NewEncoder(response).Encode(res)
		return
	}

	collection, err := db.GetDatabaseCollection("Users")
	if err != nil {
		res.Message = "Error While Creating User, Try Again"
		res.Success = false
		json.NewEncoder(response).Encode(res)
		return
	}

	var result model.User
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = collection.FindOne(ctx, bson.D{{"username", user.Username}}).Decode(&result)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			insertResult, err := collection.InsertOne(ctx, user)
			if err != nil {
				res.Message = "Error While Creating User, Try Again"
				res.Success = false
				json.NewEncoder(response).Encode(res)
				return
			}
			user.ID = insertResult.InsertedID.(primitive.ObjectID)
			res.Data = user
			res.Message = "User signed up successfully!"
			res.Success = true
			response.WriteHeader(http.StatusCreated)
			json.NewEncoder(response).Encode(res)
			return
		}

		res.Message = "Error While Creating User, Try Again"
		res.Success = false
		response.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(response).Encode(res)
		return
	}

	res.Message = "Username already Exists!!"
	res.Success = false
	response.WriteHeader(http.StatusNotAcceptable)
	json.NewEncoder(response).Encode(res)
	return
}

// LoginUser handles request for Sign in
func LoginUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user model.User
	body, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	collection, err := db.GetDatabaseCollection("Users")

	if err != nil {
		log.Fatal(err)
	}
	var result model.User
	var res model.UserResponseModel

	err = collection.FindOne(context.TODO(), bson.D{{"email", user.Email}}).Decode(&result)

	if err != nil {
		res.Message = "Invalid Email"
		res.Success = false
		json.NewEncoder(response).Encode(res)
		return
	}

	if result.Password != user.Password {
		res.Message = "Invalid password"
		res.Success = false
		json.NewEncoder(response).Encode(res)
		return
	}

	res.Message = "Login Success!"
	res.Success = true
	res.Data = result
	json.NewEncoder(response).Encode(res)
}

// func UserProfile(response http.ResponseWriter, request *http.Request) {
// 	response.Header().Set("Content-Type", "application/json")
// 	tokenString := request.Header.Get("Authorization")
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Don't forget to validate the alg is what you expect:
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method")
// 		}
// 		return []byte("secret"), nil
// 	})
// 	var result model.User
// 	var res model.ResponseResult
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		result.Username = claims["username"].(string)
// 		result.FirstName = claims["firstname"].(string)
// 		result.LastName = claims["lastname"].(string)
// 		json.NewEncoder(response).Encode(result)
// 		return
// 	} else {
// 		res.Error = err.Error()
// 		json.NewEncoder(response).Encode(res)
// 		return
// 	}
// }

// InstructionHandler handles request for getting Instructions
func InstructionHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var res model.InstructionResponseModel

	data, err := ioutil.ReadFile("/Users/ankit.gohel/src/golang-book/src/QuizChallange/Resource/instruction.txt")
	if err != nil {
		res.Message = "Something went wrong!"
		res.Success = false
		return
	}

	res.Message = "Instructions fetched successfully!"
	res.Success = true
	res.Data = string(data)

	json.NewEncoder(response).Encode(res)
}
