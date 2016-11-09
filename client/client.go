package client

import (
	"net/http"
	"time"
)

type Client struct {
	HttpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		HttpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (c *Client) AddKey(key, value string) error {

}

func (c *Client) GetKey(key string) (string, error) {

}