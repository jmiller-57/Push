import { AuthProvider } from "../auth/AuthContext";
import Router from "./Router"
import "../styles/app.css";

export default function App() {
  return (
    <AuthProvider>
      <Router />
    </AuthProvider>
  );
}