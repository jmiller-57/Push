import React, { useState } from "react";
import axios from "axios"

export default function AuthForm( { onAuth }) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [isLogin, setIsLogin] = useState(false);
  const [error, setError] = useState("");

  const handleSubmit = async e => {
    e.preventDefault();
    setError("");
    try {
      if (isLogin) {
        const res = await axios.post("http://localhost:8080/api/login", { username, password });
        if (res.data.token) {
          onAuth(res.data.token)
        } else {
          setError("No token received");
        }
      } else {
        await axios.post("http://localhost:8080/api/users", { username, password });
        const res = await axios.post("http://localhost:8080/api/login", { username, password });
        if (res.data.token) {
          onAuth(res.data.token)
        } else {
          setError("No token received after registration");
        }
      }
    } catch (err) {
      setError("Auth failed: " + (err.response?.data || err.message));
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>{isLogin ? "Login" : "Register"}</h2>
      <input
        placeholder="Username"
        value={username}
        onChange={e => setUsername(e.target.value)}
        required
      />
      <input
        placeholder="Password"
        type="password"
        value={password}
        onChange={e => setPassword(e.target.value)}
        required
      />
      <button type="submit">{isLogin ? "Login" : "Register"}</button>
      <button type="button" onClick={() => setIsLogin(!isLogin)}>
        {isLogin ? "Switch to Register" : "Switch to Login"}
      </button>
      {error && <div style={{ color: "red" }}>{error}</div>}
    </form>
  );
}