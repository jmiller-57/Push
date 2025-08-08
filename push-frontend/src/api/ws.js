export function createGameSocket({ gameId, getToken, onMessage }) {
  let ws;
  let retry = 0;
  const WS_BASE = (import.meta.env.VITE_WS_BASE ?? "ws://localhost:8080/ws");

  function connect() {
    const token = getToken?.();
    const url = `${WS_BASE}/games/${gameId}?token=${encodeURIComponent(token || "")}`;
    ws = new WebSocket(url);

    ws.onoopen = () => { retry = 0; };
    ws.onclose = () => {
      if (retry < 8) {
        setTimeout(connect, 500 * 2 ** retry);
        retry++;
      }
    };

    ws.onmessage = (evt) => {
      try { onMessage?.(JSON.parse(evt.data)); } catch {}
    };
  }

  connect();

  return {
    send: (type, payload) => {
      const msg = JSON.stringify({ type, payload });
      ws?.readyState === 1 && ws.send(msg);
    },
    close: () => ws?.close(),
  };
}