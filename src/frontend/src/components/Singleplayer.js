import React from "react";
import Layout from "./Layout"; // Updated path
import CurrentUserGame from "./CurrentUserGame"; // Updated path

const Singleplayer = ({ gameState, setGameState }) => {
  return (
    <Layout>
      <div className="CurrentUserGame">
        <CurrentUserGame gameState={gameState} setGameState={setGameState} />
      </div>
    </Layout>
  );
};

export default Singleplayer;
