import { useState } from "react";
import Login from "../auth/Login";
import Lobby from "../lobby/Lobby";
import { GameProvider } from "../game/GameContext";
import GameTable from "../game/GameTable";
import { useAuth } from "../auth/AuthContext";

export default function Routher() {
  const { token } = useAuth();
  const [roomId, setRoomId] = useState(null);

  if (!token) return <Login onSuccess={() => {}} />;
  if (!roomId) return <Lobby onEnterRoom={setRoomId} />;

  return (
    <GameProvider gameId={roomId}>
      <GameTable />
    </GameProvider>
  );
}