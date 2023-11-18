import React from "react";
import { Link } from "react-router-dom";
import Stats from "./Stats"; // Import Stats component
import { useAuth0 } from "@auth0/auth0-react"; // Import useAuth0 hook

const Layout = ({ children }) => {
  const { isAuthenticated, user } = useAuth0(); // useAuth0 inside the component

  return (
    <div>
      <div className="headerContainer">
        <Link to="/">
          <h1>Vaasle</h1>
        </Link>
        {isAuthenticated && user && (
          <Link to="/stats" className="statsButton">
            Stats
          </Link>
        )}
      </div>
      {/* Render child components */}
      {children}
      {/* Render the Stats component for authenticated users */}
      {isAuthenticated && user && <Stats />}
    </div>
  );
};

export default Layout;
