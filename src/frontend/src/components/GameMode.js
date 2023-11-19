import { useAuth0 } from "@auth0/auth0-react";
import React from "react";
import { Link } from "react-router-dom";
import LogoutButton from "./LogoutButton";
import LoginButton from "./LoginButton";
import ErrorBadge from "./ErrorBadge";
import Layout from "./Layout";

export default function GameMode() {
  const { isAuthenticated } = useAuth0();
  return (
    <>
      <Layout>
        {isAuthenticated != null && !isAuthenticated && (
          <ErrorBadge text={"You must be logged into to play"} />
        )}
        <div className="gameMode">
          {isAuthenticated != null && isAuthenticated && (
            <>
              <h2>Select Game Mode</h2>
              <div>
                <Link to="/singleplayer" className="general-rounded-button">
                  Single Player
                </Link>
                <Link to="/multiplayersetup" className="general-rounded-button">
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
      </Layout>
    </>
  );
}
