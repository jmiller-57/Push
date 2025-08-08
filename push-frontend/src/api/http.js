const API_BASE = import.meta.env.VITE_API_BASE ?? "http://localhost:8080/api";

export function createHttp(getToken) {
  async function request(path, opts = {}) {
    const headers = new Headers(opts.headers || {});
    const token = getToken?.();

    if (token) headers.set("Authorization", `Bearer ${token}`);
    headers.set("Content-Type", "application/json");

    const res = await fetch(`${API_BASE}${path}`, { ...opts, headers });
    if (!res.ok) {
      const text = await res.text().catch(() => "");
      throw new Error(text || res.statusText);
    }

    if (res.status === 204) return null;
    return res.json();
  }

  return {
    post: (p, body) => request(p, { method: "POST", body: JSON.stringify(body) }),
    get:  (p, body) => request(p),
  };
}