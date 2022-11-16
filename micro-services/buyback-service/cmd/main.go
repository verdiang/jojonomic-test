package main

import (
	"buyback-service/configs"
	"buyback-service/internal/handlers"
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	log := initLogger()
	c, _ := configs.NewConfig(log)
	h := handlers.NewHandler(c)

	c.Route.HandleFunc("/api/buyback", h.Buyback()).Methods(http.MethodPost)

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
