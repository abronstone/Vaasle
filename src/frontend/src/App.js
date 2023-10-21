import React, { useState, useEffect, useCallback } from "react";
import Wordle from "./components/Wordle";
import { newGameApi } from "./components/util/apiCalls";

function App() {
  const [gameState, setGameState] = useState(null);

  const initialGameState = useCallback(async () => {
    try {
      const maxGuesses = 5;
      const wordLength = 5;
      const initialData = await newGameApi(maxGuesses, wordLength);
      setGameState(initialData);
    } catch (error) {
      console.error("Failed to initialize game state:", error);
    }
  }, []);

  // Use useEffect to call initialGameState when the component mounts
  useEffect(() => {
    initialGameState();
  }, [initialGameState]);

  return (
    <div className="App">
      <h1>Wordle (vaas.ai)</h1>
      <Wordle gameState={gameState} setGameState={setGameState} />
    </div>
  );
}

export default App;
