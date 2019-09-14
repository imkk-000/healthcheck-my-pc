package main

import (
	"healthcheck/server/api"
	"healthcheck/server/service"
	"log"
	"net/http"
)

func main() {
	healthcheckAPI := &api.HealthcheckAPI{
		HealthcheckServiceImpl: &service.HealthcheckServiceImpl{},
	}
	if err := http.ListenAndServe(":5000", healthcheckAPI); err != nil {
		log.Fatal(err)
	}
}
