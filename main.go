package main

import (
	"context"
	"net/http"

	"github.com/bawilliams/vandelay-api/router"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Info("starting")

	log.Info("init all endpoints")
	mux, err := router.InitAllEndPoints(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Warn("failed to set up router")
		return
	}

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Warn("failed to start server")
	}

	log.Info("done processing")
}
