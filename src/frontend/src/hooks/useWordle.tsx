import { useState } from 'react';

interface LetterObject {
  key: string;
  color: string;
}

interface UseWordleReturn {
  turn: number;
  currentGuess: string;
  guesses: (LetterObject[] | undefined)[];
  isCorrect: boolean;
  handleKeyup: (event: KeyboardEvent) => void;
  usedKeys: { [key: string]: string };
}

const useWordle = (solution: string): UseWordleReturn => {
  const [turn, setTurn] = useState<number>(0);
  const [currentGuess, setCurrentGuess] = useState<string>('');
  const [guesses, setGuesses] = useState<(LetterObject[] | undefined)[]>(Array(6).fill(undefined));
  const [history, setHistory] = useState<string[]>([]);
  const [isCorrect, setIsCorrect] = useState<boolean>(false);
  const [usedKeys, setUsedKeys] = useState<{ [key: string]: string }>({});

  const formatGuess = (): LetterObject[] => {
    let solutionArray = [...solution];
    let formattedGuess: LetterObject[] = [...currentGuess].map((l) => ({
      key: l,
      color: 'grey'
    }));

    formattedGuess.forEach((l, i) => {
      if (solution[i] === l.key) {
        formattedGuess[i].color = 'green';
        solutionArray[i] = '';
      }
    });

    formattedGuess.forEach((l, i) => {
      if (solutionArray.includes(l.key) && l.color !== 'green') {
        formattedGuess[i].color = 'yellow';
        solutionArray[solutionArray.indexOf(l.key)] = '';
      }
    });

    return formattedGuess;
  }

  const addNewGuess = (formattedGuess: LetterObject[]): void => {
    if (currentGuess === solution) {
      setIsCorrect(true);
    }

    setGuesses(prevGuesses => {
      const newGuesses = [...prevGuesses];
      newGuesses[turn] = formattedGuess;
      return newGuesses;
    });

    setHistory(prevHistory => [...prevHistory, currentGuess]);

    setTurn(prevTurn => prevTurn + 1);

    setUsedKeys(prevUsedKeys => {
      const newKeys = { ...prevUsedKeys };
      formattedGuess.forEach(l => {
        const currentColor = newKeys[l.key];
        if (l.color === 'green') {
          newKeys[l.key] = 'green';
          return;
        }
        if (l.color === 'yellow' && currentColor !== 'green') {
          newKeys[l.key] = 'yellow';
          return;
        }
        if (l.color === 'grey' && currentColor !== 'green' && currentColor !== 'yellow') {
          newKeys[l.key] = 'grey';
        }
      });
      return newKeys;
    });

    setCurrentGuess('');
  }

  const handleKeyup = ({ key }: KeyboardEvent): void => {
    // The rest of the logic remains the same...
    // ...
  }

  return { turn, currentGuess, guesses, isCorrect, handleKeyup, usedKeys };
}

export default useWordle;
