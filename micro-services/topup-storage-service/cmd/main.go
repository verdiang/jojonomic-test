package main

import (
	"topup-storage-service/configs"
	"topup-storage-service/internal/handlers"

	"github.com/sirupsen/logrus"
)

func main() {
	log := initLogger()
	c, _ := configs.NewConfig(log)
	h := handlers.NewHandler(log, c.DB())

	h.TopUpStorage(log, c)

}

func initLogger() logrus.FieldLogger {
	return logrus.StandardLogger()
}
