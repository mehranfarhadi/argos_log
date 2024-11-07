package bootstrap_test

import (
	"github.com/mehranfarhadi/argos_log/bootstrap"
	"log"
	"testing"
)

func TestConnection(t *testing.T) {
	address := "cert.argos.vision"
	port := 8883
	username := "mohammad"
	password := "mehran"
	clientId := "mohammad"

	_, err := bootstrap.Connect(address, port, username, password, clientId)
	//func Connect(address string, port int, userName string, password string, clientID string, liveTopic string, pubTopic string, pubMessage string)
	if err != nil {
		log.Fatalf("Error connecting to MQTT: %v \n", err)
	}
}

func TestSubscribe(t *testing.T) {
	address := "cert.argos.vision"
	port := 8883
	username := "mohammad"
	password := "mehran"
	clientId := "mohammad"

	m, err := bootstrap.Connect(address, port, username, password, clientId)
	//func Connect(address string, port int, userName string, password string, clientID string, liveTopic string, pubTopic string, pubMessage string)
	if err != nil {
		log.Fatalf("Error connecting to MQTT: %v \n", err)
	}

	m.Sub(bootstrap.Video_topic)
}

func TestPublish(t *testing.T) {
	address := "cert.argos.vision"
	port := 8883
	username := "mohammad"
	password := "mehran"
	clientId := "mohammad"

	m, err := bootstrap.Connect(address, port, username, password, clientId)
	//func Connect(address string, port int, userName string, password string, clientID string, liveTopic string, pubTopic string, pubMessage string)
	if err != nil {
		log.Fatalf("Error connecting to MQTT: %v \n", err)
	}

	m.Publish(bootstrap.Live_topic, "t")
}
