import React from "react";
import * as deck from "@letele/playing-cards"
import { CARD_HEIGHT, CARD_OVERLAP, CARD_WIDTH } from "../cardUtils";

export default function OpponentHand({ count, scale = 0.5 }) {
  const BackCard = deck["B2"];
  
  return (
    <div
      id="opponent_hand"
      style={{
        display: "flex",
        alignItems: "center",
        position: "relative",
        transform: `scale(${scale})`,
      }}
    >
      {Array.from({ length: count }).map((_, idx) => (
        <div
          id={`card_${idx}`}
          key={idx}
          style={{
            marginLeft: idx === 0 ? 0 : `-${CARD_OVERLAP}px`,
            position: "relative",
            zIndex: idx,
          }}
        >
          <BackCard />
        </div>
      ))}
    </div>
  );
}