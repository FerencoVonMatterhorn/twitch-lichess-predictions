package client

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (c Client) GetGame(ctx context.Context) (*GameRes, error) {
	log.Debug("started api call function")

	log.Debug("creating request")
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user/%s/current-game", c.BaseURL, c.Username), nil)
	if err != nil {
		return nil, err
	}

	log.Debug("setting context and headers for the request")
	req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	log.Debug("executing request with http client")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	log.Debugf("request status code is %d", res.StatusCode)

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("an error occured, the error code is %d", res.StatusCode)
		log.Errorf("an error occured, the error code is %d", res.StatusCode)
	}

	var gameRes GameRes
	if err := json.NewDecoder(res.Body).Decode(&gameRes); err != nil {
		return nil, err
	}

	return &gameRes, nil
}

type GameRes struct {
	ID         string  `json:"id"`
	Rated      bool    `json:"rated"`
	Variant    string  `json:"variant"`
	Speed      string  `json:"speed"`
	Perf       string  `json:"perf"`
	CreatedAt  int64   `json:"createdAt"`
	LastMoveAt int64   `json:"lastMoveAt"`
	Status     string  `json:"status"`
	Players    Players `json:"players"`
	Opening    Opening `json:"opening"`
	Moves      string  `json:"moves"`
	Clock      Clock   `json:"clock"`
}

type Players struct {
	White White `json:"white"`
	Black Black `json:"black"`
}

type White struct {
	User       `json:"user"`
	Rating     int `json:"rating"`
	RatingDiff int `json:"ratingDiff"`
}

type Black struct {
	User       User `json:"user"`
	Rating     int  `json:"rating"`
	RatingDiff int  `json:"ratingDiff"`
}

type User struct {
	Name   string `json:"name"`
	Title  string `json:"title"`
	Patron bool   `json:"patron"`
	ID     string `json:"id"`
}

type Opening struct {
	Eco  string `json:"eco"`
	Name string `json:"name"`
	Ply  int    `json:"ply"`
}

type Clock struct {
	Initial   int `json:"initial"`
	Increment int `json:"increment"`
	TotalTime int `json:"totalTime"`
}
