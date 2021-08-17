package client

import (
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	apiKey     string
	Username   string
	HTTPClient *http.Client
}

func NewClient() *Client {
	return &Client{
		BaseURL:  "https://lichess.org/api",
		apiKey:   "test",
		Username: "Ferenco",
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
