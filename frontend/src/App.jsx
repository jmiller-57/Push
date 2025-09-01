import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { useState } from "react"
import AuthForm from "./AuthForm";
import GameRoom from "./GameRoom";
import Lobby from "./Lobby";
import './App.css';

function App() {
  const [token, setToken] = useState(localStorage.getItem("jwt") || "");

  const handleAuth = t => {
    setToken(t);
    localStorage.setItem("jwt", t);
  };

  if (!token || token === "") return <AuthForm onAuth={handleAuth} />;
  return (
    <Router>
      <Routes>
        <Route path='/lobby/rooms' element={<Lobby token={token} />} />
        <Route path='/lobby/rooms/:id' element={<GameRoom token={token} />} />
        <Route path='*' element={<Lobby token={token} />} />
      </Routes>
    </Router>
  );
}
export default App;
