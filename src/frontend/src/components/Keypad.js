import React from 'react'
export default function Keypad({ usedKeys }) {
    const letters = ["a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"];
    return (
        <div className='keypad'>
            {letters != null && letters.map((letter) => {
                const color = usedKeys.has(letter) ? usedKeys.get(letter) : 'unused'
                return (
                    <div key={letter} className={color} >{letter}</div>
                )
            })}
        </div>
    )
}
