package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	First_name   *string            `json:"first_name"`
	Last_name    *string            `json:"last_name"`
	Age          *int               `json:"age"`
	Sex          *string            `json:"sex"`
	PhoneNumber  *string            `json:"phone_number"`
	Login        *string            `json:"login"`
	Password     *string            `json:"password"`
	Email        *string            `json:"email"`
	Token        *string            `json:"token"`
	RefreshToken *string            `json:"refresh_token"`
	User_type    *string            `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	User_id      string             `json:"user_id"`
	Created_at   time.Time          `json:"created_at"`
	// Role     *string             `json:"role"`

}
