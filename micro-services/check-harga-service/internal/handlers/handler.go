package handlers

import (
	"check-harga-service/configs"
	"check-harga-service/internal/models"
	"net/http"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type handler struct {
	db  *gorm.DB
	log logrus.FieldLogger
}

// Handlers interface
type Handlers interface {
	CheckPrice(log logrus.FieldLogger, c *configs.Config) func(w http.ResponseWriter, r *http.Request)
}

// NewHandler for get handler
func NewHandler(log logrus.FieldLogger, db *gorm.DB) Handlers {
	return &handler{
		db:  db,
		log: log,
	}
}

func (h *handler) CheckPrice(log logrus.FieldLogger, c *configs.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var harga models.Harga

		if err := h.db.Model(&harga).Select("harga_buyback", "harga_topup").Order("created_at DESC").First(&harga).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				h.log.Error(err.Error())
				response(w, http.StatusNotFound, true, "Kafka not ready")
				return
			}

			h.log.Error(err.Error())
			response(w, http.StatusBadRequest, true, "Kafka not ready")
			return

		}

		response(w, http.StatusOK, false, harga)
	}

}
