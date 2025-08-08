import { useEffect, useState } from "react";
import { useAuth } from "../auth/AuthContext";
import CreateRoom from "./CreateRoom"
import RoomList from "./RoomList"

export default function Lobby({ onEnterRoom }) {
  const { http } = useAuth();
  const [rooms, setRooms] = useState([]);

  async function load() {
    const data = await http.get("/rooms");
    setRooms(data.rooms ?? []);
  }
  useEffect(() => {load(); const id=setInterval(load, 3000); return () => clearInterval(id); }, []);

  async function createRoom(maxPlayers = 6) {
    const room = await http.post("/rooms", { maxPlayers });
    onEnterRoom(room.id);
  }

  async function joinRoom(id) {
    await http.post(`/rooms/${id}/join`, {});
    onEnterRoom(id);
  }

  return (
    <div className="lobby">
      <CreateRoom onCreate={createRoom}/>
      <RoomList rooms={rooms} onJoin={joinRoom}/>
    </div>
  );
}