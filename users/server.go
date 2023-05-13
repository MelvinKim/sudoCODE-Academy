package main

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/MelvinKim/users/presentation"
)

const PORT = 9000

func main() {
	ctx := context.Background()

	srv := presentation.PrepareServer(ctx, PORT)

	if err := srv.ListenAndServe(); err != nil {
		log.Errorf("server start up error: %v", err)
		return
	}

	log.Infof("server up and running on port %d", PORT)
}
