import React, { useEffect, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import Section from "./Section";

const MainMenu = ({setIsLoggedIn}) => {
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

  const actions = [
    { id: "1", url: "/courses", title: "Courses" },
    { id: "2", url: "/profile", title: "Profile" },
    { id: "3", url: "/logout", title: "Logout" },
  ];

  return (
    <Section>
      <div className="main-menu">
        <span>Welcome, {user.username}!</span>
        <nav>
          <ul>
            {actions.map(action => (
              <li key={action.id}>
                <Link to={action.url}>{action.title}</Link>
              </li>
            ))}
          </ul>
        </nav>
      </div>
    </Section>
  );
};

export default MainMenu;
