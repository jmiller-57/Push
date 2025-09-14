import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";
import { useDrag, useDrop } from "react-dnd";
import PlayerHand from "./components/PlayerHand";
import DeckStack from "./components/DeckStack";
import FaceUpCard from "./components/FaceUpCard";
import { CARD_HEIGHT } from "./cardUtils";

export default function GameRoom({ token }) {
  const { id } = useParams(); // room ID from URL
  const [room, setRoom] = useState(null);
  const [gameState, setGameState] = useState(null);
  const [gameStarted, setGameStarted] = useState(false);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const [hoverdCardIdx, setHoveredCardIdx] = useState(null);
  const [hand, setHand] = useState([]);

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

  useEffect(() => {
    if (gameState && gameState.Players) {
      const me = gameState.Players.find(p => p.Hand);
      if (me) setHand(me.Hand);
    }
  }, [gameState]);

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
          <div key={member.id}>{member.username}</div>
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
              <div key={idx} style={{ marginBottom: "32px" }}>
                {player.Name} (Cards: {player.HandCount})
                {player.Hand && (
                  <PlayerHand 
                    hand={hand}
                    setHand={setHand}
                    hoveredIdx={hoverdCardIdx}
                    setHoveredIdx={setHoveredCardIdx}
                  />
                )}
              </div>
            ))}
          </ul>
          <div style={{ marginTop: `${CARD_HEIGHT}px` }}>
            <div style={{ 
              display: "flex",
              justifyContent: "left",
              alignItems: "left",
              position: "relative",
              zIndex: 1000
            }}>
              <span style={{ marginRight: "12px"}}>Face up card:</span>
              <FaceUpCard card={gameState.FaceUpCard} />
              {gameState.DeckCount > 0 && (
                <DeckStack deckCount={gameState.DeckCount} />
              )}
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
// TODO: Fix other players being hidden behind "my" cards