import React, { useEffect, useState } from 'react'
import useWordle from '../hooks/useWordle'

// components
import Grid from './Grid'
import Keypad from './Keypad'
import Modal from './Modal'

export default function Wordle({ solution }) {
  const { currentGuess, guesses, turn, isCorrect, handleKeyup, usedKeys } = useWordle(solution)
  const[showModal, setShowModal] = useState(false)
  // attach keyup listening to event object
  useEffect(() => {
    window.addEventListener('keyup', handleKeyup)

    if (isCorrect){
      console.log("you won the game!")
      setTimeout(() => setShowModal(true),2000)
      window.removeEventListener('keyup', handleKeyup)
    }

    if (turn > 5 && !isCorrect) {
      console.log("unlucky, you ran out of guesses")
      setTimeout(() => setShowModal(true),2000)
      window.removeEventListener('keyup', handleKeyup)
    }

    return () => window.removeEventListener('keyup', handleKeyup)
  }, [handleKeyup, isCorrect, turn])


  return (
    <div>
      {/* <div>solution - {solution}</div> */}
      <div>Current Guess - {currentGuess}</div>
      <Grid guesses={guesses} currentGuess={currentGuess} turn={turn} />
      <Keypad usedKeys = {usedKeys}/>
      {showModal && <Modal isCorrect = {isCorrect} turn = {turn} solution = {solution}/>}
    </div>
  )
}