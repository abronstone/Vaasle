const useWordle = (gameState) => {
  if(gameState == null || gameState.guesses == null) return {turn: 0, guesses: [], isCorrect: false, usedKeys: []}
  
  const turn = gameState.guesses.length

  const abbreviatedColorToFullColorMap = new Map([
    ["G", "green"],
    ["Y", "yellow"],
    ["X", "grey"],
  ])
  

  // const guesses = gameState.guesses.reduce((guessesDictionary, guess) => {
  //   for(let i = 0; i < guess.length; i++) {
  //     guessesDictionary[guess[i][0]] = abbreviatedColorToFullColorMap.get(guess[i][1])
  //   }
  // }, {})

  // transform guesses into map
  const guessesMap = new Map()

  gameState.guesses.forEach((guess) => {
    guess.forEach((letter) => {
      guessesMap.set(letter[0], abbreviatedColorToFullColorMap.get(letter[1]))
    })
  })

  // transform map into an object
  const guesses = Object.fromEntries(guessesMap)

  const isCorrect = gameState.state === "won"

  // make a map of all the unique used letter color pairs from guesses
  const usedKeysMap = new Map()
  
  Object.keys(guesses).forEach((key) => {
    usedKeysMap.set(key, guesses[key])
  })

  const usedKeys = Object.fromEntries(usedKeysMap)
  
  console.log({turn, guesses, isCorrect, usedKeys})
  return {turn, guesses, isCorrect, usedKeys}
}

export default useWordle




