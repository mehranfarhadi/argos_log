package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mehranfarhadi/argos_log/bootstrap"
	"log"
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
}

func (m MqttClient) Publish(topic string, message string) {
	token := m.ClientMQ.Publish(topic, 1, false, message)
	token.Wait()
	if token.Error() != nil {
		fmt.Printf("Error publishing message: %v\n", token.Error())
	}
	time.Sleep(time.Second)
}

func Connect(env *bootstrap.Env) MqttClient {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", env.MQAddress, env.MQPort))

	tlsConfig, err := NewTlsConfig()
	if err != nil {
		log.Fatal(err)
	}
	opts.SetTLSConfig(tlsConfig)

	opts.SetClientID(env.MQClientID)
	opts.SetUsername(env.MQUsername)
	opts.SetPassword(env.MQPassword)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// Optionally, keep the connection open for more interactions
	// Defer disconnect until you're sure interactions are finished
	return MqttClient{ClientMQ: client}
}
