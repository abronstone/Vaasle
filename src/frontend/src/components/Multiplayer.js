import React from "react";
import Layout from "./Layout";
import CurrentUserGame from "./CurrentUserGame";
import ExternalUserGame from "./ExternalUserGame";

const Multiplayer = ({
  gameState,
  setGameState,
  externalGamesState,
  isMultiplayerEnabled,
  isAuthenticated,
}) => {
  return (
    <Layout>
      <div className="App">
        <div className="CurrentUserGame">
          <CurrentUserGame gameState={gameState} setGameState={setGameState} />
        </div>
        {isMultiplayerEnabled &&
          isAuthenticated &&
          externalGamesState != null &&
          Array.from(externalGamesState.externalUserGamesMap.values()).map(
            (game, index) => (
              <div
                className={`ExternalUserGame ${
                  index % 2 === 0 ? "odd" : "even"
                }`}
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
