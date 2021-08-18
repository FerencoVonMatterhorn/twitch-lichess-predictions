package main

import (
	"context"
	"github.com/ferencovonmatterhorn/twitch-lichess-predictions/pkg/config"
	"github.com/ferencovonmatterhorn/twitch-lichess-predictions/pkg/lichess"
	log "github.com/sirupsen/logrus"
)

func main() {

	conf, err := config.Parse()
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("%s", conf)

	client := lichess.NewClient(conf.Credentials.Username, conf.Credentials.APIKey)

	ctx := context.TODO()

	game, err := client.GetCurrentGameForUser(ctx)
	if err != nil {
		return
		log.Error("Error in main.go")
	}
	log.Debugf("Current or last game id for lichess player %s is %s", conf.Credentials.Username, game.ID)
}
