package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Object string

const (
	Person     Object = "person"
	Bicycle    Object = "bicycle"
	Car        Object = "car"
	Motorcycle Object = "motorcycle"
	Bus        Object = "bus"
	Train      Object = "train"
	Truck      Object = "truck"
)

type CameraCont struct {
	CameraId   primitive.ObjectID `bson:"camera_id"`
	FromZoneID primitive.ObjectID `bson:"from_zone_id"`
	ToZoneID   primitive.ObjectID `bson:"to_zone_id"`
	CreatedAt  time.Time          `bson:"created_at"`
	Time       time.Time          `bson:"time"`
	Type       Object             `bson:"type"`
	Count      int                `bson:"count"`
}
