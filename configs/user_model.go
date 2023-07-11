package configs

import "go.mongodb.org/mongo-driver/bson/primitive"

// User Model to represent my application data
type User struct {
	UserID   primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Location string             `json:"location,omitempty" validate:"required"`
	Title    string             `json:"title,omitempty" validate:"required"`
}
