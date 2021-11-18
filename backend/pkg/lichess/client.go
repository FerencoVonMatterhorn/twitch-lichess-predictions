package lichess

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	Username   string
	HTTPClient *http.Client
}

func NewClient(username string) *Client {
	return &Client{
		BaseURL:  "https://lichess.org/api",
		Username: username,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (client Client) GetCurrentGameForUser(ctx context.Context) (*GameRes, error) {
	log.Debug("Starting GetCurrentGameForUser function")

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/user/%s/current-game", client.BaseURL, client.Username), nil)
	if err != nil {
		return nil, err
	}

	request.WithContext(ctx)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	result, err := client.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer result.Body.Close()

	if result.StatusCode < http.StatusOK || result.StatusCode >= http.StatusBadRequest {
		log.Errorf("An error occured, the error code is %d", result.StatusCode)
		return nil, err
	}

	var gameRes GameRes
	if err := json.NewDecoder(result.Body).Decode(&gameRes); err != nil {
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
