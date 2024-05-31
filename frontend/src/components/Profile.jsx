import React from 'react'
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import Section from './Section';
import {profile} from '../constants';

const Profile = ({setIsLoggedIn}) => {
    const [user, setUser] = useState(null);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchUserData = async () => {
            const token = localStorage.getItem("token");
            if (!token) {
                navigate('/login');
                return;
            }

            try {
                const response = await fetch("http://localhost:8080/users/me", {
                    headers: {
                        "Authorization": `Bearer ${token}`,
                    },
                });

                if (response.ok) {
                    const data = await response.json();
                    console.log(data);
                    setUser(data);
                } else {
                    console.error("Failed to fetch user data");
                    navigate('/login');
                }
            } catch (error) {
                console.error("Error fetching user data", error);
                navigate('/login');
            }
        };

        fetchUserData();
    }, [navigate]);

     if (!user) {
        return <div>Loading...</div>;
    }
  
  
    return (
    
    <Section>
   
   <div className="flex justify-center items-center h-screen">
        <div className="p-8 rounded-lg shadow-lg max-w-md w-full bg-gray-800">
        <h2>Profile</h2>

        <p> {user.name}</p>
        <p>{user.username}!</p>
        <p>{user.email}</p>
        
        {user.is_admin ?  <p>Admin</p> : <p>Normal User</p>}
      
  </div>
    </div>
  </Section>

  )
}

export default Profile

