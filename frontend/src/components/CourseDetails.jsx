import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import Section from './Section';
import Button from './Button';
import Alert from './Alert';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faStar, faStarHalfAlt } from '@fortawesome/free-solid-svg-icons';
import { faStar as faStarRegular } from '@fortawesome/free-regular-svg-icons';

const StarRating = ({ rating }) => {
  // Ensure rating is a number and within the range 0 to 5
  const validRating = Number.isFinite(rating) ? Math.max(0, Math.min(rating, 5)) : 0;

  const fullStars = Math.floor(validRating);
  const halfStar = validRating - fullStars >= 0.5;
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
const FileItem = ({ file }) => (
  <div key={file.id} className="flex items-center justify-between p-2 border rounded mb-2">
    <div>
      <p className="font-semibold">{file.name}</p>
    </div>
  </div>
);
const CommentForm = ({ handleAddComment }) => {
  const [newComment, setNewComment] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    handleAddComment(newComment);
    setNewComment('');
  };

  return (
    <div className="mb-4">
      <form onSubmit={handleSubmit}>
        <label className="block text-sm font-medium text-gray-400">Add a comment:</label>
        <textarea
          value={newComment}
          onChange={(e) => setNewComment(e.target.value)}
          className="w-full p-2 border rounded"
        />
        <Button type="submit" className="mt-2">Submit</Button>
      </form>
    </div>
  );
};

const FileUploadForm = ({ courseId, onUploadSuccess }) => {
  const [file, setFile] = useState(null);
  const [errorMessage, setErrorMessage] = useState('');
  const [alertType, setAlertType] = useState(null);
  const [showAlert, setShowAlert] = useState(false);

  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const token = localStorage.getItem('token');

    if (!token) {
      setErrorMessage("Error with token");
      setAlertType('error');
      setShowAlert(true);
      return;
    }

    if (!file) {
      setErrorMessage("Please select a file");
      setAlertType('error');
      setShowAlert(true);
      return;
    }

    const formData = new FormData();
    formData.append('file', file);

    try {
      const response = await fetch(`http://localhost:8080/courses/${courseId}/files`, {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${token}`,
        },
        body: formData,
      });

      if (response.ok) {
        setFile(null);
        setErrorMessage("File uploaded successfully");
        setAlertType('success');
        setShowAlert(true);
        if (onUploadSuccess) {
          onUploadSuccess();
        }
      } else {
        setErrorMessage("Failed to upload file");
        setAlertType('error');
        setShowAlert(true);
      }
    } catch (error) {
      setErrorMessage("Error uploading file");
      setAlertType('error');
      setShowAlert(true);
    }
  };

  return (
    <div className="mb-4">
      {showAlert && <Alert message={errorMessage} type={alertType} onClose={() => setShowAlert(false)} />}
      <form onSubmit={handleSubmit}>
        <label className="block text-sm font-medium text-gray-400">Upload a file:</label>
        <input type="file" onChange={handleFileChange} className="w-full p-2 border rounded" />
        <Button type="submit" className="mt-2">Upload</Button>
      </form>
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
  const [comments, setComments] = useState([]); // Initialize as an empty array
  const navigate = useNavigate();
  const [files, setFiles] = useState([]);
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

          // Fetch comments
          const commentsResponse = await fetch(`http://localhost:8080/courses/${id}/comments`, {
            headers: {
              "Authorization": `Bearer ${token}`,
            },
          });

          if (commentsResponse.ok) {
            const commentsData = await commentsResponse.json();
            setComments(Array.isArray(commentsData) ? commentsData : []);
          } else {
            console.error("Failed to fetch comments");
            setComments([]);
          }

          // Fetch user details to check role and ownership
          const userResponse = await fetch(`http://localhost:8080/users/me`, {
            headers: {
              "Authorization": `Bearer ${token}`,
            },
          });

          if (userResponse.ok) {
            const userData = await userResponse.json();
            setIsAdmin(userData.is_admin);

            // Check if the user is the owner of the course
            const ownerResponse = await fetch(`http://localhost:8080/courses/${id}/owner`, {
              method: 'POST',
              headers: {
                "Authorization": `Bearer ${token}`,
              },
            });
              // Fetch files
          const filesResponse = await fetch(`http://localhost:8080/courses/${id}/files`, {
            headers: {
              "Authorization": `Bearer ${token}`,
            },
          });

          if (filesResponse.ok) {
            const filesData = await filesResponse.json();
            setFiles(Array.isArray(filesData) ? filesData : []);
          } else {
            console.error("Failed to fetch files");
            setFiles([]);
          }
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

  const handleAddComment = async (comment) => {
    const token = localStorage.getItem("token");

    try {
        const response = await fetch(`http://localhost:8080/courses/${id}/comments`, {
            method: 'POST',
            headers: {
                "Authorization": `Bearer ${token}`,
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ comment }),
        });

        if (response.ok) {
            setErrorMessage("Comment added successfully");
            setAlertType('success');
            setShowAlert(true);
        } else {
            const errorData = await response.json();
            console.error("Failed to add comment", errorData);
            setErrorMessage("Failed to add comment");
            setAlertType('error');
            setShowAlert(true);
        }
    } catch (error) {
        console.error("Error adding comment", error);
        setErrorMessage("Error adding comment");
        setAlertType('error');
        setShowAlert(true);
    }
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
        <div className="p-1 bg-gradient-to-r from-cyan-400 via-yellow-500 to-pink-500 rounded-lg shadow-lg max-w-md w-full ml-4">
          <div className="p-8 rounded-lg shadow-lg max-w-md w-full bg-gray-800 text-white">
            {averageRating !== null && (
              <>
                <h2 className="text-2xl font-bold mb-4">Average Rating</h2>
                <StarRating rating={averageRating} />
                <p className="text-lg mt-2">{averageRating.toFixed(1)} / 5.0</p>
              </>
            )}
            <h2 className="text-2xl font-bold mb-4 mt-4">User Comments</h2>
            {Array.isArray(comments) && comments.length > 0 ? (
              comments.map((comment, index) => (
                <div key={index} className="mb-4">
                  <p className="text-lg font-bold">{comment.username}</p>
                  <StarRating rating={comment.rating} />
                  <p className="text-lg">{comment.comment}</p>
                </div>
              ))
            ) : (
              <p>No comments available.</p>
            )}
            {isEnrolled && (
              <CommentForm handleAddComment={handleAddComment} />
            )}
            {(isAdmin || isOwner) && (
              <FileUploadForm courseId={id} />
            )}

<h2 className="text-2xl font-bold mb-4 mt-4">Course Files</h2>
            {Array.isArray(files) && files.length > 0 ? (
              files.map((file, index) => (
                <FileItem key={index} file={file} />
              ))
            ) : (
              <p>No files available.</p>
            )}
          </div>
        </div>
      </div>
    </Section>
  );
};

export default CourseDetails;
