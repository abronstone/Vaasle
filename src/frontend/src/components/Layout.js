import React from "react";
import { Link } from "react-router-dom";
import { useAuth0 } from "@auth0/auth0-react";
import logo from "./second.png"; // with import

const Layout = ({ children }) => {
  const { isAuthenticated } = useAuth0();

  return (
    <div>
      <div className="headerContainer">
        <Link to="/">
          <img src={logo} alt="Vaasle spelled out in Wordle style blocks"/>
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
