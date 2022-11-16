package main

import (
	"buyback-storage-service/configs"
	"buyback-storage-service/internal/handlers"

	"github.com/sirupsen/logrus"
)

func main() {
	log := initLogger()
	c, _ := configs.NewConfig(log)
	h := handlers.NewHandler(log, c.DB())

	h.BuybackStorage(log, c)

}

func initLogger() logrus.FieldLogger {
	return logrus.StandardLogger()
}
