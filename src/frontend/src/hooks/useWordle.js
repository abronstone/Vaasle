import { useState } from 'react'

const useWordle = (solution) => {
  // track user turn
  const [turn, setTurn] = useState(0)

  // track what user is guessing
  const [currentGuess, setCurrentGuess] = useState('')

  // each guess is an array, having length of array be 6
  const [guesses, setGuesses] = useState([...Array(6)])

  // history of guesses of their
  const [history, setHistory] = useState([])

  // track if the guess is correct
  const [isCorrect, setIsCorrect] = useState(false)

  // coloring used keys, values are color keys are the letter
  const [usedKeys, setUsedKeys] = useState({}) // {a: 'green', b: 'yellow, c: 'green'}

  // format a guess into an array of letter objects along with coloring them
  // e.g. [{key: 'a', color: 'green'}]
  const formatGuess = () => {
    let solutionArray = [...solution]
    let formattedGuess = [...currentGuess].map((l) => {
      return {key: l, color: 'grey'}  // each object/letter defaults to grey
    })

    // find green letters
    formattedGuess.forEach((l, i) => {
      if (solution[i] === l.key) {
        formattedGuess[i].color = 'green'
        solutionArray[i] = null
      }
    })

    // find any yellow letters
    formattedGuess.forEach((l, i) => {
      if (solutionArray.includes(l.key) && l.color !== 'green') {
        formattedGuess[i].color = 'yellow'
        solutionArray[solutionArray.indexOf(l.key)] = null
      }
    })

    return formattedGuess
  }

  // add a new guess to the guesses state
  // update the isCorrect state if the guess is correct
  // add one to the turn state
  const addNewGuess = (formattedGuess) => {
    if (currentGuess === solution) {
      setIsCorrect(true)
    }
    setGuesses(prevGuesses => {
      let newGuesses = [...prevGuesses]
      newGuesses[turn] = formattedGuess
      return newGuesses
    })
    setHistory(prevHistory => {
      return [...prevHistory, currentGuess]
    })
    setTurn(prevTurn => {
      return prevTurn + 1
    })
    setUsedKeys((prevUsedKeys) => {
      let newkeys = {...prevUsedKeys}
      formattedGuess.forEach((l) => {
        // if it exists, get color else its empty
        const currentColor = newkeys[l.key]
        // assign the letter to be color green
        if (l.color === 'green'){
          newkeys[l.key] = 'green'
          return
        }
        // only update color if found it was yellow, and it already doesnt have a green property
        if (l.color === 'yellow' && currentColor !== 'green'){
          newkeys[l.key] = 'yellow'
          return
        }
        if (l.color === 'grey' && currentColor !== 'green' && currentColor !== 'yellow'){
          newkeys[l.key] = 'grey'
          return
        }
      })
      return newkeys
    })
    setCurrentGuess('')
  }

  // handle keyup event & track current guess
  // if user presses enter, add the new guess
  // using regex to ensure it's an English character and not CAPS LOCK, shift, etc.
  // hard coding with 5
  const handleKeyup = ({ key }) => {
    if (key === 'Enter') {
      // only add guess if turn is less than 6
      if (turn >= 6) {
        console.log('You used all your guesses!')
        return
      }
      // do not allow duplicate words
      if (history.includes(currentGuess)) {
        console.log('You already tried that word.')
        return
      }
      // ensure word is 5 characters long
      if (currentGuess.length !== 5) {
        console.log('Word must be 5 characters long.')
        return
      }
      const formatted = formatGuess()
      addNewGuess(formatted)
    }
    if (key === 'Backspace') {
      setCurrentGuess(prev => prev.slice(0, -1))
      return
    }
    if (/^[A-Za-z]$/.test(key)) {
      if (currentGuess.length < 5) {
        setCurrentGuess(prev => prev + key)
      }
    }
  }

  return {turn, currentGuess, guesses, isCorrect, handleKeyup, usedKeys}
}

export default useWordle;
