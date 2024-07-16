import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Button from './Button'; // Assuming you have a Button component
import ShowMoreText from './ShowMoreText'; // Import the new component

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
        const response = await fetch('/backend/courses', {
          method: 'GET',
        });

        if (response.ok) {
          const data = await response.json();
          const activeCourses = data.filter(course => course.is_active);
          setCourses(activeCourses);
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
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
        {courses.map((course) => (
          <div key={course.id_course} className="bg-gray-800 rounded-lg shadow-lg w-80">
            <div className="p-1 bg-gradient-to-r from-cyan-400 via-yellow-500 to-pink-500 rounded-lg shadow-lg w-full h-full">
              <div className="bg-gray-800 rounded-lg shadow-lg flex flex-col h-full">
                <img
                  src={`/backend/uploads/${course.picture_path}`}
                  alt={course.name}
                  className="w-full h-64 object-cover rounded-t-lg"
                />
                <div className="flex flex-col justify-between p-4 text-center flex-grow">
                  <div>
                    <h3 className="text-xl font-semibold mb-2">{course.name}</h3>
                    <ShowMoreText text={course.description} maxLength={100} /> {/* Use the component here */}
                  </div>
                  <Button className="mx-auto mt-4" onClick={() => navigate(`/courses/${course.id}`)}>
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
