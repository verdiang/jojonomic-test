package configs

import (
	"context"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	kafka "github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

// NewConfig config for this application
func NewConfig(log logrus.FieldLogger) (*Config, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Info("Error loading .env file")
	}

	route := mux.NewRouter()
	kafka, err := kafka.DialLeader(context.Background(), "tcp", os.Getenv("KAFKA_URL"), os.Getenv("KAFKA_TOPIC"), 0)
	if err != nil {
		log.Fatalf("kafka connection error : %v", err.Error())
	}

	config := &Config{
		Kafka: kafka,
		Route: route,
		Log:   log,
	}

	return config, nil
}

// Config configuration structure for this application
type Config struct {
	Kafka *kafka.Conn
	Route *mux.Router
	Log   logrus.FieldLogger
}
