import React from 'react';
import Wordle from './components/Wordle';
import { useCallback, useEffect, useState } from 'react';  // Added useEffect and useState
import { newGameApi } from './components/util/apiCalls';

function App() {
  const [gameState, setGameState] = useState(null);  

  const initialGameState = useCallback(async () => {
    const maxGuesses = 5;
    const wordLength = 5;

    const gameState = await newGameApi(maxGuesses, wordLength);
    setGameState(gameState);
  }, []);


  // Added useEffect to call gameState when the component mounts
  useEffect(() => {
    const fetchData = async () => {
      const data = await initialGameState();
      setGameState(data);
    };

    fetchData();
  }, [gameState]);

  return (
    <>
      <div className="App">
        <h1>Wordle (vaas.ai)</h1>
        <Wordle gameState={gameState} setGameState={setGameState} />
      </div>
    </>
  );
}

export default App;
