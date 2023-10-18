import React, { useEffect, useState } from "react";

interface Letter {
  id: number;
  key: string;
}

interface Props {
  usedKeys: { [key: string]: boolean };
}

export default function Keypad({ usedKeys }: Props): JSX.Element {
  const [letters, setLetters] = useState<Array<Letter> | null>(null);

  useEffect(() => {
    fetch("http://localhost:3001/letters")
      .then((res) => res.json())
      .then((json) => {
        setLetters(json);
      });
  }, []);

  return (
    <div className="keypad">
      {letters &&
        letters.map((l) => {
          const color = usedKeys[l.key] ? "gray" : "white";
          return (
            <button key={l.id} style={{ backgroundColor: color }}>
              {l.key}
            </button>
          );
        })}
    </div>
  );
}
