import { useState } from "react";
import { playGameApi } from "./util/apiCalls";
import "./App.css";

function App() {
  const [playGameApiData, setPlayGameApiData] = useState("No response");

  const handleEngineButtonClick = async () => {
    try {
      const data = await playGameApi();
      setPlayGameApiData(data);
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <>
      <h1 className="mb-3">Welcome to Wordle! (Thin Thread)</h1>
      <button
        className="bg-slate-500 mt-3 mb-3"
        onClick={handleEngineButtonClick}
      >
        Click to call the play-game container
      </button>
      <h3 className="mb-1">Data from the play-game container:</h3>
      <p>{playGameApiData}</p>
    </>
  );
}

export default App;
