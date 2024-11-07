package consumer

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mehranfarhadi/argos_log/bootstrap"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type Consumer struct {
	MqttClient *bootstrap.MqttClient
	Mongo      mongo.Database
	Topic      string
}

// NewConsumer initializes a new Consumer instance
func NewConsumer(mqttClient *bootstrap.MqttClient, mongo mongo.Database, topic string) *Consumer {
	return &Consumer{
		MqttClient: mqttClient,
		Mongo:      mongo,
		Topic:      topic,
	}
}

// StartListening starts the consumer to listen on the specified topic
func (c *Consumer) StartListening() {
	messageHandler := func(client mqtt.Client, msg mqtt.Message) {
		var logData entity.Log
		if err := json.Unmarshal(msg.Payload(), &logData); err != nil {
			log.Printf("Failed to unmarshal message: %v\n", err)
			return
		}

		// Populate any additional fields if necessary, like ID and timestamps
		logData.ID = primitive.NewObjectID()
		logData.CreatedAt = time.Now()

		// Insert the log data into MongoDB
		if err := c.saveToDatabase(logData); err != nil {
			log.Printf("Failed to save log to database: %v\n", err)
		} else {
			log.Printf("Successfully saved log from topic %s\n", msg.Topic())
		}
	}

	// Subscribe to the topic
	token := c.MqttClient.ClientMQ.Subscribe(c.Topic, 1, messageHandler)
	token.Wait()
	if token.Error() != nil {
		log.Fatalf("Subscription error: %v\n", token.Error())
	}
	log.Printf("Subscribed to topic %s\n", c.Topic)
}
