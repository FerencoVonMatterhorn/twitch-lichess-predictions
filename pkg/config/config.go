package config

import (
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Credentials LichessCredentials
	Logger      LoggerConfig
}

type LichessCredentials struct {
	Username string
}

type LoggerConfig struct {
	Loglevel string
}

func (c Config) String() string {
	return fmt.Sprintf("Set loglevel to %s and username to %s", c.Logger.Loglevel, c.Credentials.Username)
}

func Parse() (Config, error) {
	loglevel := flag.String("l", log.InfoLevel.String(), "Set the Loglevel")
	username := flag.String("u", "", "Set the Lichess User for authentication")
	flag.Parse()
	c := Config{
		Logger: LoggerConfig{
			Loglevel: *loglevel,
		},
		Credentials: LichessCredentials{
			Username: *username,
		},
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
