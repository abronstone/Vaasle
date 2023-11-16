import React from "react";
import { useAuth0 } from "@auth0/auth0-react";
import Stats from "./components/Stats";
import Layout from "./components/Layout";
import GameMode from "./components/GameMode";
import Multiplayer from "./components/Multiplayer";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import MultiplayerSetUp from "./components/MultiplayerSetUp";
import Singleplayer from "./components/Singleplayer";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<GameMode />} />
        <Route
          path="/singleplayer"
          element={
            <Singleplayer />
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
        <Route
          path="/multiplayer"
          element={
            <Multiplayer
            />
          }
        />
      </Routes>
    </Router>
  );
}

export default App;
