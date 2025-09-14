import React from "react";
import Card from "./Card";

export default function FaceUpCard({ card }) {
  if (!card) return null;
  return (
    <div style={{ marginLeft: "12px" }}>
      <Card card={card} idx={0} isHovered={false} setHoveredIdx={() => {}} />
    </div>
  );
}