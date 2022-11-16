package handlers

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"topup-storage-service/configs"
	"topup-storage-service/internal/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type handler struct {
	db  *gorm.DB
	log logrus.FieldLogger
}

type Handlers interface {
	TopUpStorage(log logrus.FieldLogger, c *configs.Config)
}

// NewHandler for get handler
func NewHandler(log logrus.FieldLogger, db *gorm.DB) Handlers {
	return &handler{
		db:  db,
		log: log,
	}
}

func (h *handler) TopUpStorage(log logrus.FieldLogger, c *configs.Config) {
	ctx := context.Background()
	for {
		msg, err := c.Kafka.FetchMessage(ctx)
		if err != nil {
			log.Printf("error : %s", err.Error())
			break
		}

		log.Printf("message at topic/partition/offset %v/%v/%v: %s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key))

		var mt models.Topup
		var mr models.Rekening
		if err := json.Unmarshal(msg.Value, &mt); err != nil {
			log.Printf("unmarshall data error : %s", err.Error())
		}

		mt.Type = os.Getenv("KAFKA_TOPIC")
		if err := c.DB().Model(&mt).Create(&mt).Error; err != nil {
			log.Error("Fail to update new trasanction, error : %s", err.Error())
		}

		mr.Norek = mt.Norek
		if err := c.DB().Model(&mr).Where("norek = ?", mr.Norek).First(&mr).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				mr.ID = mt.ID
				c.DB().Model(&mr).Create(&mr)
			}

			log.Error("Fail to create new rekening, error : %s", err.Error())
		}

		saldo, _ := strconv.ParseFloat(mt.Gram, 64)
		saldo = mr.Saldo + saldo
		if err := c.DB().Model(&mr).Where("norek", mt.Norek).Update("saldo", saldo).Error; err != nil {
			log.Error("Fail to update saldo for norek : %s", mr.Norek)
		}

		if err := c.Kafka.CommitMessages(ctx, msg); err != nil {
			log.Printf("Fail to commit message, error : %s", err.Error())
		}

		log.Info("Success to update data")
	}
}
