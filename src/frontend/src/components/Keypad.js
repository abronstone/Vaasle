    import React, { useEffect, useState } from 'react'
    // ! do we have keyboard on front end or in database

    export default function Keypad({usedKeys}) {
    const [letters, setLetters] = useState(null)
    useEffect(() => {
        fetch('http://localhost:3001/letters')
        .then(res => res.json())
        .then(json => {
            setLetters(json)
        })
    },[])
    return (
    <div className='keypad'>
        {/* letters are from the db file */}
        {/* wait till we have letters then map thru them */}
        {letters && letters.map((l) => {
            const color = usedKeys[l.key]
            return (
                // l.key => letter of keyboard
                <div key={l.key} className={color} >{l.key}</div>
            )
        })}
    </div>
    )
    }
