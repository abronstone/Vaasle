import React from "react";
// components
import CurrentUserRow from "./CurrentUserRow";

export default function CurrentUserGrid({ guesses, currentGuess, turn, status }) {
  return (
    <div>
      {/* Render as many rows as there are turns, where the last row is always filled with currentGuess */}
      <p>Turn {turn}</p>
      {guesses && guesses.length > 0 && guesses.map((guess, i) => {
        return (
          <CurrentUserRow key={i} guesses={guess} />
        )
      })}
      {/* Render the current guess row if the game is not over */}
      {(status !== "won" || status !== "lost") && 
        <CurrentUserRow currentGuess={currentGuess} />
      }
    </div>
  );
}

