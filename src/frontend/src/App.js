import React, { useState, useEffect, useCallback } from "react";
import CurrentUserGame from "./components/CurrentUserGame";
import ExternalUserGame from "./components/ExternalUserGame";
import {
  newGameApi,
  getExternalUserGamesApi,
} from "./components/util/apiCalls";
import { useAuth0 } from "@auth0/auth0-react";
import Stats from "./components/Stats";
import Layout from "./components/Layout";
import GameMode from "./components/GameMode";
import Multiplayer from "./components/Multiplayer";
import { BrowserRouter as Router, Route, Link, Routes } from "react-router-dom";
import MultiplayerSetUp from "./components/MultiplayerSetUp";
import Singleplayer from "./components/Singleplayer";

function App() {
  const { isAuthenticated, user } = useAuth0();
  const [externalGamesState, setExternalGamesState] = useState(null);
  const [isMultiplayerEnabled, setIsMultiplayerEnabled] = useState(true);

  // Fetch state of externalUserGames for multiplayer when the current user submits a guess
  // (the game state in the CurrentUserGame component changes)
  // useEffect(() => {
  //   if (
  //     gameState != null &&
  //     gameState.metadata.gameID != null &&
  //     isMultiplayerEnabled
  //   ) {
  //     const fetchNewExternalUserGames = async () => {
  //       try {
  //         const res = await getExternalUserGamesApi(gameState.metadata.gameID);
  //         setExternalGamesState(res);
  //       } catch (e) {
  //         setError("Failed to retrieve external user games" + e);
  //       }
  //     };

  //     fetchNewExternalUserGames();
  //   }
  // }, [gameState, isMultiplayerEnabled]);

  return (
    <Router>
      <Routes>
        <Route path="/" element={<GameMode />} />
        <Route
          path="/singleplayer"
          element={
            <Singleplayer   />
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
        {/* <Route path="/multiplayersetup" element={<MultiplayerSetUp />} />
        <Route
          path="/multiplayer"
          element={
            <Multiplayer
              gameState={gameState}
              setGameState={setGameState}
              externalGamesState={externalGamesState}
              isMultiplayerEnabled={isMultiplayerEnabled}
              isAuthenticated={isAuthenticated}
            />
          }
        /> */}
      </Routes>
    </Router>
  );
}

export default App;
