const formatLetters = (letters) => {
  const abbreviatedColorToFullColorMap = new Map([
    ["G", "green"],
    ["Y", "yellow"],
    ["X", "grey"],
  ])

  const formattedDictionary = {}

  letters.forEach((letter) => {
    formattedDictionary[letter[0]] = abbreviatedColorToFullColorMap.get(
      letter[1])
    
  })

  return formattedDictionary
}

export default formatLetters