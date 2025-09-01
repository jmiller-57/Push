import React, { useEffect, useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

export default function Lobby({ token }) {
  const [rooms, setRooms] = useState([]);
  const [error, setError] = useState("");
  const [roomName, setRoomName] = useState("");
  const [creating, setCreating] = useState(false);
  const [selectedRoom, setSelectedRoom] = useState(null);
  const [roomDetails, setRoomDetails] = useState(null);

  const navigate = useNavigate();

  useEffect(() => {
    axios
      .get("http://localhost:8080/api/lobby/rooms/list", {
        headers: { Authorization: "Bearer " + token }
      })
      .then(res => setRooms(Array.isArray(res.data) ? res.data : []))
      .catch(err => setError("Failed to load rooms: " + (err.response?.data || err.message)));
  }, [token, creating]);

  useEffect(() => {
    if (selectedRoom) {
      axios
        .get(`http://localhost:8080/api/lobby/rooms/${selectedRoom}`, {
          headers: { Authorization: "Bearer " + token }
        })
        .then(res => setRoomDetails(res.data))
        .catch(err => setError("Failed to load room details: " + (err.response?.data || err.message)));
    } else {
      setRoomDetails(null);
    }
  }, [selectedRoom, token]);

  // Handle create room
  const handleCreateRoom = async e => {
    e.preventDefault();
    setError("");
    try {
      await axios.post(
        "http://localhost:8080/api/lobby",
        { roomname: roomName },
        { headers: { Authorization: "Bearer " + token } }
      );
      setRoomName("");
      setCreating(!creating);
    } catch (err) {
      setError("Failed to create room: " + (err.response?.data || err.message));
    }
  };

  // Handle joining a room
  const handleJoinRoom = async roomId => {
    setError("");
    try {
      await axios.post(
        "http://localhost:8080/api/lobby/rooms/join",
        { room_id: roomId },
        { headers: { Authorization: "Bearer " + token } }
      );
      navigate(`/lobby/rooms/${roomId}`);
    } catch (err) {
      setError("Failed to join room: " + (err.response?.data || err.message));
    }
  };

  if (error) return <div style={{ color: "red" }}>{error}</div>;
  return (
    <div>
      <h2>Lobby</h2>
      <form onSubmit={handleCreateRoom}>
        <input
          value={roomName}
          onChange={e => setRoomName(e.target.value)}
          placeholder="Room Name"
          required
        />
        <button type="submit">Create Room</button>
      </form>
      {error && <div style={{ color: "red"}}>{error}</div>}
      <ul>
        {rooms.map(room => (
          <li key={room.id}>
            {room.roomname}
            <button onClick={() => handleJoinRoom(room.id)}>Join</button>
            <button onClick={() => setSelectedRoom(room.id)}>Details</button>
          </li>
        ))}
      </ul>
      {roomDetails && (
        <div>
          <h3>Room Details</h3>
          <p>Name: {roomDetails.roomname}</p>
          <p>Created by: {roomDetails.creator}</p>
          <h4>Members:</h4>
          <ul>
            {roomDetails.members.map(member => (
              <li key={member.id}>{member.username}</li>
            ))}
          </ul>
          <button onClick={() => setSelectedRoom(null)}>Close</button>
        </div>
      )}
    </div>
  );
}