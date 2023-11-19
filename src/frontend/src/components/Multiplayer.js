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

  // State variables
  // Contains response from /refreshGame endpoint  modified to be useful to the FE
  const [externalGamesState, setExternalGamesState] = useState(null);
  // Contains unmodified response from /getGame endpoint for the current user's game
  const [currentUserGameState, setCurrentUserGameState] = useState(null);
  // Error handling and flags
  const [error, setError] = useState(null);
  const [hasGameStarted, setHasGameStarted] = useState(false);
  const [externalUserHasWon, setExternalUserHasWon] = useState(false);

  const handleError = (message, error) => {
    console.error(message + error);
    setError(message + error);
  };

  // Try and join the multiplayer game specified with the id in the url 
  // and get the state of the game assigned to the current user and pass
  // it to the CurrentUserGame component
  const getInitialMultiplayerGameState = useCallback(async () => {
    try {
      if (!isAuthenticated) return;
      // Values are hardcoded for now but could be dynamically set
      const maxGuesses = 6;
      const wordLength = 5;

      const joinedGameState = await joinMultiplayerGameApi(multiplayerGameId, maxGuesses, wordLength, user.sub);
      // Handle if the game has already been won by someone else
      if (joinedGameState.state != null && joinedGameState.state === "won" && joinedGameState.winnerID !== user.sub) {
        setExternalUserHasWon(true)
      }

      const currentUserGameId = joinedGameState?.games[user.sub];
      if (currentUserGameId == null) {
        throw new Error("No current user game ID found");
      }
      const currentUserGame = await getGameApi(currentUserGameId);
      setCurrentUserGameState(currentUserGame);
    } catch (error) {
      handleError("Failed to initialize multiplayer game: ", error);
    }
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isAuthenticated, multiplayerGameId]);
  
  useEffect(() => {
    getInitialMultiplayerGameState();
  }, [getInitialMultiplayerGameState]);

  // Start a game and update the FE's internal state
  const handleStart = useCallback(async () => {
    try {
        const startGameRes = await startMultiplayerGameApi(multiplayerGameId);
        if (startGameRes == null || !startGameRes) {
          throw new Error("Backend failed to start new multiplayer game");
        }
        setHasGameStarted(true);
    } catch (error) {
      handleError("Failed to start new multiplayer game", error);
    }
  }, [multiplayerGameId]);

  // Refresh a multiplayer game by calling /refreshMultiplayerGame
  // Take the relevant keys of its response object and setExternalGamesState with
  // FE friendly versions of these keys
  const fetchNewExternalUserGames = useCallback(async () => {
    try {
      if (multiplayerGameId == null) {
        throw new Error("No multiplayer game id found");
      }
      const refreshedMultiplayerGameState = await refreshMultiplayerGameApi(multiplayerGameId);

      // If someone else has won the game, pop update the flag to popup the MultiplayerModal
      if (refreshedMultiplayerGameState.state != null && refreshedMultiplayerGameState.state === "won" && refreshedMultiplayerGameState.winnerID !== user.sub) {
        setExternalUserHasWon(true)
      }

      // Transform userCorrections and userNames into FE friendly equivalents
      const externalUserGamesMap = new Map(
        Object.entries(refreshedMultiplayerGameState.userCorrections).filter(([key]) => key !== user.sub)
      );
      const externalUserIdsToNamesMap = new Map(Object.entries(refreshedMultiplayerGameState.userNames))

      setExternalGamesState({ state: refreshedMultiplayerGameState.state, externalUserGamesMap, externalUserIdsToNamesMap, winnerID: refreshedMultiplayerGameState.winnerID, word: refreshedMultiplayerGameState.word });
    } catch (e) {
      handleError("Failed to retrieve external user games", e);
    }
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [multiplayerGameId]);

  useEffect(() => {
    fetchNewExternalUserGames();
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [currentUserGameState]);

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

  // Render different lobby screens based on whether the user is a host or not
  if (!hasGameStarted && isHost) {
    return (
      <Layout>
        {error != null && <ErrorBadge text={error} />}
        <h3>Send this code to your friends to allow them to join!</h3>
        <p>{multiplayerGameId}</p>
        <p>Once you confirm that all of your friends are on this screen, tell everyone to start the game at the same time!</p>
        <div>
          <button onClick={handleStart} className="general-rounded-button">Start Game</button>
        </div>
      </Layout>
    );
  }

  else if (!hasGameStarted) {
    return (
      <Layout>
        {error != null && <ErrorBadge text={error} />}
        <h3>Please wait for the host to tell everyone to start, then click start game</h3>
        <p>While you wait, send your friends the code to join!</p>
        <p>{multiplayerGameId}</p>
        <div>
          <button onClick={handleStart} className="general-rounded-button">Start Game</button>
        </div>
      </Layout>
    )
  }

  // Once the game has started, render the main game screen
  else {
    return (
      <>
        <Layout>
          <h3>Send this code to your friends to allow them to join!</h3>
          <p>{multiplayerGameId}</p>
          <div className="multiplayer-container">
            <div className="CurrentUserGame">
              <CurrentUserGame errorProp={error} gameState={currentUserGameState} setGameState={setCurrentUserGameState} />
            </div>
            {renderExternalUserGames()}
          </div>
        </Layout>
        {externalUserHasWon && <MultiplayerModal winner={externalGamesState.externalUserIdsToNamesMap.get(externalGamesState.winnerID)} solution={externalGamesState.word} />}
      </>
    );
  }
};

export default Multiplayer;
