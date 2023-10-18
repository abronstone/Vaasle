import React from "react";
import Row from "./Row";

interface Props {
  guesses: Array<string>;
  currentGuess: string;
  turn: number;
}

export default function Grid({
  guesses,
  currentGuess,
  turn,
}: Props): JSX.Element {
  return (
    <div>
      {guesses.map((g: string, i: number) => {
        if (turn === i) {
          return <Row key={i} currentGuess={currentGuess} />;
        }
        return <Row key={i} guess={g} />;
      })}
    </div>
  );
}
