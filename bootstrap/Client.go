package bootstrap

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

type MqttClient struct {
	ClientMQ mqtt.Client
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s \n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected to server")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v\n", err)
}

func (m MqttClient) Sub(topic string) {

	token := m.ClientMQ.Subscribe(topic, 1, messagePubHandler)
	token.Wait()
	if token.Error() != nil {
		fmt.Printf("Error subscribing to topic: %v\n", token.Error())
	}
	fmt.Println("Error", token)
}

func (m MqttClient) Publish(topic string, message string) {
	token := m.ClientMQ.Publish(topic, 1, false, message)
	token.Wait()
	if token.Error() != nil {
		fmt.Printf("Error publishing message: %v\n", token.Error())
	}
	time.Sleep(time.Second)
}

func Connect(address string, port int, userName string, password string, clientID string) (*MqttClient, error) {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", address, port))

	tlsConfig, err := NewTlsConfig()
	if err != nil {
		return nil, err
	}
	opts.SetTLSConfig(tlsConfig)

	opts.SetClientID(clientID)
	opts.SetUsername(userName)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	// Optionally, keep the connection open for more interactions
	// Defer disconnect until you're sure interactions are finished
	return &MqttClient{ClientMQ: client}, nil
}
