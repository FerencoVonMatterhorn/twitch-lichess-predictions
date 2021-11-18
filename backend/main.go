package main

import (
	"context"
	"github.com/ferencovonmatterhorn/twitch-lichess-predictions/backend/pkg/config"
	"github.com/ferencovonmatterhorn/twitch-lichess-predictions/backend/pkg/lichess"
	"github.com/ferencovonmatterhorn/twitch-lichess-predictions/backend/pkg/twitch"
	log "github.com/sirupsen/logrus"
)

func main() {

	conf, err := config.Parse()
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("%s", conf)

	lichessClient := lichess.NewClient(conf.Credentials.LichessCredentials.Username)
	twitchClient := twitch.NewClient(conf.Credentials.TwtichCredentials.BroadcasterId, conf.Credentials.TwtichCredentials.APIKey, conf.Credentials.TwtichCredentials.ClientId)

	ctx := context.TODO()

	game, err := lichessClient.GetCurrentGameForUser(ctx)
	if err != nil {
		return
		log.Error("Error in lichess Client")
	}

	prediction, err := twitchClient.GetCurrentPredictionForUser(ctx)
	if err != nil {
		return
		log.Error("Error in twitch Client")
	}
	log.Infof("Current or last game id for lichess player %s is %s", conf.Credentials.LichessCredentials.Username, game.ID)
	log.Infof("Status of last prediction is %s", prediction.Data[0].Status)
}
