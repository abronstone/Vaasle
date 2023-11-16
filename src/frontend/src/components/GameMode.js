import { useAuth0 } from "@auth0/auth0-react";
import React from "react";
import { Link } from "react-router-dom";
import LogoutButton from "./LogoutButton";
import LoginButton from "./LoginButton";
import ErrorBadge from "./ErrorBadge";

export default function GameMode() {
  const { isAuthenticated } = useAuth0();
  return (
    <>
      <Link to="/">
        <h1>Vaasle</h1>
      </Link>
      {isAuthenticated != null && !isAuthenticated && (
        <ErrorBadge text={"You must be logged into to play"} />
      )}
      <div className="gameMode">
        {isAuthenticated != null && isAuthenticated && (
          <>
            <h2>Select Game Mode</h2>
            <div>
              <Link to="/singleplayer" className="gameMode-button">
                Single Player
              </Link>
              <Link to="/multiplayersetup" className="gameMode-button">
                Multiplayer
              </Link>
            </div>
          </>
        )}
      </div>
      <div className="auth0button-container">
        {isAuthenticated != null && isAuthenticated ? (
          <LogoutButton />
        ) : (
          <LoginButton />
        )}
      </div>
    </>
  );
}
