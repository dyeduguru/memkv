package client_test

import (
	"github.com/dyeduguru/memkv/client"
	"github.com/dyeduguru/memkv/server"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var (
	cfg        = &server.ServerConfig{Host: "localhost", Port: 5555}
	kvExpected = &client.KeyValue{Key: "key1", Value: "value1"}
)

func TestClient(t *testing.T) {
	client := client.New(cfg.Host, cfg.Port)
	err := client.AddKey(kvExpected.Key, kvExpected.Value)
	assert.NoError(t, err)
	v, err := client.GetKey(kvExpected.Key)
	assert.NoError(t, err, "")
	assert.True(t, v == kvExpected.Value)
}

func TestMain(m *testing.M) {
	server := server.Server(cfg)
	go server.ListenAndServe()
	time.Sleep(5 * time.Second)
	os.Exit(m.Run())
}
