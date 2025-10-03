import React from "react";
import * as deck from "@letele/playing-cards";
import { CARD_WIDTH, CARD_HEIGHT } from "../cardUtils";

export default function DeckStack({ deckCount }) {
  const BackCard = deck["B2"];
  const visibleCount = Math.min(deckCount, 10);
  const STEP = 4;

  return (
    <div
      style={{
        position: "relative",
        width: `${CARD_WIDTH}px`,
        height: `${CARD_HEIGHT}px`,
        overflow: "visible",
      }}
    >
      {Array.from({ length: visibleCount }).map((_, i) => (
        <div
          key={i}
          style={{
            position: "absolute",
            bottom: `${i * STEP}px`,
            left: 0,
            zIndex: i,
            lineHeight: 0
          }}
        >
          <BackCard />
        </div>
      ))}
    </div>
  );
}