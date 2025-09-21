import React from "react";
import OpponentHand from "./OpponentHand";
import { CARD_HEIGHT } from "../cardUtils";

export default function OpponentSeat({ name, handCount }) {
  return (
      <OpponentHand count={handCount} />
  );
}