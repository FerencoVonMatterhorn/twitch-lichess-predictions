package twitch

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Client struct {
	BaseURL       string
	BroadcasterID int
	apiKey        string
	clientId      string
	HTTPClient    *http.Client
}

func NewClient(broadcasterID int, apiKey string, clientId string) *Client {
	return &Client{
		BaseURL:       "https://api.twitch.tv/helix",
		BroadcasterID: broadcasterID,
		apiKey:        apiKey,
		clientId:      clientId,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (client Client) GetCurrentPredictionForUser(ctx context.Context) (*PredictionRes, error) {
	log.Debug("Starting GetCurrentPredictionForUser function")

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/predictions?broadcaster_id=%d", client.BaseURL, client.BroadcasterID), nil)
	if err != nil {
		return nil, err
	}

	request.WithContext(ctx)
	request.Header.Set("Authorization: ", fmt.Sprintf("Bearer %s", client.apiKey))
	request.Header.Set("Client-Id: ", fmt.Sprintf("%s", client.clientId))

	result, err := client.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer result.Body.Close()

	if result.StatusCode < http.StatusOK || result.StatusCode >= http.StatusBadRequest {
		log.Errorf("An error occured, the error code is %d", result.StatusCode)
		return nil, err
	}

	var predictionRes PredictionRes
	if err := json.NewDecoder(result.Body).Decode(&predictionRes); err != nil {
		return nil, err
	}

	return &predictionRes, nil
}

type PredictionRes struct {
	Data       []Data `json:"data"`
	Pagination struct {
	} `json:"pagination"`
}

type Data struct {
	ID               string      `json:"id"`
	BroadcasterID    string      `json:"broadcaster_id"`
	BroadcasterName  string      `json:"broadcaster_name"`
	BroadcasterLogin string      `json:"broadcaster_login"`
	Title            string      `json:"title"`
	WinningOutcomeID interface{} `json:"winning_outcome_id"`
	Outcomes         []Outcomes  `json:"outcomes"`
	PredictionWindow int         `json:"prediction_window"`
	Status           string      `json:"status"`
	CreatedAt        time.Time   `json:"created_at"`
	EndedAt          interface {
	} `json:"ended_at"`
	LockedAt interface {
	} `json:"locked_at"`
}

type Outcomes struct {
	ID            string      `json:"id"`
	Title         string      `json:"title"`
	Users         int         `json:"users"`
	ChannelPoints int         `json:"channel_points"`
	TopPredictors interface{} `json:"top_predictors"`
	Color         string      `json:"color"`
}
