import React, { useEffect, useState } from 'react';
import CourseUserBlock from './CourseUserBlock';
import Cookies from 'js-cookie';

const MyCourses = () => {
    const [courses, setCourses] = useState([]);

    useEffect(() => {
        const token = Cookies.get('token');

        fetch('http://localhost:8080/users/courses', {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`
            },
        })
        .then(response => response.json())
        .then(data => {
            setCourses(data);
        })
        .catch((error) => {
            console.error('Error:', error);
        });
    }, []);

    return (
        <div style={{ display: 'grid', gridTemplateColumns: 'repeat(2, 1fr)', gap: '20px' }}>
            {courses.map((course) => (
                <CourseUserBlock
                    key={course.id_course}
                    name={course.name}
                    pathphoto={course.picture_path}
                    descripcion={course.description}
                    categorys={course.categories.map(category => category.name)}
                />
            ))}
        </div>
    );
};

export default MyCourses;