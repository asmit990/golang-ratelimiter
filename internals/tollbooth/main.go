package main

import (
	"encoding/json"
	"log"
	"net/http"

	tollbooth "github.com/didip/tollbooth/v7"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func endpointHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	msg := Message{
		Status: "success",
		Body:   "Request processed successfully",
	}

	err := json.NewEncoder(writer).Encode(msg)
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {

	msg := Message{
		Status: "error",
		Body:   "You have reached the request limit.",
	}

	jsonMessage, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Failed to marshal rate limit message: %v", err)
	}

	
	limiter := tollbooth.NewLimiter(1, nil)
	limiter.SetMessageContentType("application/json")
	limiter.SetMessage(string(jsonMessage))

	
	http.Handle("/ping", tollbooth.LimitFuncHandler(limiter, endpointHandler))

	log.Println("Server running on http://localhost:9080")
	if err := http.ListenAndServe(":9080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
