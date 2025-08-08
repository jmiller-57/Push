import { useGame } from "./GameContext";
import Hand from "./Hand";
import DrawPile from "./DrawPile";
import TurnBar from "./TurnBar";

export default function GameTable() {
  const { state } = useGame();
  if (!state) return <div>Loading game...</div>

  const me = state.players?.find(p => p.isMe);
  const myTurn = state.currentPlayerId === me?.id;

  return (
    <div className="table">
      <div className="piles">
        <DrawPile count={state.drawPileCount}/>
      </div>
      
      <div className="players">
        {state.players?.map(p => (
          <div key={p.id} className={`player ${p.id===me?.id?"me":""}`}>
            <div>{p.name} {state.currentPlayerId===p.id ? "⭐" : ""}</div>
            <div>Cards: {p.handCount}</div>
          </div>
        ))}
      </div>

      <Hand card={me?.hand ?? []} myTurn={myTurn}/>
      <TurnBar myTurn={myTurn}/>
    </div>
  );
}