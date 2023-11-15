import React, { useState, useEffect, useCallback } from "react";
import Layout from "./Layout";
import CurrentUserGame from "./CurrentUserGame";
import ExternalUserGame from "./ExternalUserGame";
import { refreshMultiplayerGameApi, joinMultiplayerGameApi, getGameApi, startMultiplayerGameApi } from "./util/apiCalls";
import { useAuth0 } from "@auth0/auth0-react";
import { useParams } from "react-router-dom";

const Multiplayer = () => {
  const { isAuthenticated, user } = useAuth0()
  const { multiplayerGameId } = useParams()

  const [multiplayerGameState, setMultiplayerGameState] = useState(null)
  const [externalGamesState, setExternalGamesState] = useState([])
  const [currentUserGameState, setCurrentUserGameState] = useState(null)
  const [error, setError] = useState(null);
  const [hasGameStarted, setHasGameStarted] = useState(false)

  const getInitialMultiplayerGameState = useCallback(async () => {
    try {
      if (!isAuthenticated) return
      const maxGuesses = 6
      const wordLength = 5
      const joinedGameState = await joinMultiplayerGameApi(multiplayerGameId, maxGuesses, wordLength, user.sub)
      console.log('joinMultiplayerGameApi res', joinedGameState)
      setMultiplayerGameState(joinedGameState)

      console.log('multiplayerGameState,', multiplayerGameState)
    } catch (error) {
      console.error("Failed to initialize multiplayer game: " + error);
      setError("Failed to initialize multiplayer game: " + error);
    }
  }, [isAuthenticated, multiplayerGameId, multiplayerGameState]);

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
  }, [multiplayerGameState, currentUserGameState])

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
      const res = await startMultiplayerGameApi(multiplayerGameId)
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

  const fetchNewExternalUserGames = useCallback(async () => {
    console.log('executing fetchNewExternalUserGames!')
    try {
      if (multiplayerGameId == null) {
        throw new Error("No multiplayer game id found")
      }

      const res = await refreshMultiplayerGameApi(multiplayerGameId);
      console.log('res from refreshMultiplayerGameApi', res)

      // Transform res.userCorrections into a map of userIds onto their array of guesses
      // with the current user filtered out
      // Eg. '123456' : [['G', 'Y', 'X', 'G', 'Y'],['Y', 'Y', 'X', 'Y', 'Y']]
      const externalUserGamesMap = new Map(
        Object.entries(res.userCorrections).filter(([key, value]) => key !== user.sub)
      );

      const externalUserGamesObject = {
        state: res.state,
        externalUserGamesMap
      }

      setExternalGamesState(externalUserGamesObject);
    } catch (e) {
      console.error("Failed to retrieve external user games" + e)
      setError("Failed to retrieve external user games" + e);
    }
  }, [multiplayerGameId, currentUserGameState, multiplayerGameState])

  //   Fetch state of externalUserGames for multiplayer when the current user submits a guess
  // (the game state in the CurrentUserGame component changes)
  useEffect(() => {
    console.log('fetchNewExternalUserGames useEffect triggered!')
    fetchNewExternalUserGames();
  }
    , [currentUserGameState]);

  if (!hasGameStarted) {
    return (
      <Layout>
        <div>
          <button onClick={handleStart} className="general-rounded-button">Start Game</button>
        </div>
      </Layout>
    )
  }
  return (
    <Layout>
      <h3>
        Send this code to your friends to allow them to join!
      </h3>
      <p>{multiplayerGameId}</p>
      <div className="multiplayer-container">
        <div className="CurrentUserGame">
          <CurrentUserGame errorProp={error} gameState={currentUserGameState} setGameState={setCurrentUserGameState} />
        </div>
        {externalGamesState?.externalUserGamesMap && Array.from(externalGamesState.externalUserGamesMap.entries()).map(
          ([user, game], index) => (
            <div
              className={`ExternalUserGame ${index % 2 === 0 ? "odd" : "even"
                }`}
              key={index}
            >
              <ExternalUserGame externalUserGameGuesses={game} userName={user} />
            </div>
          )
        )}
      </div>
    </Layout>
  );
};

export default Multiplayer;
