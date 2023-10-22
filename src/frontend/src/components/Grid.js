import React from "react";
// components
import Row from "./Row";

export default function Grid({ guesses, currentGuess, turn }) {
  console.log("grid rendered");
  console.log("guesses inside grid", guesses);

  return (
    <div>
     {/* Render as many rows as there are turns, where the last row is always filled with currentGuess */}
     <p>Turn {turn}</p>
      {[...Array(turn)].map((_, i) => {
        console.log("i:", i);
        return (
        <Row key={i} guesses={guesses} />
      )})}
      <Row guesses={guesses[turn]} currentGuess={currentGuess} />
    </div>
  );
}

