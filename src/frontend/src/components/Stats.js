import React from "react";
import { Link } from "react-router-dom";

export default function Stats() {
  // TODO   make statsData to a useState hook calling gameStruct object
  // Populating dummy data
  const statsData = {
    gamesPlayed: 10,
    winPercentage: 70, // In percentage
    mostGuessedWord: "hello",
  };

  return (
    <div className="stats-container">
      <h2>Statistics</h2>
      <ul>
        <li>Games Played: {statsData.gamesPlayed}</li>
        <li>Win %: {statsData.winPercentage}</li>
        <li>Most Guessed Word: {statsData.mostGuessedWord}</li>
      </ul>
      <Link to="/" className="close-btn">
        X
      </Link>
    </div>
  );
}
