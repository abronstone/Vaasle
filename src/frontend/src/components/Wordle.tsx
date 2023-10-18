import React, { useEffect, useState } from 'react'
import useWordle from '../hooks/useWordle'

// components
import Grid from './Grid'
import Keypad from './Keypad'
import Modal from './Modal'

interface Letter {
  key: string
}

interface WordleProps {
  solution: string;
}

export default function Wordle({ solution }: WordleProps): JSX.Element {
  const [currentGuess, setCurrentGuess] = useState<Letter[]>([])
  const { guesses, turn, isCorrect, usedKeys, handleKeyup } = useWordle(solution)

  useEffect(() => {
    setCurrentGuess(Array.from({ length: solution.length }, () => ({ key: '' })))
  }, [solution])

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setCurrentGuess(event.target.value.split('').map(l => ({ key: l })))
  }

  return (
    <div className="wordle">
      <h1>Wordle</h1>
      <Grid guesses={guesses as any} currentGuess={currentGuess} turn={turn} />
      <Keypad usedKeys={usedKeys} />
      <Modal isCorrect={isCorrect} solution={solution} turn={turn} />
      <input type="text" value={currentGuess.map(l => l.key).join('')} onChange={handleInputChange} onKeyUp={handleKeyup as any} />
    </div>
  )
}