import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";

export default function GameRoom({ token }) {
  const { id } = useParams(); // room ID from URL
  const [room, setRoom] = useState(null);
  const [error, setError] = useState("");

  useEffect(() => {
    axios
      .get(`http://localhost:8080/api/lobby/rooms/${id}`, {
        headers: { Authorization: "Bearer " + token }
      })
      .then(res => setRoom(res.data))
      .catch(err =>
        setError("Failed to load room: " + (err.response?.data || err.message))
      );
  }, [id, token]);

  if (error) return <div style={{ color: "red" }}>{error}</div>;
  if (!room) return <div>Loading...</div>;

  return (
    <div>
      <h2>Game Room: {room.name}</h2>
      <p>Creator: {room.creator}</p>
      <h3>Members:</h3>
      <ul>
        {room.members.map(member => (
          <li key={member.id}>{member.username}</li>
        ))}
      </ul>
      {/* Add game UI here later */}
    </div>
  );
}
