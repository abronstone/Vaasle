import { makeGuessApi } from "../util/apiCalls"  
  
  // handle keyup event & track current guess
  // if user presses enter, add the new guess

  // TODO: Maybe pass in the currentGame state as instead of breaking it up
  // could override this in function
const handleKeyup = (key, turn, isCorrect, currentGuess, guesses, setCurrentGuess, gameId) => {
    if (key === 'Enter') {
      // only add guess if turn is less than 5
      if (turn > 5 && !isCorrect) {
        console.log('you used all your guesses!')
        return
      }
      // do not allow duplicate words
      if (guesses.includes(currentGuess)) {
        console.log('you already tried that word.')
        return
      }
      // check word is 5 chars
      if (currentGuess.length !== 5) {
        console.log('word must be 5 chars.')
        return
      }

      // make guess
      const getNewGameState = async () => {
        const gameState = await makeGuessApi(gameId, currentGuess)
        return gameState
      }

      const newGameState  = getNewGameState().then((res) => {
        return res
      })

      const abbreviatedColorToFullColorMap = new Map([
        ["G", "green"],
        ["Y", "yellow"],
        ["X", "grey"],
      ])
    
      const lastGuessDictionary = {}
    
      const newGuesses = newGameState.guesses
      
      newGuesses[newGuesses.length - 1].forEach((letter) => {
          lastGuessDictionary[letter[0]] = abbreviatedColorToFullColorMap.get(
            letter[1]
          )
        })
    
      const formatted = Object.keys(lastGuessDictionary).map((key) => {
        return { key: key, color: lastGuessDictionary[key] }
      })
    
      // addNewGuess(formatted)
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

  export default handleKeyup
