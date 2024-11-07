package repository

import (
	"context"
	"time"

	"github.com/mehranfarhadi/argos_log/domain"
	"github.com/mehranfarhadi/argos_log/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type logRepository struct {
	database   mongo.Database
	collection string
}

// NewLogRepository initializes a new log repository with a database and collection name
func NewLogRepository(db mongo.Database, collection string) domain.LogRepository {
	return &logRepository{
		database:   db,
		collection: collection,
	}
}

// Create inserts a new log into the specified MongoDB collection
func (r *logRepository) Create(ctx context.Context, log *domain.Log) error {
	log.CreatedAt = time.Now()                        // Set creation time
	collection := r.database.Collection(r.collection) // Get the collection

	_, err := collection.InsertOne(ctx, log)
	return err
}

// FindByID retrieves a log by its ObjectID from the specified collection
func (r *logRepository) FindByID(ctx context.Context, id primitive.ObjectID) (domain.Log, error) {
	collection := r.database.Collection(r.collection) // Get the collection

	var log domain.Log
	filter := bson.M{"_id": id}

	err := collection.FindOne(ctx, filter).Decode(&log)
	if err != nil {
		return log, err
	}
	return log, nil
}

// FindByTimeRange retrieves logs within a specific time range from the collection
func (r *logRepository) FindByTimeRange(ctx context.Context, startTime, endTime time.Time) ([]domain.Log, error) {
	collection := r.database.Collection(r.collection) // Get the collection

	filter := bson.M{
		"created_at": bson.M{
			"$gte": startTime,
			"$lte": endTime,
		},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var logs []domain.Log
	if err = cursor.All(ctx, &logs); err != nil {
		return nil, err
	}

	return logs, nil
}
