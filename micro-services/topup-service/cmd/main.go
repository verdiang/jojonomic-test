package main

import (
	"fmt"
	"net/http"
	"os"
	"topup-service/configs"
	"topup-service/internal/handlers"

	"github.com/sirupsen/logrus"
)

func main() {
	log := initLogger()
	c, _ := configs.NewConfig(log)
	h := handlers.NewHandler(c)

	c.Route.HandleFunc("/api/topup", h.TopUp()).Methods(http.MethodPost)

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), c.Route))
}

func initLogger() logrus.FieldLogger {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)
	return logrus.StandardLogger()
}
