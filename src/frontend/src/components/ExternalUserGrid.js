import React from "react";
// components
import ExternalUserRow from "./ExternalUserRow";

export default function ExternalUserGrid({ corrections, turn, userName }) {
  return (
    <div>
      {/* Render as many rows as there are turns, where the last row is always filled with currentGuess */}
      <p>{userName}</p>
      <p>Turn {turn}</p>
      {corrections && corrections.length > 0 && corrections.map((correction, i) => {
        return (
          <ExternalUserRow key={i} corrections={correction} />
        )
      })}
    </div>
  );
}

