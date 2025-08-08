import { createContex, useContext, useEffect, useReducer, useRef } from "react";
import { createGameSocket } from "../api/ws";
import { useAuth } from "../auth/AuthContext";

const GameCtx = createContex(null);

const initial = {
  gameId: null,
  state: null,
  connected: false,
  error: null
};

function reducer (s, a) {
  switch(a.type) {
    case "CONNECT": return { ...s, gameId: a.gameId, connected: true, error: null };
    case "STATE": return { ...s, state: a.payloard };
    case "ERROR": return { ...s, error: a.message };
    case "DISCONNECT": return initial;
    default: return s;
  }
}

export function GameProvider({ gameId, children }) {
  const [state, dispatch] = useReducer(reducer, { ...initial, gameId });
  const { token } = useAuth();
  const socketRef = useRef(null);

  useEffect(() => {
    if (!gameId || !token) return;
    dispatch({ type: "CONNECT", gameId });

    socketRef.current = createGameSocket({
      gameId,
      getToken: () => token,
      onMessage: (msg) => {
        if (msg.type === "STATE") dispatch({ type: "STATE", payload: msg.payload });
        if (msg.type === "ERROR") dispatch({ type: "ERROR", message: msg.payload?.message || "Error" });
      }
    });
    return () => socketRef.current?.close();
  }, [gameId, token]);

  function send(type, payload) { socketRef.current?.send(type, payload); }

  return (
    <GameCtx.Provider value={{ ...state, send }}>
      {children}
    </GameCtx.Provider>
  );
}

export function useGame() { return useContext(GameCtx); }