import React from 'react'

export default function MultiplayerModal({ solution, winner }) {
  return (
    <div className="modal">
      <div>
        <h1>{winner} Won!</h1>
        <p>The word was: </p>
        {solution && <p className="solution">{solution}</p>}
      </div>
    </div>
  )
}