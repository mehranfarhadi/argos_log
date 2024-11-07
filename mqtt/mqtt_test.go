package mqtt_test

import (
	"github.com/mehranfarhadi/argos_log/bootstrap"
	"github.com/mehranfarhadi/argos_log/mqtt"
	"testing"
)

func TestConnection(t *testing.T) {
	env := &bootstrap.Env{
		MQClientID: "mohammad",
		MQPort:     8883,
		MQAddress:  "cert.argos.vision",
		MQPassword: "mehran",
		MQUsername: "mohammad",
	}

	_ = mqtt.Connect(env)

}

func TestPublish(t *testing.T) {
	env := &bootstrap.Env{
		MQClientID: "mohammad",
		MQPort:     8883,
		MQAddress:  "cert.argos.vision",
		MQPassword: "mehran",
		MQUsername: "mohammad",
	}

	s := mqtt.Connect(env)
	s.Publish(mqtt.Camera_log_topic, "mamad")
}

func TestSubscribe(t *testing.T) {
	env := &bootstrap.Env{
		MQClientID: "mohammad",
		MQPort:     8883,
		MQAddress:  "cert.argos.vision",
		MQPassword: "mehran",
		MQUsername: "mohammad",
	}

	s := mqtt.Connect(env)
	s.Sub(mqtt.Camera_log_topic)
}
