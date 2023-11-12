import React, { useState, useEffect, useCallback } from "react";
import CurrentUserGame from "./components/CurrentUserGame";
import ExternalUserGame from "./components/ExternalUserGame";
import {
  newGameApi,
  getExternalUserGamesApi,
} from "./components/util/apiCalls";
import { useAuth0 } from "@auth0/auth0-react";
import ErrorBadge from "./components/ErrorBadge";
import Stats from "./components/Stats";
import Layout from "./components/Layout";
import GameMode from "./components/GameMode";
import { BrowserRouter as Router, Route, Link, Routes } from "react-router-dom";
import MultiplayerSetUp from "./components/MultiplayerSetUp";

function App() {
  const { isAuthenticated, user } = useAuth0();
  const [gameState, setGameState] = useState(null);
  const [externalGamesState, setExternalGamesState] = useState(null);
  const [error, setError] = useState(null);
  // TODO replace this with multiple pages
  const [isMultiplayerEnabled, setIsMultiplayerEnabled] = useState(true);

  const initialGameState = useCallback(async () => {
    try {
      if (!isAuthenticated) return;
      const maxGuesses = 5;
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

  // Fetch state of externalUserGames for multiplayer when the current user submits a guess
  // (the game state in the CurrentUserGame component changes)
  useEffect(() => {
    if (
      gameState != null &&
      gameState.metadata.gameID != null &&
      isMultiplayerEnabled
    ) {
      const fetchNewExternalUserGames = async () => {
        try {
          const res = await getExternalUserGamesApi(gameState.metadata.gameID);
          setExternalGamesState(res);
        } catch (e) {
          setError("Failed to retrieve external user games" + e);
        }
      };

      fetchNewExternalUserGames();
    }
  }, [gameState, isMultiplayerEnabled]);

  // TODO make this component the landing page for two separate pages
  // 1) A single player page where CurrentUserGame is the only component (with Vassle header still)
  // 2) A multiplayer page where the CurrentUserGame and x amount of ExternalUserGame components
  // are on the same page
  return (
    <Router>
      <Routes>
        <Route path="/" element={<GameMode />} />
        <Route
          path="/singleplayer"
          element={
            <Layout>
              <div className="App">
                <CurrentUserGame
                  gameState={gameState}
                  setGameState={setGameState}
                />
                {isMultiplayerEnabled &&
                  isAuthenticated &&
                  externalGamesState != null &&
                  Array.from(
                    externalGamesState.externalUserGamesMap.values()
                  ).map((game, index) => (
                    <div
                      className={`ExternalUserGame ${
                        index % 2 === 0 ? "odd" : "even"
                      }`}
                    >
                      <ExternalUserGame externalUserGameGuesses={game} />
                    </div>
                  ))}
              </div>
            </Layout>
          }
        />

        <Route
          path="/stats"
          element={
            <Layout>
              <Stats />
            </Layout>
          }
        />
        <Route path="/multiplayersetup" element={<MultiplayerSetUp />} />
      </Routes>
    </Router>
  );
}

export default App;
