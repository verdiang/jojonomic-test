package configs

import (
	"input-harga-storage-service/internal/models"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	kafka "github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// NewConfig config for this application
func NewConfig(log logrus.FieldLogger) (*Config, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Info("Error loading .env file")
	}

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	route := mux.NewRouter()
	brokers := strings.Split(os.Getenv("KAFKA_URL"), ",")
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  os.Getenv("KAFKA_GROUP_ID"),
		Topic:    os.Getenv("KAFKA_TOPIC"),
		MinBytes: 0,
		MaxBytes: 10e6,
	})

	config := &Config{
		Kafka: kafkaReader,
		Route: route,
		Log:   log,
		Db:    db,
	}

	config.RunMigration()

	return config, nil
}

// Config configuration structure for this application
type Config struct {
	Kafka *kafka.Reader
	Route *mux.Router
	Log   logrus.FieldLogger
	Db    *gorm.DB
}

// DB call gorm for this application
func (c *Config) DB() *gorm.DB {
	return c.Db
}

// RunMigration run automigration
func (c *Config) RunMigration() {
	c.Log.Info("Migrating DBs")
	err := c.Db.AutoMigrate(
		&models.Harga{},
	)

	if err != nil {
		c.Log.Error(err.Error())
	}
}
