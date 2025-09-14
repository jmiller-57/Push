import React from "react";
import Card from "./Card";
import { CARD_HEIGHT } from "../cardUtils";

export default function PlayerHand({ hand, setHand, hoveredIdx, setHoveredIdx}) {
  const moveCard = (from, to) => {
    const updated = [...hand];
    const [moved] = updated.splice(from, 1);
    updated.splice(to, 0, moved);
    setHand(updated);
  };

  return (
    <div
      style={{
        display: "flex",
        alignItems: "center",
        marginTop: "8px",
        height: `${CARD_HEIGHT}px`,
        position: "relative",
      }}
    >
      {hand.map((card, idx) => (
        <Card
          key={idx}
          card={card}
          idx={idx}
          moveCard={moveCard}
          isHovered={hoveredIdx === idx}
          setHoveredIdx={setHoveredIdx}
        />
      ))}
    </div>
  );
}