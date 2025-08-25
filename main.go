package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type SlackResponse struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

func coinHandler(w http.ResponseWriter, r *http.Request) {
	result := "Heads"
	if rand.Intn(2) == 1 {
		result = "Tails"
	}

	resp := SlackResponse{
		ResponseType: "in_channel",
		Text:         ":coin: " + result + "!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/coin", coinHandler)
	http.ListenAndServe(":8080", nil)
}
