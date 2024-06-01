import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Section from './Section';

const Profile = ({ setIsLoggedIn }) => {
  const [user, setUser] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchUserData = async () => {
      const token = localStorage.getItem('token');
      if (!token) {
        navigate('/login');
        return;
      }

      try {
        const response = await fetch('http://localhost:8080/users/me', {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        if (response.ok) {
          const data = await response.json();
          console.log(data);
          setUser(data);
        } else {
          console.error('Failed to fetch user data');
          navigate('/login');
        }
      } catch (error) {
        console.error('Error fetching user data', error);
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
        </div>
      </div>
    </Section>
  );
};

export default Profile;
