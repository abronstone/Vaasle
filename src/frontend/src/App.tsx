import { useEffect, useState } fro"react";
import Wordle from "./components/Wordle";
// import "./index.css";

interface Solution {
  id: number;
  word: string;
}

function App(): JSX.Element {
  const [solution, setSolution] = useState<string | null>(null);

  useEffect(() => {
    fetch("http://localhost:3001/solutions")
      .then((res) => res.json())
      .then((json: Array<Solution>) => {
        // random int between 0 & 14
        const randomSolution = json[Math.floor(Math.random() * json.length)];
        // get the word property of the solution object
        setSolution(randomSolution.word);
      });
  }, [setSolution]);

  return (
    <div className="App">{solution && <Wordle solution={solution} />}</div>
  );
}

export default App;
