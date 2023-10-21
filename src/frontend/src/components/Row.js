import React from "react";

export default function Row({ guesses, currentGuess }) {
  console.log("guesses inside row", guesses);
  // Convert the 'guesses' object to an array of objects.
  const guessesArray = guesses
    ? Object.keys(guesses).map((key) => {
        return { color: guesses[key], key: key };
      })
    : null;

  console.log("guessesArray", guessesArray);

  if (guessesArray) {
    return (
      <div className="row past">
        {guessesArray.map((l, i) => {
          console.log("l.key", l.key);
          console.log("l.color", l.color);
          return (
            <div key={i} className={l.color}>
              {l.key}
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
