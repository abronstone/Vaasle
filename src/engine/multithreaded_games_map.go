package main

import (
	"sync"
	"vaas/structs"
)

// Globally scoped map containing all Wordle game instances.
// Key (string) is a UUID string, and value is the ID's associated game struct.
var games gamesMap = gamesMap{
	mu:            sync.Mutex{},
	games:         make(map[string]*structs.Game),
	circularCache: [10]string{},
	newest:        0,
}

// Goroutine-safe map for storing games.
type gamesMap struct {
	mu            sync.Mutex
	games         map[string]*structs.Game
	circularCache [10]string
	newest        int
}

// Get the game associated with the given ID.
func (g *gamesMap) load(key string) (value *structs.Game, ok bool) {
	g.mu.Lock()
	defer g.mu.Unlock()

	val, ok := g.games[key]
	return val, ok
}

// Store a game into the games map.
func (g *gamesMap) store(key string, value *structs.Game) {
	g.mu.Lock()
	defer g.mu.Unlock()

	delete(g.games, g.circularCache[g.newest])
	g.circularCache[g.newest] = key
	g.newest = (g.newest + 1) % len(g.circularCache)

	g.games[key] = value
}

// Registers a game into the global games map.
func registerGame(game *structs.Game) {
	games.store(game.Metadata.GameID, game)
}
