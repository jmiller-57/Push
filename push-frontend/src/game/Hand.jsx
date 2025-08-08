import Card from "./Card";
import { useGame } from "./GameContext";

export default function Hand({ cards, myTurn }) {
  const { send } = useGame();

  function discard(cardId) {
    if (!myTurn) return;
    send("DISCARD", {cardId });
  }

  return (
    <div className="hand">
      {cards.map(c => (
        <Card key={c.id} card={c} onClick={() => discard(c.id)} />
      ))}
    </div>
  );
}