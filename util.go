package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeAsJSON(w http.ResponseWriter, v any) {
	data, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}
