import { useEffect, useState } from "react";
import useWordle from "../hooks/useWordle";

// components
import Grid from "./Grid";
import Keypad from "./Keypad";
import Modal from "./Modal";

interface Props {
  solution: string;
}

export default function Wordle({ solution }: Props): JSX.Element {
  const { currentGuess, guesses, turn, isCorrect, handleKeyup, usedKeys } =
    useWordle(solution);
  const [showModal, setShowModal] = useState<boolean>(false);

  // attach keyup listening to event object
  useEffect(() => {
    window.addEventListener("keyup", handleKeyup);

    if (isCorrect) {
      console.log("you won the game!");
      setTimeout(() => setShowModal(true), 2000);
      window.removeEventListener("keyup", handleKeyup);
    }

    return () => {
      window.removeEventListener("keyup", handleKeyup);
    };
  }, [isCorrect, handleKeyup]);

  return (
    <div>
      <Grid guesses={guesses} currentGuess={currentGuess} turn={turn} />
      <Keypad usedKeys={usedKeys} />
      {showModal && (
        <Modal isCorrect={isCorrect} solution={solution} turn={turn} />
      )}
    </div>
  );
}
