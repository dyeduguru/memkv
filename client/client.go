package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/palantir/stacktrace"
	"net/http"
	"net/url"
	"time"
)

type KeyValue struct {
	Key, Value string
}

type Client struct {
	httpClient *http.Client
	endpoint   url.URL
}

const (
	addEndPoint  = "add-key"
	getEndPoint  = "get-value"
	jsonBodyType = "application/json; charset=utf-8"
)

func New(host string, port int) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
		endpoint: url.URL{
			Scheme: "http",
			Host:   fmt.Sprintf("%s:%d", host, port),
		},
	}
}

func (c *Client) AddKey(key, value string) error {
	endpoint := c.endpoint
	endpoint.Path = addEndPoint
	kv := &KeyValue{Key: key, Value: value}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(kv)
	resp, err := c.httpClient.Post(endpoint.String(), jsonBodyType, b)
	if err != nil {
		return stacktrace.Propagate(err, "")
	}
	if resp.StatusCode != http.StatusOK {
		return stacktrace.NewError("Unexpected status code returned: %v\n", resp.StatusCode)
	}
	return nil
}

func (c *Client) GetKey(key string) (string, error) {
	endpoint := c.endpoint
	endpoint.Path = getEndPoint
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(&KeyValue{Key: key})
	resp, err := http.Post(endpoint.String(), jsonBodyType, b)
	if err != nil {
		return "", stacktrace.Propagate(err, "")
	}
	if resp.StatusCode != http.StatusOK {
		return "", stacktrace.NewError("failed to get key: %v\n", resp.Status)
	}
	var kv KeyValue
	if err := json.NewDecoder(resp.Body).Decode(&kv); err != nil {
		return "", stacktrace.Propagate(err, "%v", resp.StatusCode)
	}
	return kv.Value, nil
}
