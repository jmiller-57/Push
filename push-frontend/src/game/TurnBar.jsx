import { useGame } from "./GameContext"

export default function TurnBar({ myTurn }) {
  const { send, state } = useGame();
  if (!myTurn) return <div className="turnbar">Waiting for other players...</div>;

  return (
    <div className="turnbar">
      <button onClick={() => send("TAKE_FACEUP", {})} disabled={!state.faceUpCard}>Take Card</button>
      <button onClick={() => send("PUSH_FACEUP", {})} disabled={!state.faceUpCard}>Push Card</button>
      <button onClick={() => send("TRY_GO_DOWN", {})}>Go Down</button>
    </div>
  )
}