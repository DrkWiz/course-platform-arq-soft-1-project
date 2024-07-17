import React from 'react';
import './CourseCreation.jsx';
import Section from './Section';
import CourseCreation from './CourseCreation.jsx';

const ParentComponent = () => {
    return (
        <Section className="mt-15">

            <CourseCreation />

        </Section >

    );
};

export default ParentComponent;