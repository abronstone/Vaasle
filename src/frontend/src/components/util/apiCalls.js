import axios from "axios";

export const newGameApi = async (maxGuesses, wordLength) => {
  try {
    // Make the POST request
    const res = await axios.post(
      "http://localhost:5002/newGame",
      {
        maxGuesses,
        wordLength,
      },
      {
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    // Check for successful status code
    if (res.status !== 200) {
      throw new Error(`HTTP error! Status: ${res.status}`);
    }

    const data = res.data;

    return data;
  } catch (e) {
    console.error("Fetch failed!!", e);
    throw new Error("Fetch failed!!");
  }
};

export const makeGuessApi = async (gameId, guess) => {
  try {
    // Make the POST request
    const res = await axios.post(
      "http://localhost:5002/makeGuess",
      {
        // This is the request body
        id: gameId,
        guess: guess,
      },
      {
        // These are the request headers
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    // Check for successful status code
    if (res.status !== 200) {
      throw new Error(`HTTP error! Status: ${res.status}`);
    }

    // Since axios automatically parses the JSON response,
    // you can directly get the data without using await
    const data = res.data;

    return data;
  } catch (e) {
    // Better to log the actual error along with a custom message
    console.error("Fetch failed!!", e);
    throw new Error("Fetch failed!!");
  }
};
