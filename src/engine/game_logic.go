package main

// Globally scoped map containing all Wordle game instances.
// Key (string) is a UUID string, and value is the ID's associated game struct.
var games map[string]*game = make(map[string]*game, 0)

// Registers a game into the global games map.
func registerGame(game *game) {
	games[game.Metadata.GameID] = game
}
