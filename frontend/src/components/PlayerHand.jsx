import React from "react";
import Card from "./Card";
import { CARD_HEIGHT } from "../cardUtils";

export default function PlayerHand({ 
  hand,
  setHand,
  hoveredIdx,
  setHoveredIdx,
  selectedIds,
  onToggleSelect,
}) {

  const moveCard = (from, to) => {
    const updated = [...hand];
    const [moved] = updated.splice(from, 1);
    updated.splice(to, 0, moved);
    setHand(updated);
  };

  const getCardId = (c) => c.ID ?? c.id;

  return (
    <div
      onClick={(e) => e.stopPropagation()}
      style={{
        display: "flex",
        alignItems: "center",
        marginTop: "8px",
        height: `${CARD_HEIGHT}px`,
        position: "relative",
      }}
    >
      {hand.map((card, idx) => {
        const cardId = getCardId(card);
        const isSelected = selectedIds?.has(cardId);
        return ( 
          <Card
            key={cardId || idx}
            card={card}
            idx={idx}
            moveCard={moveCard}
            isHovered={hoveredIdx === idx}
            setHoveredIdx={setHoveredIdx}
            isSelected={!!isSelected}
            onToggleSelect={() => onToggleSelect(cardId)}
          />
        );
      })}
    </div>
  );
}