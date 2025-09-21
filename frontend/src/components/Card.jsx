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
  isSelected = false,
  onToggleSelect
}) {
  const ref = useRef(null);
  const isDraggable = !!moveCard;

  const [, drag] = useDrag({
    type: ItemType,
    item: { idx },
    canDrag: isDraggable,
  });

  const [, drop] = useDrop({
    accept: ItemType,
    hover(item) {
      if (moveCard && item.idx !== idx) {
        moveCard(item.idx, idx);
        item.idx = idx;
      }
    },
    canDrop: () => isDraggable,
  });

  if (isDraggable) drag(drop(ref));

  const cardKey = suitRankMap[card.Rank + card.Suit];
  const Card = deck[cardKey]

  const style = {
    marginLeft: idx === 0 ? 0 : `-${CARD_OVERLAP}px`,
    position: "relative",
    flexShrink: 0,
    cursor: isDraggable ? "grab" : "pointer",
    // Keep selected card on top, then hovered, then natural stacking by idx
    zIndex: isSelected ? 1002 : isHovered ? 1001 : idx,
    // Lift the selected card to indicate selection
    transform: isSelected ? "translateY(-24px)" : "translateY(0)",
    transition: "transform 120ms ease, box-shadow 120ms ease, border 120ms ease, z-index 60ms ease",
    // Visual affordances for hover/selected
    boxShadow: isSelected
      ? "0 6px 14px rgba(0,0,0,0.25)"
      : isHovered
      ? "0 4px 16px rgba(0, 123, 255, 0.3)"
      : undefined,
    border: isSelected
      ? "2px solid #28a745"          // green when selected
      : isHovered
      ? "2px solid #007bff"          // blue on hover
      : "2px solid transparent",
    borderRadius: "8px",
  };

  return Card ? (
    <div
      ref={ref}
      style={style}
      onClick={(e) => { e.stopPropagation(); onToggleSelect?.(); }}
      onKeyDown={(e) => {
        if (e.key === "Enter" || e.key === " ") {
          e.preventDefault();
          onToggleSelect?.();
        }
      }}
      role="button"
      tabIndex={0}
      aria-pressed={isSelected}
      onMouseEnter={() => setHoveredIdx(idx)}
      onMouseLeave={() => setHoveredIdx(null)}
    >
      <Card />
    </div>
  ) : (
    <span>Unknown Card...</span>
  );
}