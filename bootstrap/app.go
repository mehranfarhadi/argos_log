package bootstrap

import "github.com/mehranfarhadi/argos_log/mongo"

type Application struct {
	Env   *Env
	Mongo mongo.Client
	Mqtt  *MqttClient
}

func App() Application {

	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	mq, err := Connect(app.Env.MQAddress, app.Env.MQPort, app.Env.MQUserName, app.Env.DBPass, app.Env.MQClientID)
	if err != nil {
		panic(err)
	}
	app.Mqtt = mq
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
