import React, { useState, useCallback } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth0 } from "@auth0/auth0-react";
import { newMultiplayerGameApi } from "./util/apiCalls";

export default function MultiplayerSetUp() {
  const { isAuthenticated, user } = useAuth0()

  const [gameID, setGameID] = useState("");
  const navigate = useNavigate();

  const handleCreateLobby = useCallback(async () => {
    if (!isAuthenticated) return;
    const maxGuesses = 6;
    const wordLength = 5;
    // Create a new multiplayer game if the user is authenticated and is not joining a game
    // Then redirect them to the multiplayer page for their newly created game
    const createdGameState = await newMultiplayerGameApi(maxGuesses, wordLength, user.sub)
    navigate(`/multiplayer/${createdGameState.multiplayerGameID}`)
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isAuthenticated, navigate])

  const handleJoin = () => {
    if (gameID) {
      navigate(`/multiplayer/${gameID}`);
    }
  };

  return (
    <div className="multiplayer-setup-container">
      <h1>Vaasle</h1>
      <h2>Create or Join Lobby</h2>
      <div className="multiplayer-setup-options-container">
        <div>
          <button
            onClick={() => handleCreateLobby()}
            className="multiplayer-setup-button"
          >
            Create Lobby
          </button>
        </div>

        <div className="join-lobby-container">
          <input
            type="text"
            className="gameID-input"
            placeholder="Enter game code..."
            value={gameID}
            onChange={(e) => setGameID(e.target.value)}
          />
          <button onClick={handleJoin} className="multiplayer-setup-button" disabled={gameID === null || gameID === ""}>
            Join Lobby
          </button>
        </div>
      </div>
    </div>
  );
}
