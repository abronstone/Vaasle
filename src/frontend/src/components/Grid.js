    import React from "react";
    // components
    import Row from "./Row";

    export default function Grid({ guesses, currentGuess, turn }) {
    return (
    <div>
        {Object.keys(guesses).map((key, i) => {
        if (turn === i) {
            return <Row key={i} currentGuess={currentGuess} />;
        }
        return <Row key={i} guess={guesses[key]} />;
        })}
    </div>
    );
    }
