package events

import (
	"othello-backend/models"
	"sync"
)

var (
	subscribers = make(map[int][]chan models.Game)
	mu          sync.RWMutex
)

func BroadcastGameChanges(game models.Game) {
	mu.RLock()
	defer mu.RUnlock()
	gameid := game.ID
	for _, sub := range subscribers[gameid] {
		select {
		case sub <- game:
		default:
		}
	}
}

func Subscribe(gameID int) chan models.Game {
	mu.Lock()
	defer mu.Unlock()

	ch := make(chan models.Game, 1)

	subscribers[gameID] = append(
		subscribers[gameID],
		ch,
	)

	return ch
}

func Unsubscribe(gameID int, target chan models.Game) {
	mu.Lock()
	defer mu.Unlock()

	channels := subscribers[gameID]

	for i, ch := range channels {
		if ch == target {

			subscribers[gameID] = append(
				channels[:i],
				channels[i+1:]...,
			)

			close(ch)

			break
		}
	}

	if len(subscribers[gameID]) == 0 {
		delete(subscribers, gameID)
	}
}
