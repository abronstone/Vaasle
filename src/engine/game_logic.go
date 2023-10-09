package main

// Globally scoped map containing all Wordle game instances.
var games map[string]*game = make(map[string]*game, 0)

// Registers a game into the global games map.
func registerGame(game *game) {
	games[game.Metadata.GameID] = game
}
