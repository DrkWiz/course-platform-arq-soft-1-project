import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Section from './Section';
import Button from './Button';
import Alert from './Alert';

const Profile = ({ setIsLoggedIn }) => {
  const [user, setUser] = useState(null);
  const navigate = useNavigate();
  const [showAlert, setShowAlert] = useState(false);
  const [errorMessage, setErrorMessage] = useState('');
  const [alertType, setAlertType] = useState(null);

  useEffect(() => {
    const fetchUserData = async () => {
      const token = localStorage.getItem('token');
      if (!token) {
        setErrorMessage('Error with token');
        setAlertType('error');
        setShowAlert(true);
        setTimeout(() => {
        navigate('/login');
        } , 1500);
        return;
      }

      try {
        const response = await fetch('/backend/users/me', {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        if (response.ok) {
          const data = await response.json();
          console.log(data);
          setUser(data);
        } else {
          setErrorMessage('Failed to fetch user data');
          setAlertType('error');
          setShowAlert(true);
          console.error('Failed to fetch user data');
          setTimeout(() => {
          navigate('/login');
          } , 1500);
        }
      } catch (error) {
        setErrorMessage('Error fetching user data');
        setAlertType('error');
        setShowAlert(true);
        console.error('Error fetching user data', error);
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
    <Section className="flex justify-center items-center h-screen -mt-20 -mb-20">
    {showAlert && <Alert message={errorMessage} type={alertType} onClose={() => setShowAlert(false)} />}
      <div className="p-1 bg-gradient-to-r from-cyan-400 via-yellow-500 to-pink-500 rounded-lg shadow-lg max-w-md w-full">
        <div className="p-8 rounded-lg shadow-lg max-w-md w-full bg-gray-800 text-white">
          <h2 className="text-2xl font-bold mb-4">Profile</h2>
          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-400">Name:</label>
            <p className="text-lg">{user.name}</p>
          </div>
          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-400">Username:</label>
            <p className="text-lg">{user.username}</p>
          </div>
          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-400">Email:</label>
            <p className="text-lg">{user.email}</p>
          </div>
          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-400">Role:</label>
            {user.is_admin ? <p className="text-green-400">Admin</p> : <p className="text-blue-400">Normal User</p>}
          </div>

          {user.is_admin ? <Button className="w-half bg-gray-800 text-white hover:bg-gray-800 rounded text-2xl font-semibold" onClick={() => navigate('/create')}>Create Course</Button> : null}
        </div>
      </div>
    </Section>
  );
};

export default Profile;
