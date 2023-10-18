import { useState, useEffect } from "react";

interface UseWordleReturn {
  currentGuess: string;
  guesses: Array<string>;
  turn: number;
  isCorrect: boolean;
  handleKeyup: (event: KeyboardEvent) => void;
  usedKeys: { [key: string]: boolean };
}

const useWordle = (solution: string): UseWordleReturn => {
  // track user turn
  const [turn, setTurn] = useState<number>(0);

  // track what user is guessing
  const [currentGuess, setCurrentGuess] = useState<string>("");

  // each guess is an array, having length of array be 6
  const [guesses, setGuesses] = useState<Array<string>>([...Array(6)]);

  // history of guesses of their
  const [history, setHistory] = useState<Array<string>>([]);

  // track if the guess is correct
  const [isCorrect, setIsCorrect] = useState<boolean>(false);

  // coloring used keys, values are color keys are the letter
  const [usedKeys, setUsedKeys] = useState<{ [key: string]: boolean }>({});

  // handle keyup event
  const handleKeyup = (event: KeyboardEvent): void => {
    const { key } = event;

    if (/[a-z]/.test(key) && !usedKeys[key]) {
      const newGuess = currentGuess + key;
      setCurrentGuess(newGuess);

      if (newGuess.length === solution.length) {
        const newGuesses = [...guesses];
        newGuesses[turn] = newGuess;
        setGuesses(newGuesses);

        const newHistory = [...history, newGuess];
        setHistory(newHistory);

        if (newGuess === solution) {
          setIsCorrect(true);
        } else {
          setTurn(turn + 1);
          setCurrentGuess("");
          setUsedKeys({ ...usedKeys, [key]: true });
        }
      } else {
        setUsedKeys({ ...usedKeys, [key]: true });
      }
    }
  };

  // attach keyup listening to event object
  useEffect(() => {
    window.addEventListener("keyup", handleKeyup);

    if (isCorrect) {
      console.log("you won the game!");
      window.removeEventListener("keyup", handleKeyup);
    }

    return () => {
      window.removeEventListener("keyup", handleKeyup);
    };
  }, [isCorrect, handleKeyup]);

  return { currentGuess, guesses, turn, isCorrect, handleKeyup, usedKeys };
};

export default useWordle;
