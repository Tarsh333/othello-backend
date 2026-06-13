package controllers

import (
	"encoding/json"
	"net/http"
	"othello-backend/data"
	"othello-backend/events"
	"othello-backend/helpers"
	"othello-backend/models"
)

func AddGameHandler(w http.ResponseWriter, r *http.Request) {
	// an 4 digit otp to be generated here
	game := models.Game{}
	var req models.AddGameReq
	otp := helpers.GenerateGameCode()
	game.ID = otp
	json.NewDecoder(r.Body).Decode(&req)
	game.P1 = req.P1
	game.Bs = *models.GetNewBoard()
	err := data.AddGameData(game)
	if err != nil {
		helpers.WriteJSON(w, http.StatusConflict, "")

	} else {
		helpers.WriteJSON(w, http.StatusCreated, game)
	}
}
func JoinGameHandler(w http.ResponseWriter, r *http.Request) {
	var req models.JoinGameReq
	json.NewDecoder(r.Body).Decode(&req)
	otp := helpers.FetchGameCode(r)
	gd, err := data.JoinGame(otp, req.P2)
	if err != nil {
		w.WriteHeader(403)
		json.NewEncoder(w).Encode("")
	} else {
		events.BroadcastGameChanges(gd)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("")
	}
}
func MoveHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req models.MoveReq
	json.NewDecoder(r.Body).Decode(&req)
	otp := helpers.FetchGameCode(r)
	gd, err := data.UpdateGameData(otp, req.Move, req.PlayerTurn)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("")
	} else {
		w.WriteHeader(http.StatusOK)
		events.BroadcastGameChanges(gd)
		json.NewEncoder(w).Encode("")
	}
}
func GetGameData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	otp := helpers.FetchGameCode(r)
	gameData, err := data.GetGameData(otp)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("")
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(gameData)
	}
}
