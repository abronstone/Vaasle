const useWordle = (gameState) => {
  if(!gameState) return {turn: 0, guesses: [], isCorrect: false, usedKeys: []}
  
  const turn = gameState.guesses.length

  const guesses = gameState.guesses.reduce((newArr, guess) => {
    newArr.push(guess[0])
    return newArr
  }, [])

  const isCorrect = gameState.state === "won"

  const abbreviatedColorToFullColorMap = new Map([
    ["G", "green"],
    ["Y", "yellow"],
    ["X", "grey"],
  ])

  const usedKeysDictionary = {}

  guesses.forEach((guess) => {
    guess.forEach((letter) => {
      usedKeysDictionary[letter[0]] = abbreviatedColorToFullColorMap.get(
        letter[1]
      )
    })
  })

  const usedKeys = Object.keys(usedKeysDictionary).map((key) => {
    return { key: key, color: usedKeysDictionary[key] }
  })

  return {turn, guesses, isCorrect, usedKeys}
}

export default useWordle




