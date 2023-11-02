import axios from "axios";

export const newGameApi = async (maxGuesses, wordLength, userId) => {
  try {
    const res = await axios.post(
      "http://localhost:5002/newGame",
      {
        maxGuesses,
        wordLength,
        userId,
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
    throw new Error("Fetch failed!!", e);
  }
};

export const makeGuessApi = async (gameId, guess) => {
  console.log("makeGuessApi called");
  try {
    // Make the POST request
    const res = await axios.post(
      "http://localhost:5002/makeGuess",
      {
        id: gameId,
        guess: guess,
      },
      {
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    console.log("makeGuessApi response: ", res);
    if (res.status !== 200) {
      throw new Error(`HTTP error! Status: ${res.status}`);
    }

    const data = res.data;
    
    return data;
  } catch (e) {
    console.error("Fetch failed!!", e);
    throw new Error("Fetch failed!!", e);
  }
};

/* 

@param {string} userName
@return {boolean} true if the user was created, false otherwise

Calls gateway to create a user with the given userName
*/

export const createUserApi = async (userName, userId) => {
  console.log("Creating user with username: ", userName);
  try {
    const res = await axios.put("http://localhost:5002/createUser", {
      userName: userName,
      id: userId,
    }, { headers: { "Content-Type": "application/json", } });

    if (res.status !== 200) {
      console.log("Create user failed");
      return false;
    }
    console.log("Create user succeeded");
    return true;
  } catch (error) {
    console.log("Create user failed with error: ", error);
    return false;
  }
};

/*

@param {string} userName
@return {boolean} true if you can login, false otherwise

Calls gateway to (login) validate that the user exists in the database

*/
export const loginApi = async (userName) => {
  console.log("loginApi called");
  try {
    const res = await axios.put("http://localhost:5002/login/" + userName);
    if (res.status !== 200) {
      console.log("Login returned false");
      return false;
    }
    console.log("Login returned true");
    return true;
  } catch (error) {
    console.log("Login failed with error: ", error.message);
    return false;
  }
};