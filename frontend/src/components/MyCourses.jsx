import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Section from './Section';
import Button from './Button'; // Assuming you have a Button component
import Alert from './Alert';

const MyCourses = ({ setIsLoggedIn }) => {
  const [courses, setCourses] = useState([]);
  const navigate = useNavigate();
  const [showAlert, setShowAlert] = useState(false);
  const [errorMessage, setErrorMessage] = useState('');
  const [alertType, setAlertType] = useState(null); 

  useEffect(() => {
    const fetchCourses = async () => {
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
        const response = await fetch('/backend/users/courses/all', {
          method: 'GET',
          headers: {
            "Authorization": `Bearer ${token}`,
          },
        });

        if (response.ok) {
          const data = await response.json();
          setCourses(data);
        } else {
          setErrorMessage("Failed to fetch courses");
          setAlertType("error");
          setShowAlert(true);
          console.error("Failed to fetch courses");
        }
      } catch (error) {
        setErrorMessage("Error fetching courses");
        setAlertType("error");
        setShowAlert(true);
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
    {showAlert && <Alert message={errorMessage} type={alertType} onClose={() => setShowAlert(false)} />}
      <div className="flex justify-center">
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 ">
          {courses.map((course) => (
            <div key={course.id_course} className="bg-gray-800 rounded-lg shadow-lg w-80 ">
              <div className="p-1 bg-gradient-to-r from-cyan-400 via-yellow-500 to-pink-500 rounded-lg shadow-lg w-82">
              <div className="bg-gray-800 rounded-lg shadow-lg w-79 ">
              <img
                src={`/backend/uploads/${course.picture_path}`}
                
                alt={course.name}
                className="w-full h-64 object-cover rounded-t-lg"
              />
              <div className="p-4 text-center">
                <h3 className="text-xl font-semibold mb-2">{course.name}</h3>
                <p className="text-sm text-gray-400 mb-4">{course.description}</p>
                <Button className="mx-auto"
                  onClick={() => navigate(`/courses/${course.id_course}`)}
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
    </Section>
  );
};

export default MyCourses;
