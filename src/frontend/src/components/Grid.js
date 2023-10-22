import React from "react";
// components
import Row from "./Row";

export default function Grid({ guesses, currentGuess, turn }) {
  return (
    <div>
      {/* Render as many rows as there are turns, where the last row is always filled with currentGuess */}
      <p>Turn {turn}</p>
      {[...Array(turn)].map((_, i) => {
        return (
          <Row key={i} guesses={guesses[i]} />
        )
      })}
      <Row guesses={guesses[turn]} currentGuess={currentGuess} />
    </div>
  );
}

