package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `bson: "_id"`
	UserId       string             `json: "user_id"`
	FirstName    string             `json: "first_name" validate: "required, min=2, max=50"`
	LastName     string             `json: "first_name" validate: "required, min=2, max=50"`
	Email        string             `json: "email" validate: "email, required"`
	Phone        string             `json: "phone" validate: "required"`
	Password     string             `json: "password" validate: "required, min=6"`
	UserType     string             `json: "user_type" validate: "required, eq=admin|eq=user"`
	Token        string             `json: "token"`
	RefreshToken string             `json: "refresh_token"`
	CreatedAt    time.Time          `json: "created_at"`
	UpdatedAt    time.Time          `json: "updated_at"`
}
