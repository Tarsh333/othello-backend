package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"othello-backend/controllers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("POST /game", controllers.AddGameHandler)            // will have user details and return otp
	mux.HandleFunc("POST /game/{id}/join", controllers.JoinGameHandler) //will have otp and user details and return done/not
	mux.HandleFunc("POST /game/{id}/move", controllers.MoveHandler)     //perform move in game by user
	mux.HandleFunc("GET /game/{id}", controllers.GetGameData)
	mux.HandleFunc("GET /game/{id}/events", controllers.HandleSSE)
	fmt.Println("server running at port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("server not running at port 8080")
	}
}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	resp := map[string]string{
		"status": "ok",
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(resp)
}
