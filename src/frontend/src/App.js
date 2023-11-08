import React, { useState, useEffect, useCallback } from "react";
import CurrentUserGame from "./CurrentUserGame";
import MultiplayerGame from "./MultiplayerGame";
import { newGameApi, getExternalUserGamesApi } from "./components/util/apiCalls";
import { useAuth0 } from "@auth0/auth0-react";
import ErrorBadge from "./components/ErrorBadge";

function App() {
  const { isAuthenticated, user } = useAuth0();
  const [gameState, setGameState] = useState(null);
  const [externalGamesState, setExternalGamesState] = useState(null)
  const [error, setError] = useState(null)
  // TODO replace this with multiple pages 
  const [isMultiplayerEnabled, setIsMultiplayerEnabled] = useState(true);

  const initialGameState = useCallback(async () => {
    try {
      if (!isAuthenticated) return;
      const maxGuesses = 5;
      const wordLength = 5;
      const initialData = await newGameApi(maxGuesses, wordLength, user.sub);
      setGameState(initialData);
    } catch (error) {
      setError("Failed to initialize game state: " + error)
    }
  }, [isAuthenticated]);

  // Use useEffect to call initialGameState when the component mounts
  useEffect(() => {
    initialGameState();
  }, [initialGameState]);

  // Fetch state of externalUserGames for multiplayer when the current user submits a guess
  // (the game state in the CurrentUserGame component changes)
  useEffect(() => {
    if (gameState != null && gameState.metadata.gameID != null && isMultiplayerEnabled) {
      const fetchNewExternalGameState = async () => {
        try {
          const res = await getExternalUserGamesApi(gameState.metadata.gameID);
          setExternalGamesState(res);
        }
        catch (e) {
          setError("Failed to retrieve external user games" + e)
        }
      }

      fetchNewExternalGameState();
    }
  }, [gameState, isMultiplayerEnabled])

  return (
    <div className="App">
      <h1>Vaasle</h1>
      {error != null && <ErrorBadge text={error} />}
      <CurrentUserGame gameState={gameState} setGameState={setGameState} />
      {isMultiplayerEnabled && isAuthenticated && externalGamesState != null && Array.from(externalGamesState.externalUserGamesMap.values()).map((game) =>
        <MultiplayerGame externalUserGameGuesses={game} />
      )}
    </div>
  );
}

export default App;
