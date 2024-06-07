import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import Section from './Section';
import Button from './Button';
import Alert from './Alert';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faStar, faStarHalfAlt } from '@fortawesome/free-solid-svg-icons';
import { faStar as faStarRegular } from '@fortawesome/free-regular-svg-icons';

const StarRating = ({ rating }) => {
  const fullStars = Math.floor(rating);
  const halfStar = rating - fullStars >= 0.5;
  const emptyStars = 5 - fullStars - (halfStar ? 1 : 0);

  return (
    <div className="flex">
      {Array(fullStars).fill().map((_, i) => (
        <FontAwesomeIcon key={`full-${i}`} icon={faStar} className="text-yellow-500" />
      ))}
      {halfStar && <FontAwesomeIcon icon={faStarHalfAlt} className="text-yellow-500" />}
      {Array(emptyStars).fill().map((_, i) => (
        <FontAwesomeIcon key={`empty-${i}`} icon={faStarRegular} className="text-yellow-500" />
      ))}
    </div>
  );
};
const CourseDetails = () => {
  const { id } = useParams();
  const [course, setCourse] = useState(null);
  const [isEnrolled, setIsEnrolled] = useState(false);
  const [isAdmin, setIsAdmin] = useState(false);
  const [isOwner, setIsOwner] = useState(false);
  const [averageRating, setAverageRating] = useState(null);
  const navigate = useNavigate();
  const [showAlert, setShowAlert] = useState(false);
  const [errorMessage, setErrorMessage] = useState('');
  const [alertType, setAlertType] = useState(null);

  useEffect(() => {
    const fetchCourseData = async () => {
      const token = localStorage.getItem("token");
      if (!token) {
        setErrorMessage("Error with token");
        setAlertType('error');
        setShowAlert(true);
        setTimeout(() => {
          navigate('/login');
        }, 1500);
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

          // Fetch average rating
          const ratingResponse = await fetch(`http://localhost:8080/courses/${id}/rating`, {
            headers: {
              "Authorization": `Bearer ${token}`,
            },
          });

          if (ratingResponse.ok) {
            const ratingData = await ratingResponse.json();
            setAverageRating(ratingData); // Assuming the response has an `average_rating` field
          } else {
            setAverageRating(0);
            console.error("Failed to fetch average rating");
          }

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
              const ownerData = await ownerResponse.json();
              setIsOwner(ownerData);

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
                setErrorMessage("Failed to fetch enrollment status");
                setAlertType('error');
                setShowAlert(true);
              }
            } else {
              console.error("Failed to check course ownership");
              setErrorMessage("Failed to check course ownership");
              setAlertType('error');
              setShowAlert(true);
            }
          } else {
            console.error("Failed to fetch user data");
            setErrorMessage("Failed to fetch user data");
            setAlertType('error');
            setShowAlert(true);
            setTimeout(() => {
              navigate('/login');
            }, 1500);
          }
        } else {
          console.error("Failed to fetch course data");
          setErrorMessage("Failed to fetch course data");
          setAlertType('error');
          setShowAlert(true);
          setTimeout(() => {
            navigate('/login');
          }, 1500);
        }
      } catch (error) {
        console.error("Error fetching course data", error);
        setErrorMessage("Error fetching course data");
        setAlertType('error');
        setShowAlert(true);
        setTimeout(() => {
          navigate('/login');
        }, 1500);
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
        setErrorMessage("Enrolled in course successfully");
        setAlertType('success');
        setShowAlert(true);
        setIsEnrolled(true);
      } else {
        console.error("Failed to enroll in course");
        setErrorMessage("Failed to enroll in course");
        setAlertType('error');
        setShowAlert(true);
      }
    } catch (error) {
      console.error("Error enrolling in course", error);
      setErrorMessage("Error enrolling in course");
      setAlertType('error');
      setShowAlert(true);
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
        setErrorMessage("Unenrolled from course successfully");
        setAlertType('success');
        setShowAlert(true);
        setIsEnrolled(false);
      } else {
        console.error("Failed to unenroll from course");
        setErrorMessage("Failed to unenroll from course");
        setAlertType('error');
        setShowAlert(true);
      }
    } catch (error) {
      console.error("Error unenrolling from course", error);
      setErrorMessage("Error unenrolling from course");
      setAlertType('error');
      setShowAlert(true);
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
        setErrorMessage("Course deleted successfully");
        setAlertType('success');
        setShowAlert(true);
        setTimeout(() => {
          navigate('/mainmenu');
        }, 1500);
      } else {
        console.error("Failed to delete course");
        setErrorMessage("Failed to delete course");
        setAlertType('error');
        setShowAlert(true);
      }
    } catch (error) {
      console.error("Error deleting course", error);
      setErrorMessage("Error deleting course");
      setAlertType('error');
      setShowAlert(true);
    }
  };

  const handleModify = () => {
    navigate(`/courses/${id}/edit`);
  };

  if (!course) {
    return <div>Loading...</div>;
  }

  return (
    <Section className="mt-[13rem] mb-[13rem]" customPaddings>
      {showAlert && <Alert message={errorMessage} type={alertType} onClose={() => setShowAlert(false)} />}
      <div className="flex justify-center items-center h-screen">
        <div className="p-1 bg-gradient-to-r from-cyan-400 via-yellow-500 to-pink-500 rounded-lg shadow-lg max-w-md w-full">
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
              <label className="block text-sm font-medium text-gray-400">Categories:</label>
              <ul className="list-disc list-inside text-lg">
                {course.categories?.map(category => (
                  <li key={category.id}>{category.name}</li>
                ))}
              </ul>
            </div>
            <div className="mb-4">
              <label className="block text-sm font-medium text-gray-400">Course Image:</label>
              <img
                src={`http://localhost:8080/uploads/${course.picture_path}`}
                alt={course.name}
                style={{
                  width: '100%',
                  height: 'auto',
                  objectFit: 'cover',
                  borderRadius: '4px',
                }}
              />
            </div>
            <div className="mt-4 flex justify-center items-center">
              <div className="mt-4 mr-2">
                {isEnrolled ? (
                  <Button onClick={handleUnenroll}>Unenroll</Button>
                ) : (
                  <Button onClick={handleEnroll}>Enroll</Button>
                )}
              </div>
              {(isAdmin && isOwner) && (
                <div className="mt-4 flex">
                  <Button onClick={handleModify} className="mr-2">Modify</Button>
                  <Button onClick={handleDelete} className="text-red-400 hover:text-red-600">Delete</Button>
                </div>
              )}
            </div>
          </div>
        </div>
        {averageRating !== null && (
          <div className="p-1 bg-gradient-to-r from-cyan-400 via-yellow-500 to-pink-500 rounded-lg shadow-lg max-w-md w-full ml-4">
            <div className="p-8 rounded-lg shadow-lg max-w-md w-full bg-gray-800 text-white">
              <h2 className="text-2xl font-bold mb-4">Average Rating</h2>
              <StarRating rating={averageRating} />
              <p className="text-lg mt-2">{averageRating.toFixed(1)} / 5.0</p>
            </div>
          </div>
        )}
      </div>
    </Section>
  );
};

export default CourseDetails;
