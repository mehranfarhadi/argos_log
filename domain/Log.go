package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionLog = "log"
)

type Log struct {
	ID            primitive.ObjectID `bson:"_id"`
	CameraID      primitive.ObjectID `bson:"camera_id"`
	Ram           float32            `bson:"ram"`
	Fan           float32            `bson:"fan"`
	BoardTemp     float32            `bson:"board_temp"`
	Cpu           float32            `bson:"cpu"`
	PowerMode     float32            `bson:"power_mode"`
	BattV         float32            `bson:"batt_V"`
	PanelV        float32            `bson:"panel_V"`
	LoadI         float32            `bson:"load_I"`
	LoadV         float32            `bson:"load_v"`
	ChargeI       float32            `bson:"charge_I"`
	ChargeStage   string             `bson:"charge_stage"`
	Error         string             `bson:"error"`
	BattType      string             `bson:"batt_type"`
	ChargerStatus string             `bson:"charger_status"`
	temp          float32            `bson:"temp"`
	time          time.Time          `bson:"time"`
	CreatedAt     time.Time          `bson:"created_at"`
	fps           float32            `bson:"fps"`
}

type LogRepository interface {
	Create(c context.Context, log *Log) error
	FindByID(ctx context.Context, id primitive.ObjectID) (Log, error)
	FindByTimeRange(ctx context.Context, startTime, endTime time.Time) ([]Log, error)
}

type LogUsecase interface {
	CreateLog(ctx context.Context, log *Log) error
	GetLogByID(ctx context.Context, id primitive.ObjectID) (*Log, error)
	GetLogsByTimeRange(ctx context.Context, startTime, endTime time.Time) ([]Log, error)
}
