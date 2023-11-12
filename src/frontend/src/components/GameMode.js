import React from "react";
import { Link } from "react-router-dom";

export default function GameMode() {
  return (
    <div className="gameMode">
      <h1>Vaasle</h1>
      <h2>Select Game Mode</h2>

      <div>
        <Link to="/singleplayer" className="gameMode-button">
          Single Player
        </Link>
        <Link to="/multiplayersetup" className="gameMode-button">
          Multiplayer
        </Link>
      </div>
    </div>
  );
}
