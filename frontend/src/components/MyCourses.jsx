import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Section from './Section';
import Button from './Button'; // Assuming you have a Button component

const MyCourses = ({ setIsLoggedIn }) => {
  const [courses, setCourses] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchCourses = async () => {
      const token = localStorage.getItem("token");
      if (!token) {
        navigate('/login');
        return;
      }

      try {
        const response = await fetch('http://localhost:8080/users/courses/all', {
          method: 'GET',
          headers: {
            "Authorization": `Bearer ${token}`,
          },
        });

        if (response.ok) {
          const data = await response.json();
          setCourses(data);
        } else {
          console.error("Failed to fetch courses");
        }
      } catch (error) {
        console.error("Error fetching courses", error);
      }
    };

    fetchCourses();
  }, [navigate]);

  if (courses.length === 0) {
    return <div>Loading...</div>;
  }

  return (
    <Section>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {courses.map((course) => (
          <div key={course.id_course} className="p-4 bg-gray-800 rounded-lg shadow-lg">
            <img 
              src={`http://localhost:8080/uploads/${course.picture_path}`} 
              alt={course.name} 
              className="w-full h-auto rounded-t-lg"
            />
            <div className="p-4">
              <h3 className="text-xl font-semibold mb-2">{course.name}</h3>
              <p className="text-sm text-gray-400 mb-4">{course.description}</p>
              <Button onClick={() => navigate(`/courses/${course.id_course}`)}>View Details</Button>
            </div>
          </div>
        ))}
      </div>
    </Section>
  );
};

export default MyCourses;
