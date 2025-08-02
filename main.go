package main

import (
    "log"
    "net/http"

  // use your own or import the correct package paths
    // "github.com/asmitpandey/golangPRoject/internals/per-client-rate-limiting"
	// "github.com/asmitpandey/golangPRoject/internals/token-bucket"
	// "github.com/asmitpandey/golangPRoject/internals/tollbooth"	

  
)

func main() {
    log.Println("Starting server...")

    http.Handle("/per-client/ping", perclient.NewHandler())
    http.Handle("/token-bucket/ping", tokenbucket.NewHandler())
    http.Handle("/tollbooth/ping", tollbooth.NewHandler())

    log.Fatal(http.ListenAndServe(":8080", nil))
}