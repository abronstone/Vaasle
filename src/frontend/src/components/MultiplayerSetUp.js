import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";

export default function MultiplayerSetUp() {
  const [gameID, setGameID] = useState("");
  const navigate = useNavigate();

  // TODO: have user navigate to multiplayer, but dont expose gameID in url
  const handleJoin = () => {
    if (gameID) {
      navigate(`/multiplayer/${gameID}`);
    }
  };

  return (
    <div className="multiplayerSettings">
      <Link to="/">
        <h1>Vaasle</h1>
      </Link>
      <h2>Create or Join Lobby</h2>

      <div>
        <button
          onClick={() => navigate("/multiplayer")}
          className="gameMode-button"
        >
          Create Lobby
        </button>

        <input
          type="text"
          className="gameID-input"
          placeholder="Enter GameID"
          value={gameID}
          onChange={(e) => setGameID(e.target.value)}
        />
        <button onClick={handleJoin} className="gameMode-button">
          Join Lobby
        </button>
      </div>
    </div>
  );
}
