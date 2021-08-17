package main

import (
	"context"
	"github.com/ferencovonmatterhorn/twitch-lichess-predictions/pkg/client"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	log.Debug("creating new client")
	c := client.NewClient()

	log.Debug("creating fake context")
	ctx := context.TODO()

	log.Debug("getting game from lichess api")
	game, err := c.GetGame(ctx)
	if err != nil {
		return
		log.Error("Error in main.go")
	}
	log.Debugf("current or last game id is %s", game.ID)
}
