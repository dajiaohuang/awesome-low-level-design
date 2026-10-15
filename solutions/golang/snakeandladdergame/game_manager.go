package snakeandladdergame

import "sync"

type GameManager struct {
	Games []*SnakeAndLadderGame
	mu sync.Mutex
}

var instance *GameManager
var managerOnce sync.Once

func GetGameManager() *GameManager {
	managerOnce.Do(func() {
		instance = &GameManager{Games: []*SnakeAndLadderGame{}}
	})
	return instance
}

func (gm *GameManager) StartNewGame(wg *sync.WaitGroup, playerNames []string) {
	game := NewSnakeAndLadderGame(playerNames)
	gm.mu.Lock()
	gm.Games = append(gm.Games, game)
	gm.mu.Unlock()
	wg.Add(1)
	go game.Play(wg)
}
