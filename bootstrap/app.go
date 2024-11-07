package bootstrap

import (
	"github.com/mehranfarhadi/argos_log/mongo"
	"github.com/mehranfarhadi/argos_log/mqtt"
)

type Application struct {
	Env   *Env
	Mongo mongo.Client
<<<<<<< HEAD
	Mqtt  *MqttClient
=======
	Mqtt  mqtt.MqttClient
>>>>>>> 5286ff29e4269c720c2731364c8721b8fd98394e
}

func App() Application {

	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
<<<<<<< HEAD
	mq, err := Connect(app.Env.MQAddress, app.Env.MQPort, app.Env.MQUserName, app.Env.DBPass, app.Env.MQClientID)
	if err != nil {
		panic(err)
	}
	app.Mqtt = mq
=======
	app.Mqtt = mqtt.Connect(app.Env)
>>>>>>> 5286ff29e4269c720c2731364c8721b8fd98394e
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
