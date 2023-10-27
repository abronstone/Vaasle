import React from "react";
// components
import Row from "./Row";

export default function Grid({ guesses, currentGuess, turn, status }) {
  console.log("guesses in grid", guesses)
  return (
    <div>
      {/* Render as many rows as there are turns, where the last row is always filled with currentGuess */}
      <p>Turn {turn}</p>
      {guesses && guesses.length > 0 && guesses.map((guess, i) => {
        return (
          <Row key={i} guesses={guess} />
        )
      })}
      {/* Render the current guess row if the game is not over */}
      {(status !== "won" || status !== "lost") && 
        <Row currentGuess={currentGuess} />
      }
    </div>
  );
}

