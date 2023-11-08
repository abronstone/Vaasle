import React from "react";

export default function CurrentUserRow({ guesses, currentGuess }) {  
  if (guesses) {
    return (
      <div className="row past">
        {guesses.map((l, i) => {
          return (
            <div key={i} className={l.color}>
              {l.letter}
            </div>
          );
        })}
      </div>
    );
  }

  if (currentGuess) {
    let letters = currentGuess.split("");
    return (
      <div className="row current">
        {letters.map((letter, i) => (
          <div key={i} className="filled">
            {letter}
          </div>
        ))}

        {[...Array(5 - letters.length)].map((_, i) => (
          <div key={i}></div>
        ))}
      </div>
    );
  }

  // If currentGuess or guesses don't exist, then just render empty squares
  return (
    <div className="row">
      <div></div>
      <div></div>
      <div></div>
      <div></div>
      <div></div>
    </div>
  );
}
