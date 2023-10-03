import { useState } from 'react'
import { engineApi } from './util/apiCalls'
import './App.css'

function App() {
  const [engineApiData, setEngineApiData] = useState("No response")

  const handleEngineButtonClick = async () => {
    try{
      const data = await engineApi()
      setEngineApiData(data)
    }
    catch(error){
      console.log(error)
    }
  }

  return (
    <>
      <h1>Welcome to Wordle! (Thin Thread)</h1>
      <button className="bg-slate-500"onClick={handleEngineButtonClick}>
        Click to call the engine container
      </button>
      <h3>Data from the engine container:</h3>
      <p>
        {engineApiData}
      </p>
    </>
  )
}

export default App
