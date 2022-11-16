package configs

import (
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

	config := &Config{
		Route: route,
		Log:   log,
		Db:    db,
	}

	return config, nil
}

// Config configuration structure for this application
type Config struct {
	Route *mux.Router
	Log   logrus.FieldLogger
	Db    *gorm.DB
}

// DB call gorm for this application
func (c *Config) DB() *gorm.DB {
	return c.Db
}
