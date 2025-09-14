import React, { useRef } from "react";
import * as deck from "@letele/playing-cards";
import { useDrag, useDrop } from "react-dnd";
import { suitRankMap, CARD_OVERLAP, CARD_WIDTH, CARD_HEIGHT } from "../cardUtils";

const ItemType = "CARD";

export default function Card({ 
  card,
  idx,
  moveCard,
  isHovered,
  setHoveredIdx,
  style={},
}) {
  const ref = useRef(null);

  const [, drag] = useDrag({
    type: ItemType,
    item: { idx },
    canDrag: !!moveCard,
  });

  const [, drop] = useDrop({
    accept: ItemType,
    hover(item) {
      if (moveCard && item.idx !== idx) {
        moveCard(item.idx, idx);
        item.idx = idx;
      }
    },
    canDrop: () => !!moveCard,
  });

  if (moveCard) drag(drop(ref));

  const cardKey = suitRankMap[card.Rank + card.Suit];
  const Card = deck[cardKey]

  return Card ? (
    <div
      ref={ref}
      style={{
        marginLeft: idx === 0 ? 0 : `-${CARD_OVERLAP}px`,
        zIndex: isHovered ? 100 : idx,
        position: "relative",
        flexShrink: 0,
        transition: "z-index 0.1s, box-shadow 0.1s, border 0.1s",
        boxShadow: isHovered ? "0 4px 16px rgba(0, 123, 255, 0.3)" : undefined,
        border: isHovered ? "2px solid #007bff" : "2px solid transparent",
        borderRadius: "8px",
        cursor: moveCard ? "grab" : "default",
        ...style,
      }}
      onMouseEnter={() => setHoveredIdx(idx)}
      onMouseLeave={() => setHoveredIdx(null)}
    >
      <Card />
    </div>
  ) : (
    <span>Unknown Card...</span>
  );
}