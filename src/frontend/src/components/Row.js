    import React from 'react'
    // guess from Grids.js, either previous formatted guess or a undefined guess.

    export default function Row({ guess, currentGuess }) {

    // take the guess and cycle thru the letter objects if it exists
    if (guess) {
    return (
        <div className="row past">
    {/* map thru prior guesses and go to each character */}
        {guess.map((l, i) => (
            <div key={i} className={l.color}>{l.key}</div>
        ))}
        </div>
    )
    }

    // need to accomodate 
    // 5 - letters
    if (currentGuess) {
        // splits each letter into an array
        let letters = currentGuess.split('')
        return (
            // mapping thru letters of current guess
            <div className="row current">
                {letters.map((letter, i) => (
                    <div key={i} className='filled'>{letter}</div>
                ))}

                {/* putting out the empty squares as we type based off the letter len */}
                {[...Array(5-letters.length)].map((_,i) => (
                    <div key={i}></div>
                ))}
            </div>
            
        )
    }

    return (
    // make 5 dics per row for each square
    <div className="row">
        <div></div>
        <div></div>
        <div></div>
        <div></div>
        <div></div>
    </div>
    )

    }