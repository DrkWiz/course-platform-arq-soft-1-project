import React, { useState, useEffect } from "react";
import { Route, Routes, useNavigate } from "react-router-dom";
import ButtonGradient from './assets/svg/ButtonGradient';
import Benefits from './components/Benefits';
import Footer from './components/Footer';
import Header from './components/Header';
import Hero from './components/Hero';
import Pricing from './components/Pricing';
import Login from './components/Login';
import Register from './components/Register';
import MyCourses from "./components/MyCourses";
import MainMenu from './components/MainMenu';
import Profile from './components/Profile';
import CourseDetails from './components/CourseDetails';
import CourseCreation from './components/CourseCreation';
import EditCourse from './components/EditCourse';

const App = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      setIsLoggedIn(true);
    }
  }, []);

  const handleLogout = () => {
    localStorage.removeItem("token");
    setIsLoggedIn(false);
    navigate("/login");
  };

  return (
    <div className="pt-[4.75rem] lg:pt-[5.25rem] overflow-hidden">
      <Header isLoggedIn={isLoggedIn} handleLogout={handleLogout} />
      <Routes>
        <Route path="/" element={
          <>
            <Hero />
            <Benefits />
            <Pricing />
          </>
        } />
        <Route path="/login" element={<Login setIsLoggedIn={setIsLoggedIn} isLoggedIn={isLoggedIn}/>} />
        <Route path="/register" element={<Register />} />
        <Route path="/mycourses" element={<MyCourses setIsLoggedIn={setIsLoggedIn} />} />
        <Route path="/mainmenu" element={<MainMenu />} />
        <Route path="/profile" element={<Profile setIsLoggedIn={setIsLoggedIn} />} />
        <Route path="/courses/:id" element={<CourseDetails />} /> {/* Ensure this route is defined correctly */}
        <Route path="/create" element={<CourseCreation />} /> {/* Ensure this route is defined correctly */}
        <Route path="/courses/:id/edit" element={<EditCourse />} />
      </Routes>
      <Footer />
      <ButtonGradient />
    </div>
  );
};

export default App;
