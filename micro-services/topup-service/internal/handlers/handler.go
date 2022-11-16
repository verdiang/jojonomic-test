package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"topup-service/configs"
	"topup-service/internal/models"

	"github.com/segmentio/kafka-go"
	"github.com/teris-io/shortid"
)

type handler struct {
	c *configs.Config
}

// Handlers interface of handler function
type Handlers interface {
	TopUp() func(w http.ResponseWriter, r *http.Request)
}

// NewHandler for get handler
func NewHandler(c *configs.Config) Handlers {
	return &handler{
		c: c,
	}
}

func (h *handler) TopUp() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var mp models.Harga
		var mt models.Topup

		decode := json.NewDecoder(r.Body)
		reffID, _ := shortid.Generate()

		mt.ID = reffID
		if err := decode.Decode(&mt); err != nil {
			h.c.Log.Error(err.Error())
			response(w, http.StatusBadRequest, true, reffID, "Kafka not ready")
			return
		}
		defer r.Body.Close()

		g, err := strconv.ParseFloat(mt.Gram, 64)
		if err != nil {
			h.c.Log.Error(err.Error())
			response(w, http.StatusUnprocessableEntity, true, reffID, "Kafka not ready")
			return
		}

		if g > 0 {
			sg := fmt.Sprintf("%.4f", g)

			if sg[5:] != "0" {
				h.c.Log.Error("Invalid Gram Value")
				response(w, http.StatusUnprocessableEntity, true, reffID, "Kafka not ready")
				return
			}
		}

		if err := h.c.DB().Model(&mp).Select("harga_topup").Order("created_at DESC").First(&mp).Error; err != nil {
			h.c.Log.Errorf("Get harga failed, error : %s", err.Error())
			response(w, http.StatusBadRequest, true, reffID, "Kafka not ready")
			return
		}

		if mt.Harga != int(mp.HargaTopup) {
			h.c.Log.Errorf("Different from current price")
			response(w, http.StatusBadRequest, true, reffID, "Kafka not ready")
			return
		}

		byteData, err := json.Marshal(&mt)
		if err != nil {
			response(w, http.StatusUnprocessableEntity, true, reffID, "Kafka not ready")
			return
		}

		h.c.Kafka.SetWriteDeadline(time.Now().Add(10 * time.Second))
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("address-%s", r.RemoteAddr)),
			Value: byteData,
		}

		_, err = h.c.Kafka.WriteMessages(msg)
		if err != nil {
			h.c.Log.Println(err.Error())
			response(w, http.StatusBadRequest, true, reffID, "Kafka not ready")
		}

		response(w, http.StatusOK, false, reffID, nil)

	}
}
