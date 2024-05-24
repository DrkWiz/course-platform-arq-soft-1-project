import React, { useState } from "react";
import { useNavigate } from "react-router-dom"; // Para redirigir
import "./login.css";

function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    const response = await fetch("http://localhost:5000/user/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ username, password }),
    });

    if (response.ok) {
      alert("Login exitoso");
      navigate('/app'); // Redirigir al menú de inicio
    } else {
      alert("Login incorrecto");
    }
  };

  return (
    <div className="Login">
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <button type="submit">Login</button>
      </form>
      {/* 
      Parte comentada para guardar username y password en un JSON y enviarlo a una URL, esperando la confirmación 
      
      const handleLogin = async () => {
        const userData = { username, password };
        console.log("Datos del usuario:", userData);
        
        const response = await fetch("http://tudominio.com/api/login", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(userData),
        });

        if (response.ok) {
          const result = await response.json();
          if (result.success) { // Suponiendo que la respuesta tiene un campo 'success'
            console.log("Datos enviados exitosamente");
            navigate('/app'); // Redirigir al menú de inicio
          } else {
            alert("Login incorrecto");
          }
        } else {
          console.error("Error al enviar los datos");
        }
      };
      
      */}
    </div>
  );
}

export default Login;