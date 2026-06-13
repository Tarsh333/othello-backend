package data

import (
	"fmt"
	"othello-backend/models"
	"sync"
)

var (
	gameData = make(map[int]models.Game)
	mu       sync.RWMutex
)

func AddGameData(g models.Game) error {
	mu.Lock()
	defer mu.Unlock()
	_, isPresent := gameData[g.ID]
	if !isPresent {
		gameData[g.ID] = g
		return nil
	} else {
		return fmt.Errorf("game already exists")
	}
}
func JoinGame(gameID int, p2 *models.Player) (models.Game, error) {
	mu.Lock()
	defer mu.Unlock()
	gd, isPresent := gameData[gameID]
	if isPresent {
		if gd.P2 != nil {
			return models.Game{}, fmt.Errorf("player already exist")
		} else {
			gd.P2 = p2
			gameData[gameID] = gd
		}
		return gd, nil
	} else {
		return models.Game{}, fmt.Errorf("game does not exists")
	}
}
func UpdateGameData(gameID int, moves []models.Position, turn models.Cell) (models.Game, error) {
	mu.Lock()
	defer mu.Unlock()
	gd, isPresent := gameData[gameID]
	if isPresent {
		for _, move := range moves {
			gd.Bs.Board[move.Row][move.Col] = turn
			gameData[gameID] = gd
		}
		return gd, nil
	} else {
		return models.Game{}, fmt.Errorf("game does not exists")
	}
}

func GetGameData(gameID int) (models.Game, error) {
	mu.RLock()
	defer mu.RUnlock()
	gd, isPresent := gameData[gameID]
	if isPresent {
		return gd, nil
	} else {
		return models.Game{}, fmt.Errorf("game does not exists")
	}
}
