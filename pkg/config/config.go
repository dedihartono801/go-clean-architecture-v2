package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseUser     string `envconfig:"DATABASE_USER" default:"root"`
	DatabasePassword string `envconfig:"DATABASE_PASSWORD" default:"root"`
	DatabaseHost     string `envconfig:"DATABASE_HOST" default:"mysql"`
	DatabaseName     string `envconfig:"DATABASE_NAME" default:"clean-architecture"`
	DatabasePort     string `envconfig:"DATABASE_PORT" default:"3306"`
	MongoAddress     string `envconfig:"MONGO_ADDRESS" default:"mongodb://localhost:27017"`
	RedisAddress     string `envconfig:"REDIS_ADDRESS" default:"redis:6379"`
	KafkaAdress      string `envconfig:"KAFKA_ADDRESS" default:"kafka:9092"`
	ConsumerGroup    string `envconfig:"CONSUMER_GROUP" default:"my-consumer-group"`
	Topic            string `envconfig:"TOPIC" default:"my-topic"`
}

func SetupEnvFile() *Config {
	envConfig := &Config{}
	_ = godotenv.Load()
	err := envconfig.Process("", envConfig)
	if err != nil {
		log.Fatal(nil, "Fatal error ", err)
	}

	return envConfig
}

// GetEnv
func GetEnv(key string) string {
	value := os.Getenv(key)
	return value
}
