package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represent user to signup
type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}

// UserResponseModel will give result for User Response
type UserResponseModel struct {
	Data    User   `json:"data,omitempty"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

// InstructionResponseModel will give result for Instructions Response
type InstructionResponseModel struct {
	Data    string `json:"data,omitempty"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}
