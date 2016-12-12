package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, obj interface{}, status int) {
	if err, ok := obj.(error); ok {
		obj = err.Error()
	}
	data, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Failed to write response %v to the caller", string(data))
	}
}
