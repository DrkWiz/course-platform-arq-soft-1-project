import React, { useState } from 'react';

const CourseUserBlock = ({ name, pathphoto, descripcion, categorys }) => {
    const [showDetails, setShowDetails] = useState(false);

    const handleButtonClick = () => {
        setShowDetails(!showDetails);
    };

    const handleUnsubscribe = () => {
        // Send a request to unsubscribe the user from the course
        fetch('http://localhost:8080/user/desuscribir', {
            method: 'POST',
            // Add any necessary headers or body data
        })
            .then(response => {
                // Handle the response
            })
            .catch(error => {
                // Handle any errors
            });
    };

    return (
        <div className="course-user-block">
            <img src={pathphoto} alt={name} />
            <h3>{name}</h3>
            <button onClick={handleButtonClick}>Ver más</button>
            {showDetails && (
                <div className="course-details">
                    <p>{descripcion}</p>
                    <ul>
                        {categorys.map(category => (
                            <li key={category}>{category}</li>
                        ))}
                    </ul>
                    <button onClick={handleUnsubscribe}>Desuscribirse</button>
                </div>
            )}
        </div>
    );
};

export default CourseUserBlock;