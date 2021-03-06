package server

import (
	"encoding/json"
	"errors"
	"github.com/palantir/stacktrace"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
)

const maxRequestSize = 1024 * 1024 // 1MB

type Store struct {
	MemStore map[string]string
	Mutex    *sync.Mutex
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewStore() *Store {
	return &Store{
		MemStore: map[string]string{},
		Mutex:    &sync.Mutex{},
	}
}

func (st *Store) AddKey(w http.ResponseWriter, r *http.Request) {
	body, err := readBody(r)
	if err != nil {
		WriteJSON(w, err, http.StatusBadRequest)
		return
	}
	var kv KeyValue
	err = json.Unmarshal(body, &kv)
	if err != nil {
		WriteJSON(w, err, http.StatusBadRequest)
		return
	}
	st.Mutex.Lock()
	st.MemStore[kv.Key] = kv.Value
	st.Mutex.Unlock()
	w.WriteHeader(http.StatusOK)
}

func (st *Store) GetValue(w http.ResponseWriter, r *http.Request) {
	body, err := readBody(r)
	if err != nil {
		WriteJSON(w, err, http.StatusBadRequest)
	}
	var kv KeyValue
	err = json.Unmarshal(body, &kv)
	if err != nil {
		WriteJSON(w, err, http.StatusInternalServerError)
	}
	if v, ok := st.MemStore[kv.Key]; ok {
		WriteJSON(w, KeyValue{Key: kv.Key, Value: v}, http.StatusOK)
	} else {
		WriteJSON(w, errors.New("Unable to find the key"), http.StatusNotFound)
	}
}

func readBody(r *http.Request) ([]byte, error) {
	if r == nil || r.Body == nil {
		return nil, stacktrace.NewError("Request body empty")
	}
	data, err := ioutil.ReadAll(io.LimitReader(r.Body, maxRequestSize))
	if err != nil {
		return nil, stacktrace.Propagate(err, "cannot read input")
	}
	return data, nil
}
