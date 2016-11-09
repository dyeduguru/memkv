package server_test

import (
	"testing"
	"github.com/memkv/server"
	"github.com/stretchr/testify/assert"
	"os"
	"net/http"
	"fmt"
	"encoding/json"
	"bytes"
	"time"
)

var(
	cfg  = &server.ServerConfig{Host: "localhost", Port: 5555}
	kvExpected = &server.KeyValue{Key: "key1", Value:"value1"}
	params =  "application/json; charset=utf-8"
)

func TestServer(t *testing.T) {
	addEndpoint := fmt.Sprintf("http://%v:%v/add-key", cfg.Host, cfg.Port)
	getEndpoint := fmt.Sprintf("http://%v:%v/get-value", cfg.Host, cfg.Port)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(kvExpected)
	resp, err := http.Post(addEndpoint, params, b)
	assert.NoError(t, err," Error while posting keyvalue")
	b = new(bytes.Buffer)
	json.NewEncoder(b).Encode(&server.KeyValue{Key: kvExpected.Key})
	resp, err = http.Post(getEndpoint, params, b)
	assert.NoError(t, err," Error while getting keyvalue")
	var kvActual server.KeyValue
	json.NewDecoder(resp.Body).Decode(&kvActual)
	assert.NoError(t, err," Error while unmarshalling keyvalue")
	assert.True(t, kvActual.Value == kvExpected.Value, "Unexpected value. Actual:%v, Expected:%v",
		kvActual, kvExpected)
}


func TestMain(m *testing.M) {
	server := server.Server(cfg)
	go server.ListenAndServe()
	time.Sleep(time.Minute)
	os.Exit(m.Run())
}