import React from 'react';
import Wordle from './components/Wordle';
import { useCallback, useEffect, useState } from 'react';  // Added useEffect and useState
import { newGameApi } from './components/util/apiCalls';

function App() {
  const [gameData, setGameData] = useState(null);  

  const gameState = useCallback(async () => {
    const maxGuesses = 5;
    const wordLength = 5;

    const gameState = await newGameApi(maxGuesses, wordLength);
    return gameState;
  }, []);

  // Added useEffect to call gameState when the component mounts
  useEffect(() => {
    console.log("useEffect called")
    const fetchData = async () => {
      const data = await gameState();
      setGameData(data);
    };

    fetchData();
    console.log("gameData: ", gameData)
  }, [gameState]);

  return (
    <>
      <div className="App">
        <h1>Wordle (vaas.ai)</h1>
        <Wordle gameState={gameData} />
      </div>
    </>
  );
}

export default App;
