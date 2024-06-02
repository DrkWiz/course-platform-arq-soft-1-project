import React from 'react';
import SearchComponent from './SearchComponent';
import AllCourses from './AllCourses';
import './ParentComponent.css';

const ParentComponent = () => {
    return (
        <div>
            <SearchComponent />

            <div className="mb-[7rem] mt-[7rem]" style={{ display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
                <div className='text-x4 font-semibold' style={{ animation: 'moveUpDown 2s infinite' }}>↓</div>
                <div className="ml-2 mr-2 font-semibold">THERE´S MORE BELOW</div>
                <div className='text-x4 font-semibold' style={{ animation: 'moveUpDown 2s infinite' }}>↓</div>
            </div>

            <AllCourses />
        </div>
    );
};

export default ParentComponent;