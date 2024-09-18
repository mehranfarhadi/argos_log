package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionIncident = "incident"
)

type Incident struct {
	ID        primitive.ObjectID `bson:"_id"`
	CameraID  primitive.ObjectID `bson:"camera_id"`
	Image     []byte             `bson:"image"`
	Video     string             `bson:"video"`
	Message   string             `bson:"message"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
