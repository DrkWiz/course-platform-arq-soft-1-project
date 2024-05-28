// aca se hace un formulario que pide el name, username, password y email, para crear un usuario nuevo, se envia un post de un json a localhost:8080/users y luego espera la respuesta del servidor, si el usuario se crea exitosamente te sale un popup de diga "exito" y luego te envia de vuelta al menu App.js y si la respuesta sale mal te sale un popup de diga "fracaso" y luego te envia de vuelta al menu App.js

import React, { useState } from "react";
import { useNavigate } from "react-router-dom"; // Para redirigir
import "./SignUp.css";

function SignUp() {
    const [name, setName] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [email, setEmail] = useState("");
    const navigate = useNavigate();
    
    /*const handleSubmit = async (e) => {
        e.preventDefault();
        const data = { 
            Name: username, 
            Username: username, 
            Password: password, 
            Email: email 
          };
    
        const response = await fetch("localhost:8080/users", {
        method: "POST", // Método POST para enviar datos al servidor de forma segura (no se ven en la URL)
        headers: { // Cabeceras para indicar que se envía un JSON y no un formulario normal (application/json)
            "Content-Type": "application/json", // Tipo de contenido que se envía en el cuerpo de la petición
        },
        body: JSON.stringify(data), // Convertir los datos a un JSON
        });
    
        if (response.ok) {
        alert("Usuario creado exitosamente");
        navigate('/app'); // Redirigir al menú de inicio
        } else {
        alert("Usuario no creado");
        }
    };*/

    const handleSubmit = async (e) => {
        e.preventDefault();
        const data = { 
            Name: username, 
            Username: username, 
            Password: password, 
            Email: email 
        };
    
        console.log(JSON.stringify(data)); // Imprimir el JSON en la consola
    
        const response = await fetch("http://localhost:8080/users", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        });
    
        if (response.ok) {
            alert("Usuario creado exitosamente");
            window.location.href = "http://localhost:3000";
            window.close();
        } else {
            alert("Usuario no creado");
        }
    };


    
    return (
        <div className="SignIn">
        <form onSubmit={handleSubmit}>
            <input
            type="text"
            placeholder="Name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            />
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
            <input
            type="email"
            placeholder="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            />
            <button type="submit">Sign In</button>
        </form>
        </div>
    );
    }

export default SignUp; // Exportar el componente SignIn para poder usarlo en otros archivos
