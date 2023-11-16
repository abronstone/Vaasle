import React, { useEffect, useState } from "react";
// components
import ExternalUserGrid from "./ExternalUserGrid";

export default function ExternalUserGame({ externalUserGameGuesses, userName }) {
  const [corrections, setCorrections] = useState([])
  const [turnCount, setTurnCount] = useState(0)

  useEffect(() => {
    if (externalUserGameGuesses != null) {
      const newCorrections = []
      let turnCounter = 0
      // Transform the guesses array into a format that the FE can understand
      for (let turn of externalUserGameGuesses) {
        if(turn === "") continue
        turnCounter++
        const fullColorNameGuesses = []
        for (let colorCode of turn) {

          // Determine the color needed for FE based on the color code from the API
          const color =
            colorCode === "G"
              ? "green"
              : colorCode === "Y"
                ? "yellow"
                : "grey";
          fullColorNameGuesses.push(color)
        }
        newCorrections.push(fullColorNameGuesses)

      }
      setCorrections(newCorrections)
      setTurnCount(turnCounter)
    }
  }, [externalUserGameGuesses])

  return (
    <div className="wordle-container">
      <ExternalUserGrid corrections={corrections} turn={turnCount} userName={userName} />
    </div>
  );
}
