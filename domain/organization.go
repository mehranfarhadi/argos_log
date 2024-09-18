package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionOrganization = "organization"
)

type UserRole string

const (
	Admin  UserRole = "admin"
	Guest  UserRole = "guest"
	Member UserRole = "member"
)

type UserInOrganization struct {
	UserID    primitive.ObjectID `bson:"_id"`
	Role      UserRole           `bson:"role"`
	CreatedAt time.Time          `bson:"create_at"`
}

type Organization struct {
	ID                  primitive.ObjectID   `bson:"_id"`
	Name                string               `bson:"name"`
	Address             string               `bson:"address"`
	PhoneNumber         string               `bson:"phone_number"`
	CreatedAt           time.Time            `bson:"created_at"`
	UpdateAt            time.Time            `bson:"update_at"`
	UserInOrganization  []UserInOrganization `bson:"user_in_organization"`
	GroupInOrganization []UserInOrganization `bson:"group_in_organization"`
}

type GroupInOrganization struct {
	ID       primitive.ObjectID   `bson:"_id"`
	Name     string               `bson:"name"`
	UserID   []primitive.ObjectID `bson:"user_id"`
	CameraID []primitive.ObjectID `bson:"camera_id"`
}

type OrganizationRepository interface {
	Create(c context.Context, user *Organization) error
	Fetch(c context.Context) ([]Organization, error)
	FetchByPageNumber(c context.Context, page int) ([]User, error)
	FetchByTime(c context.Context, startTime, endTime *time.Time) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByUsername(c context.Context, username string) (User, error)
	GetByID(c context.Context, id string) (User, error)
}
