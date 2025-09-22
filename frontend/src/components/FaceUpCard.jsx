import React from "react";
import Card from "./Card";
import { CARD_HEIGHT, CARD_WIDTH } from "../cardUtils";

export default function FaceUpCard({ card }) {
  if (!card) return null;
  return (
    <Card card={card} idx={0} isHovered={false} setHoveredIdx={() => {}} />
  );
}