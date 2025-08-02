package main

import (
    "encoding/json"
    "golang.org/x/time/rate"
    "net/http"
)


func rateLimiter(next http.HandlerFunc) http.HandlerFunc {
    limiter := rate.NewLimiter(1, 3) 

    return func(w http.ResponseWriter, r *http.Request) {
        if !limiter.Allow() {
           
            msg := Message{
                Status: "error",
                Body:   "Too Many Requests",
            }

            w.WriteHeader(http.StatusTooManyRequests)
           
            json.NewEncoder(w).Encode(msg)
            return 
        }
        
        
        next(w, r)
    }
}