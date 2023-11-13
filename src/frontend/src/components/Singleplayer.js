import React, { useState, useEffect, useCallback } from "react";
import Layout from "./Layout";
import CurrentUserGame from "./CurrentUserGame";
import { useAuth0 } from "@auth0/auth0-react";
import { newGameApi } from "./util/apiCalls";

const Singleplayer = () => {
  const { isAuthenticated, user } = useAuth0()
  const [gameState, setGameState] = useState(null);
  const [error, setError] = useState(null);

  const initialGameState = useCallback(async () => {
    try {
      if (!isAuthenticated) return;
      const maxGuesses = 6;
      const wordLength = 5;
      const initialData = await newGameApi(maxGuesses, wordLength, user.sub);
      setGameState(initialData);
    } catch (error) {
      setError("Failed to initialize game state: " + error);
    }
  }, [isAuthenticated]);

  // Use useEffect to call initialGameState when the component mounts
  useEffect(() => {
    initialGameState();
  }, [initialGameState]);


  return (
    <Layout>
      <div className="CurrentUserGame">
        <CurrentUserGame errorProp={error} gameState={gameState} setGameState={setGameState} />
      </div>
    </Layout>
  );
};

export default Singleplayer;
