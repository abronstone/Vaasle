package main

import (
	"sync"
	"vaas/structs"
)

// Globally scoped map containing all Wordle game instances.
// Key (string) is a UUID string, and value is the ID's associated game struct.
var games gamesMap

// Goroutine-safe map for storing games.
type gamesMap struct {
	games sync.Map
}

// Get the game associated with the given ID.
func (g *gamesMap) load(key string) (value *structs.Game, ok bool) {
	val, ok := g.games.Load(key)
	if !ok {
		return nil, false
	}

	game, ok := val.(*structs.Game)
	return game, ok
}

// Get the game associated with the given ID.
func (g *gamesMap) store(key string, value *structs.Game) {
	g.games.Store(key, value)
}

// Registers a game into the global games map.
func registerGame(game *structs.Game) {
	games.store(game.Metadata.GameID, game)
}
