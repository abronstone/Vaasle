import React from 'react'
import Row from './Row'

interface Letter {
  key: string
}

interface GridProps {
  guesses: Letter[][],
  currentGuess: Letter[],
  turn: number
}

export default function Grid({ guesses, currentGuess, turn }: GridProps): JSX.Element {
  return (
    <div>
      {guesses.map((g: Letter[], i: number) => {
        if (turn === i) {
          return <Row key={i} guess={currentGuess} />;
        } else {
          return <Row key={i} guess={g} />;
        }
      })}
    </div>
  )
}