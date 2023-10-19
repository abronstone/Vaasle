import axios from "axios";

export const newGameApi = async (maxGuesses, wordLength) => {
  try {
    const res = await axios.post("http://play-game:5001/newGame", {
      headers: {
        "Content-Type": "application/json",
      },
      body: {
        "maxGuesses": maxGuesses,
        "wordLength" : wordLength
      }
    });

    if (res.status !== 200) {
      throw new Error(`HTTP error! Status: ${res.status}`);
    }

    const data = await res.data;
    return data;
  } catch (e) {
    throw new Error("Fetch failed!!");
  }
};

export const makeGuessApi = async (gameId, guess) => {
  try {
    const res = await axios.post("http://play-game:5001/makeGuess", {
      headers : {
        "Content-Type": "application/json",
      },
      body: {
        "id": gameId,
        "guess": guess
      }
    });

    if (res.status !== 200) {
      throw new Error(`HTTP error! Status: ${res.status}`);
    }

    const data = await res.data;
    return data;
  } catch (e) {
    throw new Error("Fetch failed!!");
  }
};