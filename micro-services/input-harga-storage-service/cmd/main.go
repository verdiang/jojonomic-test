package main

import (
	"input-harga-storage-service/configs"
	"input-harga-storage-service/internal/handlers"

	"github.com/sirupsen/logrus"
)

func main() {
	log := initLogger()
	c, _ := configs.NewConfig(log)
	h := handlers.NewHandler(log, c.DB())

	h.ReadPriceMessage(log, c)
}

func initLogger() logrus.FieldLogger {
	return logrus.StandardLogger()
}
