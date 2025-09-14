import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";
import * as deck from "@letele/playing-cards";
import { suitRankMap, CARD_HEIGHT, CARD_WIDTH, CARD_OVERLAP } from "./cardUtils";

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
                  <div 
                    style={{ 
                      display: "flex",
                      alignItems: "center",
                      marginTop: "8px",
                      height: `${CARD_HEIGHT}px`,
                      position: "relative"
                    }}
                  >
                    {player.Hand.map((card, idx) => {
                      const cardKey = suitRankMap[card.Rank + card.Suit];
                      const CardComponent = deck[cardKey];
                      return CardComponent ? (
                        <div
                          key={idx}
                          style={{
                            marginLeft: idx === 0 ? 0 : `-${CARD_OVERLAP}px`,
                            zIndex: idx, // ensures cards stack in order
                            position: "relative",
                            flexShrink: 0,
                          }}
                        >
                          <CardComponent  />
                        </div>
                      ) : (
                        <span key={idx}>Unknown Card...</span>
                      );
                    })}
                  </div>
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
              {gameState.FaceUpCard && (() => {
                const cardKey = suitRankMap[gameState.FaceUpCard.Rank + gameState.FaceUpCard.Suit];
                const CardComponent = deck[cardKey];
                return CardComponent ? (
                  <CardComponent />
                ) : (
                  <span>Unknown Card...</span>
                );
              })()}
              {gameState.DeckCount > 0 && (
                <div
                  style={{
                    position: "relative",
                    width: `${CARD_WIDTH}px`,
                    height: `${CARD_HEIGHT}px`,
                    marginLeft: "24px",
                  }}
                >
                  {Array.from({ length: Math.min(gameState.DeckCount, 10) }).map((_, i) => {
                    const BackCard = deck["B2"];
                    return BackCard ? (
                      <div
                        key={i}
                        style={{
                          position: "absolute",
                          top: `${(10 - Math.min(gameState.DeckCount, 10) + i) * 4}px`,
                          left: 0,
                          zIndex: i,
                          opacity: 1 - (0.05 * (Math.min(gameState.DeckCount, 10) - i - 1)),
                        }}
                      >
                        <BackCard />
                      </div>
                    ) : null;
                  })}
                </div>
              )}
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
// TODO: Fix other players being hidden behind "my" cards