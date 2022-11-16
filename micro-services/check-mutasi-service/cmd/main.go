package main

import (
	"check-mutasi-service/configs"
	"check-mutasi-service/internal/handlers"
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	log := initLogger()
	c, _ := configs.NewConfig(log)
	h := handlers.NewHandler(log, c.DB())

	c.Route.HandleFunc("/api/check-mutasi", h.CheckMutasi(log, c)).Methods("POST")

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), c.Route))

}

func initLogger() logrus.FieldLogger {
	return logrus.StandardLogger()
}
