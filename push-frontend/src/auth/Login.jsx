import { useState } from "react";
import { useAuth } from "./AuthContext";

export default function Login({ onSuccess }) {
  const { login } = useAuth();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [err, setErr] = useState("");

  async function submit(e) {
    e.preventDefault();
    try { await login(email, password); onSuccess?.(); }
    catch (e) { setErr(e.message); }
  }

  return (
    <form onSubmit={submit} className="auth-form">
      <h2>Sign in</h2>
      {err && <div className="error">{err}</div>}
      <input value={email} onChange={e => setEmail(e.target.value)} placeholder="Email" />
      <input value={password} onChange={e => setPassword(e.target.value)} placeholder="Password" />
      <button>Login</button>
    </form>
  );
}