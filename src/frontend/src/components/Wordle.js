import React, { useEffect, useState } from 'react'
import useWordle from './gameLogic/useWordle'
import handleKeyup from './gameLogic/handleKeyUp'

// components
import Grid from './Grid'
import Keypad from './Keypad'
import Modal from './Modal'

export default function Wordle({gameState}) {
  const [currentGuess, setCurrentGuess] = useState('')
  const { guesses, turn, isCorrect, usedKeys } = useWordle(gameState)
  const[showModal, setShowModal] = useState(false)
  
  // attach keyup listening to event object
  useEffect(() => {
    const keyUpEventListener = (e) => {
      handleKeyup(e.key, turn, isCorrect, currentGuess, guesses, setCurrentGuess, gameState.metadata.id)
    }

    window.addEventListener('keyup', keyUpEventListener)

    if (isCorrect){
      console.log("you won the game!")
      setTimeout(() => setShowModal(true),2000)
      window.removeEventListener('keyup', keyUpEventListener)
    }

    if (turn > 5 && !isCorrect) {
      console.log("unlucky, you ran out of guesses")
      setTimeout(() => setShowModal(true),2000)
      window.removeEventListener('keyup', keyUpEventListener)
    }

    return () => window.removeEventListener('keyup', keyUpEventListener)
  }, [handleKeyup, isCorrect, turn])


  return (
    <div>
      {/* <div>solution - {solution}</div> */}
      <div>Current Guess - {currentGuess}</div>
      <Grid guesses={guesses} currentGuess={currentGuess} turn={turn} />
      <Keypad usedKeys = {usedKeys}/>
      {showModal && <Modal isCorrect = {isCorrect} turn = {turn} />}
    </div>
  )
}