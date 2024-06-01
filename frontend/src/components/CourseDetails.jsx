import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import Section from './Section';

const CourseDetails = () => {
  const { id } = useParams();
  const [course, setCourse] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    console.log("Fetching course data for course ID:", id);
    const fetchCourseData = async () => {
      const token = localStorage.getItem("token");
      if (!token) {
        navigate('/login');
        return;
      }

      try {
        const response = await fetch(`http://localhost:8080/courses/${id}`, {
          headers: {
            "Authorization": `Bearer ${token}`,
          },
        });

        if (response.ok) {
          const data = await response.json();
          setCourse(data);
        } else {
          console.error("Failed to fetch course data");
          navigate('/login');
        }
      } catch (error) {
        console.error("Error fetching course data", error);
        navigate('/login');
      }
    };

    fetchCourseData();
  }, [id, navigate]);

  if (!course) {
    return <div>Loading...</div>;
  }

  return (
    <Section>
      <div className="flex justify-center items-center h-screen">
        <div className="p-8 rounded-lg shadow-lg max-w-md w-full bg-gray-800 text-white">
          <h2 className="text-2xl font-bold mb-4">Course Details</h2>
          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-400">Course:</label>
            <p className="text-lg">{course.name}</p>
          </div>
          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-400">Description:</label>
            <p className="text-lg">{course.description}</p>
          </div>
          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-400">Price:</label>
            <p className="text-lg">{course.price}</p>
          </div>
          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-400">Is it active?:</label>
            {course.is_active ? <p className="text-green-400">Published</p> : <p className="text-red-400">Not Published</p>}
          </div>
        </div>
      </div>
    </Section>
  );
};

export default CourseDetails;