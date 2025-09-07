import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";
import * as deck from "@letele/playing-cards";

const suitRankMap = {
  "A♠": "Sa",
  "2♠": "S2",
  "3♠": "S3",
  "4♠": "S4",
  "5♠": "S5",
  "6♠": "S6",
  "7♠": "S7",
  "8♠": "S8",
  "9♠": "S9",
  "10♠": "S10",
  "J♠": "Sj",
  "Q♠": "Sq",
  "K♠": "Sk",
  "A♥": "Ha",
  "2♥": "H2",
  "3♥": "H3",
  "4♥": "H4",
  "5♥": "H5",
  "6♥": "H6",
  "7♥": "H7",
  "8♥": "H8",
  "9♥": "H9",
  "10♥": "H10",
  "J♥": "Hj",
  "Q♥": "Hq",
  "K♥": "Hk",
  "A♣": "Ca",
  "2♣": "C2",
  "3♣": "C3",
  "4♣": "C4",
  "5♣": "C5",
  "6♣": "C6",
  "7♣": "C7",
  "8♣": "C8",
  "9♣": "C9",
  "10♣": "C10",
  "J♣": "Cj",
  "Q♣": "Cq",
  "K♣": "Ck",
  "A♦": "Da",
  "2♦": "D2",
  "3♦": "D3",
  "4♦": "D4",
  "5♦": "D5",
  "6♦": "D6",
  "7♦": "D7",
  "8♦": "D8",
  "9♦": "D9",
  "10♦": "D10",
  "J♦": "Dj",
  "Q♦": "Dq",
  "K♦": "Dk",
  "Joker*": "J1"
}

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
              <div key={idx}>
                {player.Name} (Cards: {player.Count})
                {player.Hand && (
                  <div 
                    style={{ 
                      display: "flex",
                      alignItems: "center",
                      marginTop: "8px",
                      height: "100px",
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
                            marginLeft: idx === 0 ? 0 : -24, // overlap cards by 40px
                            zIndex: idx, // ensures cards stack in order
                            position: "relative",
                            width: "60px",
                            height: "90px",
                          }}
                        >
                          <CardComponent />
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
          <div>
            <br/>
            <br/>
            <br/>
            <br/>
            <br/>
            <br/>
            <br/>
            {/* TODO: play with alignment of cards */}
          </div>
          <div>
            <div style={{ 
              display: "flex",
              justifyContent: "left",
              marginTop: "32px",
              alignItems: "left",
              position: "relative",
              zIndex: 1000
            }}>
              <span style={{ marginRight: "12px"}}>Face up card:</span>
              {gameState.FaceUpCard && (() => {
                const cardKey = suitRankMap[gameState.FaceUpCard.Rank + gameState.FaceUpCard.Suit];
                const CardComponent = deck[cardKey];
                return CardComponent ? (
                  <CardComponent/>
                ) : (
                  <span>Unknown Card...</span>
                );
              })()}
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
