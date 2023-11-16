import React from "react";
import { Link } from "react-router-dom";

const Layout = ({ children }) => {
  return (
    <div>
      <div className="headerContainer">
        <Link to="/">
          <h1>Vaasle</h1>
        </Link>
        <Link to="/stats" className="statsButton">
          Stats
        </Link>
      </div>
      {children}
    </div>
  );
};

export default Layout;
