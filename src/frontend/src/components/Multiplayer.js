import React, { useState, useEffect, useCallback } from "react";
import Layout from "./Layout";
import CurrentUserGame from "./CurrentUserGame";
import ExternalUserGame from "./ExternalUserGame";
import { refreshMultiplayerGameApi, joinMultiplayerGameApi, getGameApi, startMultiplayerGameApi } from "./util/apiCalls";
import { useAuth0 } from "@auth0/auth0-react";
import { useParams, useLocation } from "react-router-dom";
import MultiplayerModal from "./MultiplayerModal";
import ErrorBadge from "./ErrorBadge";

const Multiplayer = () => {
  const { isAuthenticated, user } = useAuth0();
  const { multiplayerGameId } = useParams();
  const location = useLocation();
  const queryParams = new URLSearchParams(location.search);
  const isHost = queryParams.get('host') === 'true';

  const [multiplayerGameState, setMultiplayerGameState] = useState(null);
  const [externalGamesState, setExternalGamesState] = useState(null);
  const [currentUserGameState, setCurrentUserGameState] = useState(null);
  const [error, setError] = useState(null);
  const [hasGameStarted, setHasGameStarted] = useState(false);
  const [externalUserHasWon, setExternalUserHasWon] = useState(false);

  const handleError = (message, error) => {
    console.error(message + error);
    setError(message + error);
  };

  const getInitialMultiplayerGameState = useCallback(async () => {
    try {
      if (!isAuthenticated) return;
      const maxGuesses = 6;
      const wordLength = 5;
      console.log('attempting to join game for user: ', user.sub)
      const joinedGameState = await joinMultiplayerGameApi(multiplayerGameId, maxGuesses, wordLength, user.sub);
      if (joinedGameState.state != null && joinedGameState.state === "won" && joinedGameState.winnerID !== user.sub) {
        setExternalUserHasWon(true)
      }
      setMultiplayerGameState(joinedGameState);
    } catch (error) {
      handleError("Failed to initialize multiplayer game: ", error);
    }
  }, [isAuthenticated, multiplayerGameId, user.sub]);
  
  useEffect(() => {
    getInitialMultiplayerGameState();
  }, [getInitialMultiplayerGameState]);

  const getInitialCurrentUserGameState = useCallback(async () => {
    try {
      const currentUserGameId = multiplayerGameState?.games[user.sub];
      if (currentUserGameId == null) {
        throw new Error("No current user game ID found");
      }  
      const res = await getGameApi(currentUserGameId);
      setCurrentUserGameState(res);
    } catch (error) {
      handleError("Failed to initialize your game: ", error);
    }  
  }, [multiplayerGameState, user.sub]);  

  const handleStart = useCallback(async () => {
    try {
      setTimeout(async () => {
        const res = await startMultiplayerGameApi(multiplayerGameId);
        if (res == null || !res) {
          throw new Error("Backend failed to start new multiplayer game");
        }
        getInitialCurrentUserGameState();
        setHasGameStarted(true);
      }, 1000)
    } catch (error) {
      handleError("Failed to start new multiplayer game", error);
    }
  }, [multiplayerGameId, getInitialCurrentUserGameState]);


  const fetchNewExternalUserGames = useCallback(async () => {
    try {
      if (multiplayerGameId == null) {
        throw new Error("No multiplayer game id found");
      }
      const refreshedMultiplayerGameState = await refreshMultiplayerGameApi(multiplayerGameId);
      console.log('refreshedMultiplayerGameState: ', refreshedMultiplayerGameState)

      if (refreshedMultiplayerGameState.state != null && refreshedMultiplayerGameState.state === "won" && refreshedMultiplayerGameState.winnerID !== user.sub) {
        setExternalUserHasWon(true)
      }

      const externalUserGamesMap = new Map(
        Object.entries(refreshedMultiplayerGameState.userCorrections).filter(([key]) => key !== user.sub)
      );
      const externalUserIdsToNamesMap = new Map(Object.entries(refreshedMultiplayerGameState.userNames))

      setExternalGamesState({ state: refreshedMultiplayerGameState.state, externalUserGamesMap, externalUserIdsToNamesMap });
      setMultiplayerGameState(refreshedMultiplayerGameState)
    } catch (e) {
      handleError("Failed to retrieve external user games", e);
    }
  }, [multiplayerGameId, user.sub]);

  useEffect(() => {
    fetchNewExternalUserGames();
  }, [currentUserGameState, fetchNewExternalUserGames]);

  // Render a grid representation of every external user's guesses 
  const renderExternalUserGames = () => {
    const games = externalGamesState?.externalUserGamesMap;
    return games && Array.from(games.entries()).map(([user, game], index) => (
      <div
        className={`ExternalUserGame ${index % 2 === 0 ? "odd" : "even"}`}
        key={index}
      >
        <ExternalUserGame externalUserGameGuesses={game} userName={externalGamesState.externalUserIdsToNamesMap.get(user) ?? 'Guest'} />
      </div>
    ));
  };

  const handleRefreshLobby = useCallback(async () => {
    try {
      if (multiplayerGameId == null) {
        throw new Error("No multiplayer game id found");
      }
      const refreshedMultiplayerGameState = await refreshMultiplayerGameApi(multiplayerGameId);
      console.log('refreshedMultiplayerGameState: ', refreshedMultiplayerGameState);

      // Common logic for setting state
      const externalUserGamesMap = new Map(
        Object.entries(refreshedMultiplayerGameState.userCorrections).filter(([key]) => key !== user.sub)
      );
      const externalUserIdsToNamesMap = new Map(Object.entries(refreshedMultiplayerGameState.userNames));
      setExternalGamesState({ state: refreshedMultiplayerGameState.state, externalUserGamesMap, externalUserIdsToNamesMap });
      setMultiplayerGameState(refreshedMultiplayerGameState);

      switch (refreshedMultiplayerGameState.state) {
        case "waiting":
          break;
        case "ongoing":
          handleStart()
          break;
        case "won":
          throw new Error("The game you are trying to join has already finished");
        case "lost":
          throw new Error("The game you are trying to join has already finished");
        default:
          console.warn(`Unhandled game state: ${refreshedMultiplayerGameState.state}`);
          return;
      }

      setError(null); // Clearing any existing errors

    } catch (error) {
      handleError("Failed to refresh lobby: ", error);
    }
  }, [multiplayerGameId, user.sub]);

  useEffect(() => {
    handleRefreshLobby()
  }, [])

  // Render a list of all current users in the lobby
  const renderCurrentUsersInLobby = () => {
    return (
      <>
        <h3>Players</h3>
        {externalGamesState != null && externalGamesState.externalUserIdsToNamesMap != null && Array.from(externalGamesState.externalUserIdsToNamesMap?.values() ?? []).map((userName) => (
          <p key={userName}>{userName}</p>
        ))}
      </>
    );
  };


  if (!hasGameStarted && isHost) {
    return (
      <Layout>
        {error != null && <ErrorBadge text={error} />}
        <h3>Send this code to your friends to allow them to join!</h3>
        <p>{multiplayerGameId}</p>
        <p>Once all of your friends are on this screen everyone start the game at the same time!</p>
        <div>
          <button onClick={handleStart} className="general-rounded-button">Start Game</button>
        </div>
        <div>
          <button onClick={handleRefreshLobby} className="general-rounded-button">Refresh Lobby</button>
        </div>
        {renderCurrentUsersInLobby()}
      </Layout>
    );
  }

  else if (!hasGameStarted) {
    return (
      <Layout>
        {error != null && <ErrorBadge text={error} />}
        <h3>Please wait for the host to join this game</h3>
        <p>Click refresh to refresh the lobby, you will automatically enter the game after pressing this once the host starts the game</p>
        <p>While you wait, send your friends the code to join!</p>
        <p>{multiplayerGameId}</p>
        <p>Once all of your friends are on this screen everyone start the game at the same time!</p>
        <div>
          <button onClick={handleRefreshLobby} className="general-rounded-button">Refresh Lobby</button>
        </div>
        {renderCurrentUsersInLobby()}
      </Layout>
    )
  }

  else {
    console.log("Current User Game State: ", currentUserGameState); // Add this for debugging
  
    return (
      <>
        <Layout>
          {/* ... rest of your code */}
          <div>
            {currentUserGameState != null ? (
              <div className="multiplayer-container">
                <div className="CurrentUserGame">
                  <CurrentUserGame errorProp={error} gameState={currentUserGameState} setGameState={setCurrentUserGameState} />
                </div>
              </div>
            ) : (
              <p>Loading game...</p> // This will indicate if the state is still null
            )}
            {renderExternalUserGames()}
          </div>
        </Layout>
        {/* ... rest of your code */}
      </>
    );
  }
  };

export default Multiplayer;
