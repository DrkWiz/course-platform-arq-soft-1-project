import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Section from './Section';
import Button from './Button'; // Assuming you have a Button component

const AllCourses = ({ setIsLoggedIn }) => {
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
        const response = await fetch('http://localhost:8080/courses', {
          method: 'GET',
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
      <div className="flex justify-center">
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 ">
          {courses.map((course) => (
            <div key={course.id_course} className="bg-gray-800 rounded-lg shadow-lg w-80 ">
              <div className="p-1 bg-gradient-to-r from-cyan-400 via-yellow-500 to-pink-500 rounded-lg shadow-lg w-82">
              <div className="bg-gray-800 rounded-lg shadow-lg w-79 ">
              <img
                src={`http://localhost:8080/uploads/${course.picture_path}`}
                
                alt={course.name}
                className="w-full h-64 object-cover rounded-t-lg justify-center align-center content-center text-center"
              />
              <div className="p-4 text-center">
                <h3 className="text-xl font-semibold mb-2">{course.name}</h3>
                <p className="text-sm text-gray-400 mb-4">{course.description}</p>
                <Button className="mx-auto"
                  onClick={() => navigate(`/courses/${course.id}`)}
                >
                  View Details
                </Button>
              </div>
            </div>
          </div>
          </div>
          ))}
        </div>
      </div>
  );
};

export default AllCourses;
