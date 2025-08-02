package main

import (
	"encoding/json"
	"fmt"

	"net/http"
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
		return
	}
}

func main() {
	
	http.Handle("/ping", rateLimiter(endpointHandler))

	fmt.Println("Server starting on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}