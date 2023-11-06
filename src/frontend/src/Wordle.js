import React, { useEffect, useState } from "react";
import { makeGuessApi, createUserApi, loginApi } from "./components/util/apiCalls";
import LoginButton from "./components/LoginButton";
import LogoutButton from "./components/LogoutButton";
import { useAuth0 } from "@auth0/auth0-react";
import ErrorBadge from "./components/ErrorBadge";

// components
import Grid from "./components/Grid";
import Keypad from "./components/Keypad";
import Modal from "./components/Modal";

export default function Wordle({ gameState, setGameState }) {
  const { isAuthenticated, user } = useAuth0()
  const [loginSuccessful, setLoginSuccessful] = useState(false)

  const [state, setState] = useState({
    currentGuess: "",
    guesses: [],
    turn: 0,
    usedKeys: new Map(),
    status: "ongoing", // "ongoing", "won", or "lost"
  });

  // These state variables are only used at the end of the game, 
  // so they are not part of the main state object
  const [showModal, setShowModal] = useState(false);
  const [isCorrect, setIsCorrect] = useState(false);
  const [solution, setSolution] = useState(null);
  const [error, setError] = useState(null);

  const handleGameEnd = () => {
    setShowModal(true);

    window.removeEventListener("keyup", handleKeyup);
  };

  // TODO: break this up into smaller functions
  // Each time a key is pressed, the handleKeyup function is called
  const handleKeyup = async (e) => {
    // If the game is over, do not allow any more guesses
    if (showModal) {
      return;
    }

    const key = e.key;

    // If the enter key is pressed, the current guess is submitted and the FE's state is updated
    if (key === "Enter") {
      // Make previous state variables easy to work with
      const { turn, currentGuess, usedKeys } = state;

      // check word is 5 chars
      if (currentGuess.length !== 5) {
        setError("Your guess must be 5 characters long.");
        return;
      }
      try {
        // When a guess is submitted, the API called to get a new game state
        // and all necessary state is updated

        // When a guess is submitted, call API to get a new game state
        const newGameState = await makeGuessApi(
          gameState.metadata.gameID,
          currentGuess
        );

        // Update various state variables based on the newGameState
        setGameState(newGameState);

        if(newGameState == null || newGameState.guesses == null) {
          setError("Your guess must be a valid english word. No duplicates are allowed.");
          return;
        }

        // If the game is over, show the modal and stop listening for keyup events
        if (newGameState.state === "won" || newGameState.state === "lost") {
          if (newGameState.word) {
            setSolution(newGameState.word);
          }
          setIsCorrect(newGameState.state === "won" ? true : false);
          setState({ ...state, status: newGameState.state });

          handleGameEnd();
          return;
        }

        // Create an array of mappings of letters to colors for the most recent guess.
        const mostRecentGuessArr = [];

        // Deconstruct the guess into the word and its color codes (eg. "GGYXG").
        if (newGameState.guesses.length > 0) {
          const [word, colorCodes] = newGameState.guesses[turn];

          Array.from(word).forEach((letter, index) => {
            const colorCode = colorCodes[index];

            // Determine the color needed for FE based on the color code.
            const color =
              colorCode === "G"
                ? "green"
                : colorCode === "Y"
                  ? "yellow"
                  : "grey";

            // Update the mapping of the letter to the color.
            mostRecentGuessArr.push({ letter, color });
          });
        }

        // Make a map of all the unique used letter color pairs from guesses
        const newUsedKeys = usedKeys;

        // Loop through the mostRecentGuessArr to add to the map
        mostRecentGuessArr.forEach(({ letter, color }) => {
          // If the letter is not already in the map, add it.
          if (usedKeys.has(letter) === false) {
            newUsedKeys.set(letter, color);
          }
          if (usedKeys.has(letter) === true && color === "green") {
            newUsedKeys.set(letter, color);
          }
        });

        setState((prevState) => ({
          ...prevState,
          turn: prevState.turn + 1,
          guesses: [...prevState.guesses, mostRecentGuessArr],
          status: newGameState.state,
          usedKeys: newUsedKeys,
          currentGuess: ""
        }));

        setError(null);

      } catch (error) {
        console.log(error);
        setError("Your guess must be a valid english word. No duplicates are allowed.");
        return; 
      }

    }
    if (key === "Backspace") {
      setState({ ...state, currentGuess: state.currentGuess.slice(0, -1) });
      return;
    }
    if (/^[A-Za-z]$/.test(key)) {
      if (state.currentGuess.length < 5) {
        setState({ ...state, currentGuess: state.currentGuess + key });
      }
    }
  };

  // attach keyup listening to event object
  useEffect(() => {
    window.addEventListener("keyup", handleKeyup);

    return () => window.removeEventListener("keyup", handleKeyup);
  }, [state, showModal]);

  useEffect(() => {
    const fetchData = async () => {
      if (isAuthenticated != null && isAuthenticated) {

        const login = await loginApi(user.sub);
        setLoginSuccessful(login);
        console.log("Login successful: ", login)

        if (loginSuccessful === false) {
          console.log("Creating user with id and username: ", user.sub, user.name)
          const createUser = await createUserApi(user.name, user.sub);
          console.log("Create user successful: ", createUser)
        }

        setError(null);
      } else {
        setError("You must be logged in to play.")
      }
    };

    fetchData();
  }, [isAuthenticated]); 

  return (
    <div className="wordle-container">
      {error != null && <ErrorBadge text={error} />}
      {isAuthenticated != null && isAuthenticated && (
        <>
          <Grid
            guesses={state.guesses}
            currentGuess={state.currentGuess}
            turn={state.turn}
            status={state.status}
          />
          <Keypad usedKeys={state.usedKeys} />
          {showModal && (
            <Modal
              isCorrect={isCorrect}
              turn={state.turn}
              solution={solution}
            />
          )}
        </>
      )}
      {isAuthenticated != null && isAuthenticated ? (<LogoutButton />) : (<LoginButton />)}
    </div>
  );
}
