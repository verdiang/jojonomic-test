package handlers

import (
	"check-saldo-service/configs"
	"check-saldo-service/internal/models"
	"encoding/json"
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
	CheckSaldo(log logrus.FieldLogger, c *configs.Config) func(w http.ResponseWriter, r *http.Request)
}

// NewHandler for get handler
func NewHandler(log logrus.FieldLogger, db *gorm.DB) Handlers {
	return &handler{
		db:  db,
		log: log,
	}
}

func (h *handler) CheckSaldo(log logrus.FieldLogger, c *configs.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var mr models.Rekening

		decode := json.NewDecoder(r.Body)
		if err := decode.Decode(&mr); err != nil {
			h.log.Error(err.Error())
			response(w, http.StatusBadRequest, true, "", "Kafka not ready")
			return
		}
		defer r.Body.Close()

		if err := h.db.Model(&mr).Where("norek = ?", mr.Norek).First(&mr).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				h.log.Error(err.Error())
				response(w, http.StatusNotFound, true, "", "Kafka not ready")
				return
			}

			h.log.Error(err.Error())
			response(w, http.StatusBadRequest, true, "", "Kafka not ready")
			return

		}

		response(w, http.StatusOK, false, "", mr)
	}

}
