package bootstrap

import (
	"github.com/mehranfarhadi/argos_log/mongo"
	"github.com/mehranfarhadi/argos_log/mqtt"
)

type Application struct {
	Env   *Env
	Mongo mongo.Client
	Mqtt  mqtt.MqttClient
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	app.Mqtt = mqtt.Connect(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
