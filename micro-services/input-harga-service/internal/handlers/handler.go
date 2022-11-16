package handlers

import (
	"encoding/json"
	"fmt"
	"input-harga-service/configs"
	"input-harga-service/internal/models"
	"net/http"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"github.com/teris-io/shortid"
)

type handler struct {
	log logrus.FieldLogger
}

// Handlers interface of handler function
type Handlers interface {
	CreatePrice() func(w http.ResponseWriter, r *http.Request)
}

// NewHandler for get handler
func NewHandler(log logrus.FieldLogger) Handlers {
	return &handler{
		log: log,
	}
}

func (h *handler) CreatePrice() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c, _ := configs.NewConfig(h.log)

		p := models.Price{}
		decode := json.NewDecoder(r.Body)
		reffID, _ := shortid.Generate()
		p.ID = reffID
		if err := decode.Decode(&p); err != nil {
			response(w, http.StatusBadRequest, true, reffID, "Kafka not ready")
			return
		}
		defer r.Body.Close()

		byteParams, err := json.Marshal(&p)
		if err != nil {
			response(w, http.StatusBadRequest, true, reffID, err.Error())
			return
		}

		c.Kafka.SetWriteDeadline(time.Now().Add(10 * time.Second))
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("address-%s", r.RemoteAddr)),
			Value: byteParams,
		}

		_, err = c.Kafka.WriteMessages(msg)
		if err != nil {
			h.log.Println(err.Error())
			response(w, http.StatusBadRequest, true, reffID, "Kafka not ready")
		}

		response(w, http.StatusOK, false, reffID, nil)
	}
}
