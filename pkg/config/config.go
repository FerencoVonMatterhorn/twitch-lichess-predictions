package config

import (
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Credentials Credentials
	Logger      LoggerConfig
}

type Credentials struct {
	LichessCredentials Lichess
	TwtichCredentials  Twitch
}

type Twitch struct {
	APIKey        string
	ClientId      string
	BroadcasterId int
}

type Lichess struct {
	Username string
}

type LoggerConfig struct {
	Loglevel string
}

func (c Config) String() string {
	return fmt.Sprintf("Set loglevel to %s and username to %s", c.Logger.Loglevel, c.Credentials.LichessCredentials.Username)
}

func Parse() (Config, error) {
	loglevel := flag.String("l", log.InfoLevel.String(), "Set the Loglevel")
	username := flag.String("u", "", "Set the Lichess User")
	apiKey := flag.String("a", "", "Set the API Key for the Twitch API")
	clientId := flag.String("c", "", "Set the Client-ID for the Twitch API")
	broadcasterId := flag.Int("b", 0, "Set the broadcaster id for Twitch API")
	flag.Parse()
	c := Config{
		Logger: LoggerConfig{
			Loglevel: *loglevel,
		},
		Credentials: struct {
			LichessCredentials Lichess
			TwtichCredentials  Twitch
		}{LichessCredentials: struct{ Username string }{Username: *username}, TwtichCredentials: struct {
			APIKey        string
			ClientId      string
			BroadcasterId int
		}{APIKey: *apiKey, ClientId: *clientId, BroadcasterId: *broadcasterId}},
	}
	return c, setLogLevel(*loglevel)
}

func setLogLevel(loglevel string) error {
	lvl, err := log.ParseLevel(loglevel)
	if err != nil {
		return err
	}
	log.SetLevel(lvl)
	return nil
}
