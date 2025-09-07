import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";

export default function GameRoom({ token }) {
  const { id } = useParams(); // room ID from URL
  const [room, setRoom] = useState(null);
  const [gameState, setGameState] = useState(null);
  const [gameStarted, setGameStarted] = useState(false);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  // Fetch room details
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

  useEffect(() => {
    fetchGameState();
    // eslint-disable-next-line
  }, [id, token]);

  // Fetch game state on mount and after starting
  const fetchGameState = () => {
    setLoading(true);
    axios
      .get(`http://localhost:8080/api/lobby/rooms/${id}/game`, {
        headers: { Authorization: "Bearer " + token }
      })
      .then(res => {
        if (res.data.gameStarted === false) {
          setGameStarted(false);
          setGameState(null);
        } else {
          setGameStarted(true);
          setGameState(res.data);
        }        
      })
      .catch(err => { 
          setError("Failed to load game state: " + (err.response?.data || err.message))
        })
      .finally(() => setLoading(false));
  };

  const handleStartGame = async () => {
    setError("");
    setLoading(true);
    try {
      const res = await axios.post(
        `http://localhost:8080/api/lobby/rooms/${id}/start`,
        {},
        { headers: { Authorization: "Bearer " + token } }
      );
      setGameStarted(true);
      setGameState(res.data);
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
      {!gameStarted && room.members.length >= 2 && (
        <button onClick={handleStartGame} disabled={loading}>
          {loading ? "Starting..." : "Start Game"}
        </button>
      )}
      {/* Show game state if started */}
      {gameStarted && gameState && (
        <div>
          <h3>Game Started!</h3>
          <ul>
            {gameState.Players.map((player, idx) => (
              <li key={idx}>
                {player.Name} (Cards: {player.Count})
                {player.Hand && (
                  <>
                    <br />
                    My Hand: {player.Hand.map(card => card.Rank + card.Suit).join(", ")}
                  </>
                )}
              </li>
            ))}
          </ul>
          <div>
            Face up card: {gameState.FaceUpCard.Rank}{gameState.FaceUpCard.Suit}
          </div>
        </div>
      )}
    </div>
  );
}
