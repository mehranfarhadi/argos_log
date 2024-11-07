package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASS"`
	DBName                 string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	MQAddress              string `mapstructure:"MQTT_ADDRESS"`
	MQPort                 int    `mapstructure:"MQTT_PORT"`
<<<<<<< HEAD
	MQUserName             string `mapstructure:"MQTT_USERNAME"`
	MQPassword             string `mapstructure:"MQTT_PASSWORD"`
	MQClientID             string `mqpsructure:"MQTT_CLIENT_ID"`
=======
	MQUsername             string `mapstructure:"MQTT_USERNAME"`
	MQPassword             string `mapstructure:"MQTT_PASSWORD"`
	MQClientID             string `mapstructure:"MQ_CLIENT_ID"`
>>>>>>> 5286ff29e4269c720c2731364c8721b8fd98394e
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
