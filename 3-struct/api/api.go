package api

import "jsoncli/config"

type Client struct {
	APIKey string
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		APIKey: cfg.Key,
	}
}
