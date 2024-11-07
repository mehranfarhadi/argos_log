// internal/usecase/log_usecase.go

package usecase

import (
	"context"
	"time"

	"github.com/mehranfarhadi/argos_log/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// logUsecase is the implementation of the LogUsecase interface
type logUsecase struct {
	logRepo        domain.LogRepository
	contextTimeout time.Duration
}

// NewLogUsecase initializes a new instance of logUsecase
func NewLogUsecase(logRepo domain.LogRepository, timeout time.Duration) domain.LogUsecase {
	return &logUsecase{
		logRepo:        logRepo,
		contextTimeout: timeout,
	}
}

// CreateLog adds a new log entry to the database
func (u *logUsecase) CreateLog(ctx context.Context, log *domain.Log) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.logRepo.Create(ctx, log)
}

// GetLogByID retrieves a log entry by its ObjectID
func (u *logUsecase) GetLogByID(ctx context.Context, id primitive.ObjectID) (*domain.Log, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	log, err := u.logRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &log, nil
}

// GetLogsByTimeRange retrieves logs created within the specified time range
func (u *logUsecase) GetLogsByTimeRange(ctx context.Context, startTime, endTime time.Time) ([]domain.Log, error) {
	logs, err := u.logRepo.FindByTimeRange(ctx, startTime, endTime)
	if err != nil {
		return nil, err
	}
	return logs, nil
}
