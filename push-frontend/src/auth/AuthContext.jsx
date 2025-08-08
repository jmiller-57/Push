import { createContext, useContext, useState } from "react";
import { createHttp } from "../api/http"

const AuthCtx = createContext(null);

export function AuthProvider({ children }) {
  const [token, setToken] = useState(null);
  const http = createHttp(() => token);

  async function login(email, password) {
    const { access_token } = await http.post("/login", { email, password });
    setToken(access_token);
  }

  async function register(email, password) {
    const { access_token } = await http.post("/register", { email, password });
    setToken(access_token);
  }

  function logout() {
    setToken(null);
  }

  return (
    <AuthCtx.Provider value={{ token, login, register, logout, http }}>
      {children}
    </AuthCtx.Provider>
  );
}

export function useAuth() { return useContext(AuthCtx); }