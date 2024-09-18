package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionCamera = "Camera"
)

type cameraType string

const (
	ArgV3 cameraType = "ArgV4"
	ArgV4 cameraType = "ArgV4"
)

type Camera struct {
	ID           primitive.ObjectID `bson:"_id"`
	CameraID     string             `bson:"CameraID"`
	CameraType   cameraType         `bson:"CameraType"`
	CreatedAt    time.Time          `bson:"create_at"`
	UpdatedAt    time.Time          `bson:"update_at"`
	Latitude     float64            `bson:"latitude"`
	Longitude    float64            `bson:"longitude"`
	CameraConfig CameraConfig       `bson:"camera_config"`
	Created_By   primitive.ObjectID `bson:"created_by"`
}

type CameraConfig struct {
	ID                primitive.ObjectID `bson:"_id"`
	SatelliteImage    []byte             `bson:"satellite_image"`
	CameraImage       []byte             `bson:"camera_image"`
	CameraCalibration [][]int            `bson:"camera_calibration"`
}

type Zone struct {
	ID            primitive.ObjectID `bson:"_id"`
	Name          string             `bson:"name"`
	Lat           float32            `bson:"lat"`
	Long          float32            `bson:"long"`
	XCoordination float32            `bson:"x_coordination"`
	YCoordination float32            `bson:"y_coordination"`
	ColorCode     string             `bson:"color_code"`
	Rotation      int16              `bson:"rotation"`
}
