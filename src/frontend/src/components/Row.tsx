import React from 'react'

interface Letter {
  key: string
}

interface RowProps {
  guess?: Letter[] | undefined,
  currentGuess?: string
}

export default function Row({ guess, currentGuess }: RowProps): JSX.Element {
  // take the guess and cycle thru the letter objects if it exists
  if (guess) {
    return (
      <div className="row past">
        {/* map thru prior guesses and go to each character */}
        {guess.map((l: Letter, i: number) => (
          <div key={i} className="cell">{l.key}</div>
        ))}
      </div>
    )
  } else {
    return (
      <div className="row">
        {/* map thru current guess and go to each character */}
        {currentGuess?.split('').map((l: string, i: number) => (
          <div key={i} className="cell">{l}</div>
        ))}
      </div>
    )
  }
}