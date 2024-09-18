package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	First_Name   string             `bson:"first_name"`
	Last_Name    string             `bson:"last_name""`
	Email        string             `bson:"email"`
	Username     string             `bson:"username"`
	Password     string             `bson:"password"`
	Role         string             `bson:"role"`
	Photo        string             `bson:"photo"`
	Verified     bool               `bson:"verified"`
	IsSuperAdmin bool               `bson:"is_super_admin"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"upadated_at"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	FetchByPageNumber(c context.Context, page int) ([]User, error)
	FetchByTime(c context.Context, startTime, endTime *time.Time) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByUsername(c context.Context, username string) (User, error)
	GetByID(c context.Context, id string) (User, error)
}
