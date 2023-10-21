import React from "react";
// components
import Row from "./Row";

export default function Grid({ guesses, currentGuess, turn }) {
  console.log("grid rendered");
  return (
    <div>
      {Object.keys(guesses).map((key, i) => {
        if (turn === i) {
          return <Row key={i} currentGuess={currentGuess} />;
        }
        console.log("guesses in Grid", guesses);
        return <Row key={i} guesses={guesses} />;
      })}
    </div>
  );
}
