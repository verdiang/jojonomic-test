package handlers

import (
	"check-mutasi-service/configs"
	"check-mutasi-service/internal/models"
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
	CheckMutasi(log logrus.FieldLogger, c *configs.Config) func(w http.ResponseWriter, r *http.Request)
}

// NewHandler for get handler
func NewHandler(log logrus.FieldLogger, db *gorm.DB) Handlers {
	return &handler{
		db:  db,
		log: log,
	}
}

func (h *handler) CheckMutasi(log logrus.FieldLogger, c *configs.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var mt []models.Topup
		var mmr models.MutationRequest

		decode := json.NewDecoder(r.Body)
		if err := decode.Decode(&mmr); err != nil {
			h.log.Error(err.Error())
			response(w, http.StatusBadRequest, true, "", "Kafka not ready")
			return
		}
		defer r.Body.Close()

		if err := h.db.Model(&mt).Where("norek", mmr.Norek).Where("created_at BETWEEN ? AND ?", mmr.StartDate, mmr.EndDate).Find(&mt).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				h.log.Error(err.Error())
				response(w, http.StatusNotFound, true, "", "Kafka not ready")
				return
			}

			h.log.Error(err.Error())
			response(w, http.StatusBadRequest, true, "", "Kafka not ready")
			return

		}

		response(w, http.StatusOK, false, "", mt)
	}

}
