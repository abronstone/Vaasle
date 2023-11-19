import React from 'react'
import { Link } from "react-router-dom";

export default function MultiplayerModal({ solution, winner }) {
  return (
    <div className="modal">
      <div>
        <h1>{winner} Won!</h1>
        <p>The word was: </p>
        {solution && <p className="solution">{solution}</p>}
        <Link to="/" className="home-nav-button">
            Return Home
        </Link>
      </div>
    </div>
  )
}