package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Auth ...
type Auth struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

// Customer ...
type Customer struct {
	ID               primitive.ObjectID `bson:"_id"`
	Username         string             `bson:"username"`
	FirstName        string             `bson:"firstname"`
	RegistrationDate time.Time          `bson:"registrationDate"`
	ReleaseDate      time.Time          `bson:"releaseDate"`
}

func testCostomer() Auth {
	return Auth{
		Username: "umekovazamat@gmail.com",
		Password: "123456",
	}
}
