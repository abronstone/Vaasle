import React from "react";

export default function ExternalUserRow({ corrections }) {  
  if (corrections) {
    return (
      <div className="row past">
        {corrections.map((correction, i) => {
          return (
            <div key={i} className={correction} />
          );
        })}
      </div>
    );
  }

  // If there are no corrections to be made, render empty squares
  return (
    <div className="row">
      <div></div>
      <div></div>
      <div></div>
      <div></div>
      <div></div>
    </div>
  );
}
