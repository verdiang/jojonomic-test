package handlers

import (
	"context"
	"encoding/json"
	"input-harga-storage-service/configs"
	"input-harga-storage-service/internal/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type handler struct {
	db  *gorm.DB
	log logrus.FieldLogger
}

type Handlers interface {
	ReadPriceMessage(log logrus.FieldLogger, c *configs.Config)
}

// NewHandler for get handler
func NewHandler(log logrus.FieldLogger, db *gorm.DB) Handlers {
	return &handler{
		db:  db,
		log: log,
	}
}

func (h *handler) ReadPriceMessage(log logrus.FieldLogger, c *configs.Config) {
	ctx := context.Background()
	for {
		msg, err := c.Kafka.FetchMessage(ctx)
		if err != nil {
			log.Printf("error : %s", err.Error())
			break
		}

		log.Printf("message at topic/partition/offset %v/%v/%v: %s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key))

		var harga models.Harga
		if err := json.Unmarshal(msg.Value, &harga); err != nil {
			log.Printf("unmarshall data error : %s", err.Error())
		}

		if err := c.DB().Model(&harga).Create(&harga).Error; err != nil {
			log.Error("Fail to update new harga, error : %s", err.Error())
		}

		if err := c.Kafka.CommitMessages(ctx, msg); err != nil {
			log.Printf("Fail to commit message, error : %s", err.Error())
		}

		log.Info("Success to update data")
	}
}
