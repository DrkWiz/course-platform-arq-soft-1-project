import React, { useState, useEffect } from 'react';
import axios from 'axios';

const MyCourses = () => {
  const [userCourses, setUserCourses] = useState([]);
  const [jsonResponse, setJsonResponse] = useState('');

  useEffect(() => {
    const fetchCourses = async () => {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get('http://localhost:8080/users/courses', {
          headers: { 'Authorization': `Bearer ${token}` }
        });
        const courses = response.data;
        console.log('Fetched user courses:', courses);
        setUserCourses(courses);
        setJsonResponse(JSON.stringify(courses, null, 2));
      } catch (error) {
        console.error('Error fetching user courses:', error);
      }
    };

    fetchCourses();
  }, []);

  return (
    <div>
      <h1>My Courses</h1>
      {/* Aquí puedes renderizar los cursos del usuario como prefieras */}
      <textarea readOnly value={jsonResponse} style={{ width: '100%', height: '200px' }} />
    </div>
  );
};

export default MyCourses;