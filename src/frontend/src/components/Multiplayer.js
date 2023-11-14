import React, { useState, useEffect, useCallback } from "react";
import Layout from "./Layout";
import CurrentUserGame from "./CurrentUserGame";
import ExternalUserGame from "./ExternalUserGame";
import { newGameApi, getExternalUserGamesApi } from "./util/apiCalls";
import { useAuth0 } from "@auth0/auth0-react";

const Multiplayer = () => {
  const { isAuthenticated, user } = useAuth0()
  const [externalGamesState, setExternalGamesState] = useState([])
  const [currentUserGameState, setCurrentUserGameState] = useState(null)
  const [error, setError] = useState(null);

  const initialGameState = useCallback(async () => {
    console.log('initialGameStateFunction entered')
    console.log('isAuthenticated', isAuthenticated)
    try {
      if (!isAuthenticated) return;
      console.log('isAuthenticated!')
      const maxGuesses = 6;
      const wordLength = 5;
      const initialData = await newGameApi(maxGuesses, wordLength, user.sub);
      setCurrentUserGameState(initialData);
      console.log('api res', initialData)
    } catch (error) {
      setError("Failed to initialize game state: " + error);
    }
  }, [isAuthenticated]);

  // Use useEffect to call initialGameState when the component mounts
  useEffect(() => {
    initialGameState();
  }, [initialGameState]);


  //   Fetch state of externalUserGames for multiplayer when the current user submits a guess
  // (the game state in the CurrentUserGame component changes)
  useEffect(() => {
    if (
      currentUserGameState != null &&
      currentUserGameState.metadata.gameID != null
    ) {
      const fetchNewExternalUserGames = async () => {
        try {
          const res = await getExternalUserGamesApi(currentUserGameState.metadata.gameID);
          setExternalGamesState(res);
        } catch (e) {
          setError("Failed to retrieve external user games" + e);
        }
      };

      fetchNewExternalUserGames();
    }
  }, [currentUserGameState]);

  return (
    <Layout>
      <div className="App">
        <div className="CurrentUserGame">
          <CurrentUserGame errorProp={error} gameState={currentUserGameState} setGameState={setCurrentUserGameState} />
        </div>
        {
          externalGamesState?.externalUserGamesMap &&
          Array.from(externalGamesState.externalUserGamesMap.values()).map(
            (game, index) => (
              <div
                className={`ExternalUserGame ${index % 2 === 0 ? "odd" : "even"
                  }`}
                key={index}
              >
                <ExternalUserGame externalUserGameGuesses={game} />
              </div>
            )
          )}
      </div>
    </Layout>
  );
};

export default Multiplayer;
