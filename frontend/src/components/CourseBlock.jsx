import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useCookies } from 'react-cookie';
import Boton from './Button';

function CourseBlock({ course }) {
  const [cookies, setCookie, removeCookie] = useCookies(['token']);
  const [isSubscribed, setIsSubscribed] = useState(false);
  const [showDetails, setShowDetails] = useState(false);

  useEffect(() => {
    axios.post('http://localhost:8080/userissuscribed', {
      token: cookies.token,
      courseId: course.id,
    })
    .then(response => setIsSubscribed(response.data.isSubscribed))
    .catch(error => console.error(error));
  }, [course.id, cookies.token]);

  const handleSubscription = () => {
    const url = isSubscribed ? 'http://localhost:8080/desuscribiruser' : 'http://localhost:8080/suscribirusuario';
    axios.post(url, {
      token: cookies.token,
      courseId: course.id,
    })
    .then(response => setIsSubscribed(!isSubscribed))
    .catch(error => console.error(error));
  };

  return (
    <div className="course-block">
      <img src={course.photo} alt={course.name} />
      <h2>{course.name}</h2>
      <Boton onClick={() => setShowDetails(!showDetails)}>Ver más</Boton>
      {showDetails && (
        <div className="course-details">
          <h3>{course.name}</h3>
          <img src={course.photo} alt={course.name} />
          <p>{course.category}</p>
          <p>{course.description}</p>
          <Boton onClick={handleSubscription}>
            {isSubscribed ? 'Desinscribirse' : `Inscribirse - ${course.price}`}
          </Boton>
        </div>
      )}
    </div>
  );
}

export default CourseBlock;