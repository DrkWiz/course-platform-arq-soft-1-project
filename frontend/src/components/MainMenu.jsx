import React, { useEffect, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import Section from "./Section";
import SearchComponent from "./SearchComponent";

const MainMenu = ({ setIsLoggedIn }) => {
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
      <div className="main-menu" style={{ flexDirection: "column", justifyContent: "center", alignItems: "center" }}>
        <div style={{ width: '50%', margin: '0 auto' }}>
          <SearchComponent />
        </div>
      </div>
    </Section>
  );
};

export default MainMenu;
