package main

import (
	"check-harga-service/configs"
	"check-harga-service/internal/handlers"
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	log := initLogger()
	c, _ := configs.NewConfig(log)
	h := handlers.NewHandler(log, c.DB())

	c.Route.HandleFunc("/api/check-harga", h.CheckPrice(log, c))

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), c.Route))

}

func initLogger() logrus.FieldLogger {
	return logrus.StandardLogger()
}
