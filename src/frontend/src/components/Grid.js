    import React from 'react'
    // components
    import Row from './Row'

    export default function Grid({ guesses, currentGuess, turn }) {
    console.log("guesses", guesses)
    console.log("type of guesses", guesses)
    return (
    <div>
    {/* mapping thru the guesses and output a row for each guess,
    guesses currently have length of 6 */}
        {guesses !== {} && guesses != null && guesses.map((g, i) => {
            // when idx matches turn vaue, return this row with the current guess
            if (turn === i) {
                return <Row key = {i} currentGuess = {currentGuess}/>
            }
        
        return <Row key={i} guess={g} /> 
        })}
    </div>
    )
    }