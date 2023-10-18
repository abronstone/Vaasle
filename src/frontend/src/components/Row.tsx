import React from "react";

interface Props {
  guess?: string;
  currentGuess?: string;
}

export default function Row({ guess, currentGuess }: Props): JSX.Element {
  const letters = guess ? guess.split("") : new Array(4).fill("");

  return (
    <div>
      {letters.map((l, i) => (
        <div key={i} className="cell">
          {currentGuess && i === currentGuess.length ? (
            <input type="text" maxLength={1} value={currentGuess[i]} />
          ) : (
            <div>{l}</div>
          )}
        </div>
      ))}
    </div>
  );
}
