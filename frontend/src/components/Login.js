import React, { useState } from "react";
import { useNavigate } from "react-router-dom"; // Para redirigir
import HeadBar from './HeadBar';
import "./login.css";
import logo from "../logo.svg";

function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    const response = await fetch("http://localhost:5000/user/login", {
      method: "POST", // Método POST para enviar datos al servidor de forma segura (no se ven en la URL) 
      headers: { // Cabeceras para indicar que se envía un JSON y no un formulario normal (application/json) 
        "Content-Type": "application/json", // Tipo de contenido que se envía en el cuerpo de la petición 
      },
      body: JSON.stringify({ username, password }), // Convertir los datos a un JSON 
    });

    if (response.ok) {
      alert("Login exitoso");
      navigate('/app'); // Redirigir al menú de inicio
    } else {
      alert("Login incorrecto");
    }
  };

  const handleLogoClick = () => {
    navigate('/')
  };

  const buttons = [
    { 
      label: 'Botón 1', 
      onClick: () => {
        navigate('/signup'); // Navegar a SignUp
      } 
    },
    { label: 'Botón 2', onClick: () => {} },
    // agregar botones según sea necesario
  ];


  /* const handleSubmit = async (e) => {
    e.preventDefault();
    const data = { username, password };
    console.log(JSON.stringify(data)); // Imprimir el JSON en la consola
  
    const response = await fetch("http://localhost:5000/user/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });
    if (response.ok) {
      alert("Login exitoso");
      navigate('/app'); // Redirigir al menú de inicio
    } else {
      alert("Login incorrecto");
    }
  }; */

  return (
    <div className="Login">
      <HeadBar logo={logo} buttons={buttons} onLogoClick={handleLogoClick} />
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