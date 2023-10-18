import React from "react";

interface ModalProps {
  isCorrect: boolean;
  solution: string;
  turn: number;
}

export default function Modal({
  isCorrect,
  solution,
  turn,
}: ModalProps): JSX.Element {
  return (
    <div className="modal">
      {isCorrect && (
        <div>
          <h1>You Win!</h1>
          <p className="solution">{solution}</p>
          <p>You found the solution in {turn} guesses :)</p>
        </div>
      )}
    </div>
  );
}
