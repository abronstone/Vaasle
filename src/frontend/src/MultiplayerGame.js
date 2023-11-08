import React, { useState } from "react";
// components
import ExternalUserGrid from "./components/ExternalUserGrid";

export default function MultiplayerGame() {
  const corrections = [['green', 'green', 'gray', 'yellow', 'gray']]

  return (
    <div className="wordle-container">
      <ExternalUserGrid corrections={corrections} turn={corrections.length} />
    </div>
  );
}
