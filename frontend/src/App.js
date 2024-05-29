import React from "react";
import { useNavigate } from "react-router-dom";
import logo from "./logo.svg";
import "./App.css";
//agrego imports para crear un boton que me lleve al login.js
import "./components/Login.js";
import "./components/login.css";
import "./components/SignUp.js";
import "./components/SignUp.css";

function App() {
  const navigate = useNavigate();
  const openLogin = () => {
   /* window.open("/login", "_blank", "noopener,noreferrer"); */
  navigate('/Login')
  };

  const openSignUp = () => {
    /* window.open("/signup", "_blank", "noopener,noreferrer"); */
    navigate('/signup')
  };

  return (
    <div className="App">

      <div className="BarraSuperiorFlotante">
        <img src={logo} className="App-logo" alt="logo" />
        <button onClick={openLogin}> Login</button>
        <button onClick={openSignUp}> Sign Up</button>
      </div>

    </div>
  );
}

export default App;
