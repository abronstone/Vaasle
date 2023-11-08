import React, { useState, useEffect, useCallback } from "react";
import CurrentUserGame from "./CurrentUserGame";
import MultiplayerGame from "./MultiplayerGame";
import { newGameApi } from "./components/util/apiCalls";
import { useAuth0 } from "@auth0/auth0-react";

function App() {
  const { isAuthenticated, user } = useAuth0();
  const [gameState, setGameState] = useState(null);
  // TODO replace this with multiple pages 
  const [isMultiplayerEnabled, setIsMultiplayerEnabled] = useState(true);

  const initialGameState = useCallback(async () => {
    try {
      if(!isAuthenticated) return;
      const maxGuesses = 5;
      const wordLength = 5;
      const initialData = await newGameApi(maxGuesses, wordLength, user.sub);
      setGameState(initialData);
    } catch (error) {
      console.error("Failed to initialize game state:", error);
    }
  }, [isAuthenticated]);

  // Use useEffect to call initialGameState when the component mounts
  useEffect(() => {
    initialGameState();
  }, [initialGameState]);

  return (
    <div className="App">
      <h1>Vassle</h1>
      <CurrentUserGame gameState={gameState} setGameState={setGameState} />
      {isMultiplayerEnabled && isAuthenticated && <MultiplayerGame gameState={gameState} setGameState={setGameState} />}
    </div>
  );
}

export default App;
