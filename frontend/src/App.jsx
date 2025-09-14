import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { useState } from "react"
import AuthForm from "./AuthForm";
import GameRoom from "./GameRoom";
import Lobby from "./Lobby";
import './App.css';
import { DndProvider } from 'react-dnd';
import { HTML5Backend } from 'react-dnd-html5-backend';

function App() {
  const [token, setToken] = useState(localStorage.getItem("jwt") || "");

  const handleAuth = t => {
    setToken(t);
    localStorage.setItem("jwt", t);
  };

  if (!token || token === "") return <AuthForm onAuth={handleAuth} />;
  return (
    <DndProvider backend={HTML5Backend}>  
      <Router>
        <Routes>
          <Route path='/lobby/rooms' element={<Lobby token={token} />} />
          <Route path='/lobby/rooms/:id' element={<GameRoom token={token} />} />
          <Route path='*' element={<Lobby token={token} />} />
        </Routes>
      </Router>
    </DndProvider>
  );
}
export default App;
