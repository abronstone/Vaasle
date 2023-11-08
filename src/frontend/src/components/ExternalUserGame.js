import React, { useEffect, useState } from "react";
// components
import ExternalUserGrid from "./ExternalUserGrid";

export default function ExternalUserGame({ externalUserGameGuesses }) {
  const [corrections, setCorrections] = useState([])

  useEffect(() => {
    if (externalUserGameGuesses != null) {
      const newCorrections = []
      // Transform the guesses array into a format that the FE can understand
      for (let turn of externalUserGameGuesses) {
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
    }
  }, [externalUserGameGuesses])

  return (
    <div className="wordle-container">
      <ExternalUserGrid corrections={corrections} turn={corrections.length} />
    </div>
  );
}
