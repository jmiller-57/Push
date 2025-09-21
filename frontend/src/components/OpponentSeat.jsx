import React from "react";
import OpponentHand from "../OpponenetHand";

export default function OpponentSeat({ name, handCount }) {
  return (
    <div style={{ display: "flex", flexDirection: "column", alignItems: "center", color: "white" }}>
      <div style={{ marginBottom: 6, fontWeight: 600 }}>{name}</div>
      <OpponentHand count={handCount} />
    </div>
  );
}