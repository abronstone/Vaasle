import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

export default function MultiplayerSetUp() {
  const [gameID, setGameID] = useState("");
  const navigate = useNavigate();

  const handleJoin = () => {
    if (gameID) {
      navigate(`/multiplayersetup/${gameID}`);
    }
  };

  return (
    <div className="multiplayerSettings">
      <h1>Vaasle</h1>
      <h2>Create or Join Lobby</h2>

      <div>
        <button
          onClick={() => navigate("/singleplayer")}
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
