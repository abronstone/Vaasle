import React from "react";
// components
import Row from "./Row";

export default function Grid({ guesses, currentGuess, turn, status }) {
  return (
    <div>
      {/* Render as many rows as there are turns, where the last row is always filled with currentGuess */}
      <p>Turn {turn}</p>
      {guesses.map((guess, i) => {
        return (
          <Row key={i} guesses={guess} />
        )
      })}
      <Row guesses={currentGuess} />
    </div>
  );
}

