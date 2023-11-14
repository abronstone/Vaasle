import React, { useState, useEffect, useCallback } from "react";
import Layout from "./Layout";
import CurrentUserGame from "./CurrentUserGame";
import ExternalUserGame from "./ExternalUserGame";
import { newMultiplayerGameApi, getExternalUserGamesApi, joinMultiplayerGameApi, getGameApi, startMultiplayerGameApi } from "./util/apiCalls";
import { useAuth0 } from "@auth0/auth0-react";
import { useParams } from "react-router-dom";

const Multiplayer = () => {
  const { isAuthenticated, user } = useAuth0()
  const [multiplayerGameState, setMultiplayerGameState] = useState(null)
  const [externalGamesState, setExternalGamesState] = useState([])
  const [currentUserGameState, setCurrentUserGameState] = useState(null)
  const [error, setError] = useState(null);
  const [hasGameStarted, setHasGameStarted] = useState(false)
  const [isHost, setIsHost] = useState(false)
  const { multiplayerGameId } = useParams()

  const getInitialMultiplayerGameState = useCallback(async () => {
    try {
      if (!isAuthenticated) return;
      const maxGuesses = 6;
      const wordLength = 5;
      // Create a new multiplayer game if the user is authenticated and is not joining a game
      if (multiplayerGameId == null) {
        const createdGameState = await newMultiplayerGameApi(maxGuesses, wordLength, user.sub)
        console.log('newMultiplayerGameApi res', createdGameState)
        setMultiplayerGameState(createdGameState);
        setIsHost(true)
      }
      // Join an existing multiplayer game if the user is authenticated and has the id for
      // some one elses game
      else {
        const joinedGameState = await joinMultiplayerGameApi(multiplayerGameId, maxGuesses, wordLength, user.sub)
        console.log('joinMultiplayerGameApi res', joinedGameState)
        setMultiplayerGameState(joinedGameState)
      }
      console.log('multiplayerGameState,', multiplayerGameState)
    } catch (error) {
      console.error("Failed to initialize multiplayer game: " + error);
      setError("Failed to initialize multiplayer game: " + error);
    }
  }, [isAuthenticated, multiplayerGameId, user.sub, multiplayerGameState]);

  const getInitialCurrentUserGameState = useCallback(async () => {
    try {
      // if (multiplayerGameState == null) {
      //   throw new Error("no multiplayer game state found")
      // }
      // if (multiplayerGameState.games == null) {
      //   throw new Error("no games found in multiplayer game state")
      // }
      const currentUserGameId = multiplayerGameState.games[user.sub]
      console.log('currentUserGameId', currentUserGameId)
      const res = await getGameApi(currentUserGameId)
      console.log('get game for current user res,', res)
      setCurrentUserGameState(res)
      console.log('currentUserGameState', currentUserGameState)
    } catch (error) {
      console.error("Failed to initialize your game: " + error);
      setError("Failed to initialize your game: " + error);
    }
  }, [multiplayerGameState, user.sub, currentUserGameState])

  // Use useEffect to call getInitialMultiplayerGameState when the component mounts
  useEffect(() => {
    getInitialMultiplayerGameState();
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  useEffect(() => {
    console.log('multiplayerGameState updated:', multiplayerGameState);
  }, [multiplayerGameState]);
  
  // Start a multiplayer game
  const handleStart = useCallback(async () => {
    try {
      const res = await startMultiplayerGameApi(multiplayerGameState.multiplayerGameID)
      if (res != null && !res) {
        throw new Error("Backend failed to start new multiplayer game")
      }
      getInitialCurrentUserGameState()
      setHasGameStarted(true)
    }
    catch (error) {
      console.error("Failed to start new multiplayer game" + error)
      setError("Failed to start new multiplayer game" + error)
    }
  }, [multiplayerGameId, getInitialCurrentUserGameState])



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

  if (!hasGameStarted) {
    return (
      <Layout>
        <div className="App">
          <button onClick={handleStart}>Start Game</button>
        </div>
      </Layout>
    )
  }
  return (
    <Layout>
      <div className="App">
        <div className="CurrentUserGame">
          <CurrentUserGame errorProp={error} gameState={currentUserGameState} setGameState={setCurrentUserGameState} />
        </div>
        {externalGamesState?.externalUserGamesMap && Array.from(externalGamesState.externalUserGamesMap.values()).map(
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
