import React, { useEffect, useState } from 'react'

interface Letter {
    key: string
}

interface KeypadProps {
    usedKeys: Record<string, string>
}

export default function Keypad({ usedKeys }: KeypadProps): JSX.Element {
    const [letters, setLetters] = useState<Letter[] | null>(null)

    useEffect(() => {
        fetch('http://localhost:3001/letters')
            .then(res => res.json())
            .then((json: Letter[]) => {
                setLetters(json)
            })
    }, [])

    return (
        <div className='keypad'>
            {/* letters are from the db file */}
            {/* wait till we have letters then map thru them */}
            {letters && letters.map((l: Letter) => {
                const color = usedKeys[l.key]
                return (
                    // l.key => letter of keyboard
                    <div key={l.key} className={color} >{l.key}</div>
                )
            })}
        </div>
    )
}