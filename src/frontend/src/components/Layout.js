import React from "react";
import { Link } from "react-router-dom";
import { useAuth0 } from "@auth0/auth0-react";

const Layout = ({ children }) => {
  const { isAuthenticated } = useAuth0(); 

  return (
    <div>
      <div className="headerContainer">
        <Link to="/">
          <h1>Vaasle</h1>
        </Link>
        {isAuthenticated && (
          <Link to="/stats" className="statsButton">
            Stats
          </Link>
        )}
      </div>
      {children}
    </div>
  );
};

export default Layout;
