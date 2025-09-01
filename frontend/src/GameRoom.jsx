import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";

export default function GameRoom({ token }) {
  const { id } = useParams(); // room ID from URL
  const [room, setRoom] = useState(null);
  const [gameState, setGameState] = useState(null);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

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

  // Fetch game state
  const fetchGameState = () => {
    setLoading(true);
    axios
      .get(`http://localhost:8080/api/lobby/rooms/${id}/game`, {
        headers: { Authorization: "Bearer " + token }
      })
      .then(res => {
        setGameState(res.data);
        console.log("Fetched game state: ", res.data);
      })
      .catch(err => { 
          if (err.response && err.response.status === 404) {
            setGameState(null);
          } else {
            setError("Failed to load game state: " + (err.response?.data || err.message))
          }
        })
      .finally(() => setLoading(false));
  };

  useEffect(() => {
    fetchGameState();
  }, [id, token]);

  useEffect(() => {
    if (gameState) {
      console.log("Updated gameState:", gameState);
    }
  }, [gameState]);

  const handleStartGame = async () => {
    setError("");
    setLoading(true);
    try {
      await axios.post(
        `http://localhost:8080/api/lobby/rooms/${id}/start`,
        {},
        { headers: { Authorization: "Bearer " + token } }
      );
      fetchGameState();
    } catch (err) {
      setError("Failed to start game: " + (err.response?.data || err.message));
    } finally {
      setLoading(false);
    }
  };

  if (error) return <div style={{ color: "red" }}>{error}</div>;
  if (!room) return <div>Loading Room...</div>;

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
      {/* Show start game button if at least 2 members and game not started */}
      {!gameState && room.members.length >= 2 && (
        <button onClick={handleStartGame} disabled={loading}>
          {loading ? "Starting..." : "Start Game"}
        </button>
      )}
      {/* Show game state if started */}
      {gameState && (
        <div>
          <h3>Game Started!</h3>
          <ul>
            {gameState.Players.map((player, idx) => (
              <li key={idx}>
                {player.Name} (Score: {player.Score})
                <br/>
                Hand: {player.Hand.map(card => card.Rank + card.Suit).join(", ")}
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
}
