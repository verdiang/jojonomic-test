package configs

import (
	"context"
	"os"
	"topup-service/internal/models"

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
	kafka, err := kafka.DialLeader(context.Background(), "tcp", os.Getenv("KAFKA_URL"), os.Getenv("KAFKA_TOPIC"), 0)
	if err != nil {
		log.Fatalf("kafka connection error : %v", err.Error())
	}

	config := &Config{
		Kafka: kafka,
		Route: route,
		Log:   log,
		Db:    db,
	}

	config.RunMigration()

	return config, nil
}

// Config configuration structure for this application
type Config struct {
	Kafka *kafka.Conn
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
		&models.Rekening{},
		&models.Topup{},
	)

	if err != nil {
		c.Log.Error(err.Error())
	}
}
