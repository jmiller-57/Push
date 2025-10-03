import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";
import { useDrag, useDrop } from "react-dnd";
import PlayerHand from "./components/PlayerHand";
import DeckStack from "./components/DeckStack";
import FaceUpCard from "./components/FaceUpCard";
import OpponentSeat from "./components/OpponentSeat";
import { CARD_WIDTH, CARD_HEIGHT } from "./cardUtils";

export default function GameRoom({ token }) {
  const { id } = useParams(); // room ID from URL
  const [room, setRoom] = useState(null);
  const [gameState, setGameState] = useState(null);
  const [gameStarted, setGameStarted] = useState(false);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const [hoveredCardIdx, setHoveredCardIdx] = useState(null);
  const [hand, setHand] = useState([]);
  const [selectedIds, setSelectedIds] = useState(new Set());

  const tableStyle = {
    position: "relative",
    width: "min(1100px, 92vw)",
    height: "72vh",
    margin: "24px auto",          // center on the page
    borderRadius: 12,
    background:
      "radial-gradient(circle at 50% 40%, #1e7f52 0%, #0f5f3e 50%, #0b3f2a 100%)",
    boxShadow: "inset 0 0 80px rgba(0,0,0,0.35)",
    overflow: "visible",
    zIndex: 1,
  };
  const centerAreaStyle = {
    position: "absolute",
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
    display: "flex",
    alignItems: "flex-end",
    justifyContent: "center",
    gap: 12,
  };
  const topOpponentsStyle = {
    height: `${CARD_HEIGHT / 2}px`,
    position: "relative",
    top: 24,
    display: "flex",
    justifyContent: "center",
    gap: 32,
    alignItems: "center",
    zIndex: 5,
  };
  const bottomHandStyle = {
    position: "fixed",
    left: "50%",
    bottom: 16,
    transform: "translateX(-50%)",
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    gap: 8,
    zIndex: 10,
  };
  const actionsStyle = { display: "flex", gap: 8 };

  const CENTER_SCALE = 0.5;

  // Clear card selections when the hand changes
  useEffect(() => {
    setSelectedIds(new Set());
  }, [hand]);

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

  const toggleSelectById = (cardId) => {
    setSelectedIds(prev => {
      const next = new Set(prev);
      next.has(cardId) ? next.delete(cardId) : next.add(cardId);
      return next;
    });
  };

  const handlePlaySelected = async () => {
    if (selectedIds.size === 0) return;

    const cardIds = Array.from(selectedIds);
    try {
      await axios.post(
        `http://localhost:8080/api/lobby/rooms/${id}/play`,
        { cardIds },
        { headers: {Authorization: "Bearer " + token } }
      );

      setHand(prev => prev.filter(c => !selectedIds.has(c.ID || c.id)));
      setSelectedIds(new Set());
    } catch (err) {
      setError("Failed to play selected cards: " + (err.response?.data || err.message));
    }
  };

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

  function ScaleBox({ id, scale = 0.5, children }) {
    return (
      <div
        id={`${id}`}
        style={{
          width: CARD_WIDTH * scale,
          height: CARD_HEIGHT * scale,
          position: "relative",
        }}
      >
        <div
          style={{
            position: "absolute",
            top: 0,
            left: 0,
            width: CARD_WIDTH,
            height: CARD_HEIGHT,
            transform: `scale(${scale})`,
            transformOrigin: "top left",
          }}
        >
          {children}
        </div>
      </div>
    );
  }

  if (error) return <div style={{ color: "red" }}>{error}</div>;
  if (!room) return <div>Loading Room...</div>;

  const players = gameState?.Players ?? [];
  const me = players.find(p => p.Hand);
  const opponents = players.filter(p => !p.Hand);

  return (
    <div id="gameroom" onClick={() => setSelectedIds(new Set())}>
      <div id="header">
        <h1 style={{ display: "inline", textAlign: "left" }}>Game Room: {room.name}</h1>
      </div>
      {/* Show start game button if at least 2 members and game not started */}
      {!gameStarted && room.members.length >= 2 && (
        <button onClick={handleStartGame} disabled={loading}>
          {loading ? "Starting..." : "Start Game"}
        </button>
      )}
      {/* Show game state if started */}
      {gameStarted && gameState && (
        <div id="tablespace" onClick={() => setSelectedIds(new Set())}>
          <div id="table" style={tableStyle}>
            <div id="opponent" style={topOpponentsStyle}>
              {opponents.map((p, i) => (
                <OpponentSeat key={i} name={p.Name} handCount={p.HandCount} />
              ))}
            </div>

            <div id="centerarea" style={centerAreaStyle}>
              <ScaleBox id={"faceupcard"} scale={CENTER_SCALE}>
                <FaceUpCard card={gameState.FaceUpCard} />
              </ScaleBox>
              
              {gameState.DeckCount > 0 && (
                <ScaleBox id={"deckstack"} scale={CENTER_SCALE}>
                  <DeckStack deckCount={gameState.DeckCount} />
                </ScaleBox>
              )}
            </div>
          </div>

          <div id="playerhand" style={bottomHandStyle} onClick={(e) => e.stopPropagation()}>
            <div 
              id="hand"
              style={{
                transformOrigin: "bottom center",
              }}
            >
              <PlayerHand
                hand={hand}
                setHand={setHand}
                hoveredIdx={hoveredCardIdx}
                setHoveredIdx={setHoveredCardIdx}
                selectedIds={selectedIds}
                onToggleSelect={toggleSelectById}
              />
            </div>
            {selectedIds.size > 0 && (
              <div id="actionbuttons" style={{ display: "flex", gap: 8 }}>
                <button onClick={() => setSelectedIds(new Set())}>Deselect All</button>
                <button onClick={handlePlaySelected}>Play Selected ({selectedIds.size})</button>
              </div>
            )}
          </div>
        </div>
      )}
    </div>
  );
}