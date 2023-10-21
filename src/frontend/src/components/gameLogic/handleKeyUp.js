import { makeGuessApi } from "../util/apiCalls";  

const handleKeyup = async (key, turn, isCorrect, currentGuess, guesses, setCurrentGuess, gameId, setGameState) => {
  if (key === 'Enter') {
    // only add guess if turn is less than 5
    if (turn > 5 && !isCorrect) {
      console.log('you used all your guesses!');
      return;
    }
    // do not allow duplicate words
    // if (guesses.includes(currentGuess)) {
    //   console.log('you already tried that word.');
    //   return;
    // }
    console.log("Here is the currentGuess in handleKeyup", currentGuess)
    // check word is 5 chars
    if (currentGuess.length !== 5) {
      console.log('word must be 5 chars.');
      return;
    }
    try {
      const newGameState = await makeGuessApi(gameId, currentGuess);
      console.log('newGameState', newGameState);

      setGameState(newGameState);
      setCurrentGuess('');

    } catch (error) {
      console.error('Failed to update game state:', error);
    }  }  // This is where the problematic curly brace was located
  if (key === 'Backspace') {
    setCurrentGuess(prev => prev.slice(0, -1));
    return;
  }
  if (/^[A-Za-z]$/.test(key)) {
    if (currentGuess.length < 5) {
      setCurrentGuess(prev => prev + key);
    }
  }
};

export default handleKeyup;
