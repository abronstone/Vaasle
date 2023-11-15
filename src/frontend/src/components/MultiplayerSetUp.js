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
  }, [isAuthenticated, navigate])

  const handleJoin = () => {
    if (gameID) {
      navigate(`/multiplayer/${gameID}`);
    }
  };

  return (
    <div className="multiplayerSettings">
      <h1>Vaasle</h1>
      <h2>Create or Join Lobby</h2>

      <div>
        <button
          onClick={() => handleCreateLobby()}
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
