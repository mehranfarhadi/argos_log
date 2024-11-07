package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mehranfarhadi/argos_log/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogController struct {
	LogUsecase domain.LogUsecase
}

// NewLogController initializes a new LogController
func NewLogController(logUsecase domain.LogUsecase) *LogController {
	return &LogController{LogUsecase: logUsecase}
}

// CreateLog handles the HTTP request to create a new log entry
func (lc *LogController) CreateLog(c *gin.Context) {
	var log domain.Log

	// Bind the JSON payload to the Log struct
	if err := c.ShouldBindJSON(&log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	// Call the usecase to create the log
	if err := lc.LogUsecase.CreateLog(c.Request.Context(), &log); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create log", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Log created successfully"})
}

// GetLogByID handles the HTTP request to fetch a log by its ID
func (lc *LogController) GetLogByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	// Fetch the log by ID
	log, err := lc.LogUsecase.GetLogByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve log", "error": err.Error()})
		return
	}
	if log == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Log not found"})
		return
	}

	c.JSON(http.StatusOK, log)
}

// GetLogsByTimeRange handles the HTTP request to fetch logs by a time range
func (lc *LogController) GetLogsByTimeRange(c *gin.Context) {
	startTimeParam := c.Query("start_time")
	endTimeParam := c.Query("end_time")

	startTime, err := time.Parse(time.RFC3339, startTimeParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid start time format", "error": err.Error()})
		return
	}

	endTime, err := time.Parse(time.RFC3339, endTimeParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid end time format", "error": err.Error()})
		return
	}

	// Fetch logs within the time range
	logs, err := lc.LogUsecase.GetLogsByTimeRange(c.Request.Context(), startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve logs", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}
