import axios from "axios";

/**
 * Calls gateway's POST /newGame endpoint
 * 
 * @param {number} maxGuesses 
 * @param {number} wordLength 
 * @param {string} userId 
 * @returns data object from the POST /newGame response
 */
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

/**
 * Calls gateway's POST /makeGuess endpoint
 * 
 * @param {string} gameId 
 * @param {string} guess 
 * @returns data object from the gateway POST /makeGuess endpoint
 */
export const makeGuessApi = async (gameId, guess) => {
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

/**
 * Calls gateway to create a user with the given userName
 * 
 * @param {string} userName 
 * @param {string} userId 
 * @returns {boolean} true if the user was created, false otherwise
 */

export const createUserApi = async (userName, userId) => {
  try {
    const res = await axios.put("http://localhost:5002/createUser", {
      userName: userName,
      id: userId,
    }, { headers: { "Content-Type": "application/json", } });

    if (res.status !== 200) {
      return false;
    }
    return true;
  } catch (error) {
    return false;
  }
};

/**
 * Calls gateway to (login) validate that the user exists in the database
 * 
 * @param {string} userName 
 * @returns {boolean} true if you can login, false otherwise
 */
export const loginApi = async (userName) => {
  try {
    const res = await axios.put("http://localhost:5002/login/" + userName);
    if (res.status !== 200) {
      return false;
    }
    return true;
  } catch (error) {
    return false;
  }
};

export const getExternalUserGamesApi = async (gameId) => {
  const externalUserIds = ['123456', '109876', '234567', '987654', ];
  // ['345678', '876543', '8525001', '8525002']
  const externalUserGamesMap = new Map();
  const state = 'ongoing';

  // User 123456's guesses
  externalUserGamesMap.set('123456', [
    ['G', 'Y', 'X', 'G', 'Y'],
    ['Y', 'Y', 'X', 'Y', 'Y'],
    ['Y', 'Y', 'X', 'Y', 'Y']
  ]);

  // User 109876's guesses
  externalUserGamesMap.set('109876', [
    ['X', 'Y', 'G', 'Y', 'G'],
    ['Y', 'G', 'Y', 'G', 'Y']
  ]);

  // User 234567's guesses
  externalUserGamesMap.set('234567', [
    ['G', 'G', 'Y', 'X', 'Y'],
    ['Y', 'Y', 'Y', 'G', 'X'],
    ['X', 'X', 'G', 'G', 'Y']
  ]);

  // User 987654's guesses
  externalUserGamesMap.set('987654', [
    ['Y', 'X', 'X', 'Y', 'G'],
    ['G', 'Y', 'Y', 'X', 'Y'],
    ['Y', 'G', 'Y', 'G', 'G']
  ]);

  // // Adding new user 345678's guesses
  // externalUserGamesMap.set('345678', [
  //   ['X', 'G', 'Y', 'Y', 'X'],
  //   ['G', 'X', 'G', 'Y', 'Y'],
  //   ['Y', 'Y', 'X', 'G', 'G']
  // ]);

  // // Adding new user 876543's guesses
  // externalUserGamesMap.set('876543', [
  //   ['Y', 'Y', 'G', 'X', 'Y'],
  //   ['X', 'Y', 'X', 'G', 'Y'],
  //   ['G', 'G', 'Y', 'Y', 'X']
  // ]);

  // externalUserGamesMap.set('8525001', [
  //   ['Y', 'Y', 'G', 'X', 'Y'],
  //   ['X', 'Y', 'X', 'G', 'Y'],
  //   ['G', 'G', 'Y', 'Y', 'X']
  // ])

  // externalUserGamesMap.set('8525002', [
  //   ['Y', 'Y', 'G', 'X', 'Y'],
  //   ['X', 'Y', 'X', 'G', 'Y'],
  //   ['G', 'G', 'Y', 'Y', 'X']
  // ])


  return {
    externalUserIds,
    externalUserGamesMap,
    state,
  }
};
