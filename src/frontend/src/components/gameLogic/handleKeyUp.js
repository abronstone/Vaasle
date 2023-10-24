import { makeGuessApi } from "../util/apiCalls";  

const handleKeyup = async(key, currentGuess, setCurrentGuess, guesses, setGuesses, turn, setTurn, gameState, setGameState, isCorrect, setIsCorrect, usedKeys, setUsedKeys) => {

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
    // check word is 5 chars
    if (currentGuess.length !== 5) {
      console.log('word must be 5 chars.');
      return;
    }
    try {

      // When a guess is submitted, the API called to get a new game state
      // and all necessary state is updated
      const newGameState = await makeGuessApi(gameState.metadata.id, currentGuess);

      setGameState(newGameState);
      setCurrentGuess('');

      setTurn(gameState.guesses.length)

      const guessesMap = new Map()

      
      // 
      newGameState.guesses.forEach((guess) => {
        const formattedGuess = Array.from(guess[0]).map((letter, index) => {
          const colorCode = guess[1][index];
          return {
            letter,
            color: colorCode === 'G' ? 'green' : colorCode === 'Y' ? 'yellow' : 'grey'
          };
        });
        guessesMap.set(guess[0], formattedGuess);
      });
    
      setGuesses(Object.fromEntries(guessesMap))
      
      setIsCorrect(gameState.state === "won")

        // make a map of all the unique used letter color pairs from guesses
      const usedKeysMap = new Map()
      
      Object.keys(guesses).forEach((key) => {
        usedKeysMap.set(key, guesses[key])
      })


    setUsedKeys(Object.fromEntries(usedKeysMap))

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
