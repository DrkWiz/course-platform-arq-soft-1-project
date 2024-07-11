import React, { useEffect, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import Section from "./Section";
import SearchComponent from "./SearchComponent";
import Alert from './Alert';

const MainMenu = ({ setIsLoggedIn }) => {
  const [user, setUser] = useState(null);
  const navigate = useNavigate();
  const [showAlert, setShowAlert] = useState(false);
  const [errorMessage, setErrorMessage] = useState('');
  const [alertType, setAlertType] = useState(null); 

  useEffect(() => {
    const fetchUserData = async () => {
      const token = localStorage.getItem("token");
      if (!token) {
        setErrorMessage("Error with token");
        setAlertType("error");
        setShowAlert(true);
        setTimeout(() => {
        navigate('/login');
        } , 1500);
        return;
      }

      try {
        const response = await fetch("/backend/users/me", {
          headers: {
            "Authorization": `Bearer ${token}`,
          },
        });

        if (response.ok) {
          const data = await response.json();
          setUser(data);
        } else {
          setErrorMessage("Failed to fetch user data");
          setAlertType("error");
          setShowAlert(true);
          console.error("Failed to fetch user data");
          setTimeout(() => {
          navigate('/login');
          } , 1500);
        }
      } catch (error) {
        setErrorMessage("Error fetching user data");
        setAlertType("error");
        setShowAlert(true);
        console.error("Error fetching user data", error);
        setTimeout(() => {
          navigate('/login');
        }, 1500);
      }
    };

    fetchUserData();
  }, [navigate]);

  if (!user) {
    return <div>Loading...</div>;
  }

 

  return (
    <Section>
         {showAlert && <Alert message={errorMessage} type={alertType} onClose={() => setShowAlert(false)} />}
      <div className="main-menu" style={{ flexDirection: "column", justifyContent: "center", alignItems: "center" }}>
        <div style={{ width: '50%', margin: '0 auto' }}>
          <SearchComponent />
        </div>
      </div>
    </Section>
  );
};

export default MainMenu;
