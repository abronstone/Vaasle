import React, { useEffect, useState } from 'react'
import useWordle from './gameLogic/useWordle'
import handleKeyup from './gameLogic/handleKeyUp'

// components
import Grid from './Grid'
import Keypad from './Keypad'
import Modal from './Modal'

export default function Wordle({gameState , setGameState}) {
  const [currentGuess, setCurrentGuess] = useState('')
  const { guesses, turn, isCorrect, usedKeys } = useWordle(gameState)
  const[showModal, setShowModal] = useState(false)  

  // attach keyup listening to event object
  useEffect(() => {
    if (!gameState) {
      return; // Exit if gameState is null
    }
    
    const keyUpEventListener = (e) => {
      console.log("gameState before handleKeyup", gameState)
      console.log("currentGuess before handleKeyup", currentGuess)      
      handleKeyup(e.key, turn, isCorrect, currentGuess, guesses, setCurrentGuess, gameState.metadata.gameID, setGameState);
      console.log("gameState after handleKeyup", gameState)
      console.log("currentGuess after handleKeyup", currentGuess)
    };

    window.addEventListener('keyup', keyUpEventListener);
  
    const handleGameEnd = () => {
      console.log(isCorrect ? "you won the game!" : "unlucky, you ran out of guesses");
      setTimeout(() => setShowModal(true), 2000);
      window.removeEventListener('keyup', keyUpEventListener);
    };
  
    if (isCorrect || (turn > 5 && !isCorrect)) {
      handleGameEnd();
    }
  
    return () => window.removeEventListener('keyup', keyUpEventListener);
  }, [handleKeyup, turn, gameState, currentGuess, guesses, isCorrect]);
  
  // 
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