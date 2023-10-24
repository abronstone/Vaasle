    import React from 'react'

    export default function Modal({ isCorrect, solution, turn }) {
    return (
    <div className="modal">
        {isCorrect && (
        <div>
            <h1>You Win!</h1>
            <p>The word was: </p>
            {solution && <p className="solution">{solution}</p>}
            <p>You found the solution in {turn} guess {turn > 1 ? "es" : ""} :)</p>
        </div>
        )}
        {!isCorrect && (
        <div>
            <h1>You lost</h1>
            <p>The word was: </p>
            {solution && <p className="solution">{solution}</p>}
            <p>Better luck next time :)</p>
        </div>
        )}
    </div>
    )
    }