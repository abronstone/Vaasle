import React from 'react'
import { Link } from "react-router-dom";

export default function Modal({ isCorrect, solution, turn }) {
return (
<div className="modal">
    {isCorrect && (
    <div>
        <h1>You Win!</h1>
        <p>The word was: </p>
        {solution && <p className="solution">{solution}</p>}
        <p>You found the solution in {turn + 1} guess{turn + 1 > 1 ? "es" : ""} :)</p>
        <Link to="/" className="general-rounded-button">
            Return Home
        </Link>
    </div>
    )}
    {!isCorrect && (
    <div>
        <h1>You lost</h1>
        <p>The word was: </p>
        {solution && <p className="solution">{solution}</p>}
        <p>Better luck next time :)</p>
        <Link to="/" className="general-rounded-button">
            Return Home
        </Link>
    </div>
    )}
</div>
)
}