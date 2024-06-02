import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import Section from './Section';
import Button from './Button'; // Assuming you have a Button component

const CourseDetails = () => {
  const { id } = useParams();
  const [course, setCourse] = useState(null);
  const [isEnrolled, setIsEnrolled] = useState(false);
  const [isAdmin, setIsAdmin] = useState(false);
  const [isOwner, setIsOwner] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
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

          // Fetch user details to check role and ownership
          const userResponse = await fetch(`http://localhost:8080/users/me`, {
            headers: {
              "Authorization": `Bearer ${token}`,
            },
          });

          if (userResponse.ok) {
            const userData = await userResponse.json();
            console.log(userData);
            setIsAdmin(userData.is_admin);

            // Check if the user is the owner of the course
            const ownerResponse = await fetch(`http://localhost:8080/courses/${id}/owner`, {
              method: 'POST',
              headers: {
                "Authorization": `Bearer ${token}`,
              },
            });

            if (ownerResponse.ok) {
              console.log("Owner response", ownerResponse);
              const ownerData = await ownerResponse.json();
              console.log(ownerData);
              setIsOwner(ownerResponse);

              // Fetch enrollment status
              const enrollmentResponse = await fetch(`http://localhost:8080/users/courses/${id}/enrolled`, {
                headers: {
                  "Authorization": `Bearer ${token}`,
                },
              });

              if (enrollmentResponse.ok) {
                const enrollmentData = await enrollmentResponse.json();
                setIsEnrolled(enrollmentData);
              } else {
                console.error("Failed to fetch enrollment status");
              }
            } else {
              console.error("Failed to check course ownership");
            }
          } else {
            console.error("Failed to fetch user data");
            navigate('/login');
          }
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

  const handleEnroll = async () => {
    const token = localStorage.getItem("token");
    try {
      const response = await fetch(`http://localhost:8080/users/courses/${id}`, {
        method: 'POST',
        headers: {
          "Authorization": `Bearer ${token}`,
          "Content-Type": "application/json",
        },
      });

      if (response.ok) {
        setIsEnrolled(true);
      } else {
        console.error("Failed to enroll in course");
      }
    } catch (error) {
      console.error("Error enrolling in course", error);
    }
  };

  const handleUnenroll = async () => {
    const token = localStorage.getItem("token");
    try {
      const response = await fetch(`http://localhost:8080/users/courses/${id}/unsubscribe`, {
        method: 'DELETE',
        headers: {
          "Authorization": `Bearer ${token}`,
          "Content-Type": "application/json",
        },
      });

      if (response.ok) {
        setIsEnrolled(false);
      } else {
        console.error("Failed to unenroll from course");
      }
    } catch (error) {
      console.error("Error unenrolling from course", error);
    }
  };

  const handleDelete = async () => {
    const token = localStorage.getItem("token");
    try {
      const response = await fetch(`http://localhost:8080/courses/delete/${id}`, {
        method: 'PUT',
        headers: {
          "Authorization": `Bearer ${token}`,
          "Content-Type": "application/json",
        },
      });

      if (response.ok) {
        navigate('/courses'); // Redirect to courses list after deletion
      } else {
        console.error("Failed to delete course");
      }
    } catch (error) {
      console.error("Error deleting course", error);
    }
  };

  const handleModify = () => {
    navigate(`/courses/${id}/edit`);
  };

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
            <label className="block text-sm font-medium text-gray-400">Start Date:</label>
            <p className="text-lg">{course.start_date}</p>
          </div>
          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-400">End Date:</label>
            <p className="text-lg">{course.end_date}</p>
          </div>
          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-400">Is it active?:</label>
            {course.is_active ? <p className="text-green-400">Published</p> : <p className="text-red-400">Not Published</p>}
          </div>
          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-400">Course Image:</label>
            <img 
              src={`http://localhost:8080/uploads/${course.picture_path}`} 
              alt={course.name} 
              style={{
                width: '100%', // make the image take up the full width of its container
                height: 'auto', // keep the original aspect ratio
                objectFit: 'cover', // cover the entire width of the container without stretching
                borderRadius: '4px', // round the corners
              }}
            />
          </div>
          <div className="mt-4">
            {isEnrolled ? (
              <Button onClick={handleUnenroll}>Unenroll</Button>
            ) : (
              <Button onClick={handleEnroll}>Enroll</Button>
            )}
          </div>
          {console.log(isAdmin, isOwner)}
          {(isAdmin && isOwner) && (
            <div className="mt-4">
              <Button onClick={handleModify} className="mr-2">Modify</Button>
              <Button onClick={handleDelete} className="bg-red-500 hover:bg-red-700">Delete</Button>
            </div>
          )}
        </div>
      </div>
    </Section>
  );
};

export default CourseDetails;
