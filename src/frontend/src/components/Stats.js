import React, { useState, useEffect } from "react";
import { useAuth0 } from "@auth0/auth0-react";
import { Link } from "react-router-dom";
import { getStatsApi } from "./util/apiCalls";

export default function Stats() {
  const [statsData, setStatsData] = useState({
    gamesPlayed: 0,
    winPercentage: 0,
    mostGuessedWord: "",
  });
  const { isAuthenticated, user } = useAuth0();

  useEffect(() => {
    // ? not sure if we're only using user or isAuth as well.
    if (user) {
      // get the user from their UUID, take the data then setStatsSata
      getStatsApi(user.sub)
        .then((data) => {
          setStatsData(data);
        })
        .catch((error) => {
          console.error("ERROR FETCHING STATS", error);
        });
    }
  }, [user]);

  return (
    <div className="stats-container">
      <h2>Statistics</h2>
      <ul>
        <li>Games Played: {statsData.gamesPlayed}</li>
        <li>Win %: {statsData.winPercentage}</li>
        <li>Most Guessed Word: {statsData.mostGuessedWord}</li>
      </ul>
      <Link to="/" className="close-btn">
        X
      </Link>
    </div>
  );
}
