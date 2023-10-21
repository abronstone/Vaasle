const useWordle = (gameState) => {
  if (gameState == null || gameState.guesses == null)
    return { turn: 0, guesses: [], isCorrect: false, usedKeys: [] };

  const turn = gameState.guesses.length;

  const guessesMap = new Map();

  gameState.guesses.forEach((guess) => {
    const formattedGuess = Array.from(guess[0]).map((letter, index) => {
      const colorCode = guess[1][index];
      return {
        letter,
        color:
          colorCode === "G" ? "green" : colorCode === "Y" ? "yellow" : "grey",
      };
    });
    guessesMap.set(guess[0], formattedGuess);
  });

  const guesses = Object.fromEntries(guessesMap);

  const isCorrect = gameState.state === "won";

  // make a map of all the unique used letter color pairs from guesses
  const usedKeysMap = new Map();

  Object.keys(guesses).forEach((key) => {
    usedKeysMap.set(key, guesses[key]);
  });

  const usedKeys = Object.fromEntries(usedKeysMap);

  return { turn, guesses, isCorrect, usedKeys };
};

export default useWordle;
