package main

import (
	"fmt"
	"input-harga-service/configs"
	"input-harga-service/internal/handlers"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	log := initLogger()
	c, _ := configs.NewConfig(log)
	h := handlers.NewHandler(log)

	c.Route.HandleFunc("/api/input-harga", h.CreatePrice()).Methods("POST")

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), c.Route))

}

func initLogger() logrus.FieldLogger {
	return logrus.StandardLogger()
}
