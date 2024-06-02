
import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import Select from "react-select";
import Section from './Section';
import Button from './Button';

const EditCourse = () => {
  const { id } = useParams();
  const [course, setCourse] = useState({
    owner: '',
    name: '',
    description: '',
    price: '',
    picture_path: '',
    start_date: '',
    end_date: '',
    is_active: false,
    categories: [],
  });
  const [categories, setCategories] = useState([]);
  const [selectedCategories, setSelectedCategories] = useState([]);
  const [loading, setLoading] = useState(true);
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
        } else {
          console.error("Failed to fetch course data");
          navigate('/login');
        }
      } catch (error) {
        console.error("Error fetching course data", error);
        navigate('/login');
      } finally {
        setLoading(false);
      }
    };

    const fetchCategories = async () => {
      const response = await fetch("http://localhost:8080/category/all");
      const data = await response.json();
      const formattedData = data.map(category => ({
        value: category.id,
        label: category.name
      }));
      setCategories(formattedData);
      console.log(formattedData);
    };


    fetchCourseData();
    fetchCategories();
  }, [id, navigate]);

  const handleChange = (e) => {
    const { name, value, type, checked } = e.target;
    setCourse((prevCourse) => ({
      ...prevCourse,
      [name]: type === 'checkbox' ? checked : value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const token = localStorage.getItem("token");

    try {
      const response = await fetch(`http://localhost:8080/courses/update/${id}`, {
        method: 'PUT',
        headers: {
          "Authorization": `Bearer ${token}`,
          "Content-Type": "application/json",
        },
        body: JSON.stringify(course),
      });

      if (response.ok) {
        navigate(`/courses/${id}`);
      } else {
        console.error("Failed to update course");
      }
    } catch (error) {
      console.error("Error updating course", error);
    }


  };

  const customStyles = {
    control: (provided) => ({
      ...provided,
      backgroundColor: '#2d3748', // Match your dark theme
      borderColor: '#4a5568', // Match your dark theme
      color: 'white',
    }),
    menu: (provided) => ({
      ...provided,
      backgroundColor: '#2d3748', // Match your dark theme
    }),
    option: (provided, state) => ({
      ...provided,
      backgroundColor: state.isSelected ? '#4a5568' : '#2d3748', // Match your dark theme
      color: 'white',
      '&:hover': {
        backgroundColor: '#4a5568', // Match your dark theme
      },
    }),
    multiValue: (provided) => ({
      ...provided,
      backgroundColor: '#4a5568', // Match your dark theme
    }),
    multiValueLabel: (provided) => ({
      ...provided,
      color: 'white',
    }),
    placeholder: (provided) => ({
      ...provided,
      color: '#a0aec0', // Match your dark theme
    }),
  };

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <Section>
      <div className="flex justify-center items-center h-screen">
        <div className="p-1 bg-gradient-to-r from-cyan-400 via-yellow-500 to-pink-500 rounded-lg shadow-lg max-w-md w-full">
          <div className="p-8 rounded-lg shadow-lg max-w-md w-full bg-gray-800 text-white">
            <h2 className="text-2xl font-bold mb-4">Edit Course</h2>
            <form onSubmit={handleSubmit}>
              <div className="mb-4">
                <label className="block text-sm font-medium text-gray-400">Course Name</label>
                <input
                  type="text"
                  name="name"
                  value={course.name}
                  onChange={handleChange}
                  className="mt-1 block w-full bg-gray-700 text-white rounded-md"
                />
              </div>
              <div className="mb-4">
                <label className="block text-sm font-medium text-gray-400">Description</label>
                <textarea
                  name="description"
                  value={course.description}
                  onChange={handleChange}
                  className="mt-1 block w-full bg-gray-700 text-white rounded-md"
                />
              </div>
              <div className="mb-4">
                <label className="block text-sm font-medium text-gray-400">Price</label>
                <input
                  type="number"
                  name="price"
                  value={course.price}
                  onChange={handleChange}
                  className="mt-1 block w-full bg-gray-700 text-white rounded-md"
                />
              </div>
              <div className="mb-4">
                <label className="block text-sm font-medium text-gray-400">Start Date</label>
                <input
                  type="text"
                  name="start_date"
                  value={course.start_date}
                  onChange={handleChange}
                  className="mt-1 block w-full bg-gray-700 text-white rounded-md"
                />
              </div>
              <div className="mb-4">
                <label className="block text-sm font-medium text-gray-400">End Date</label>
                <input
                  type="text"
                  name="end_date"
                  value={course.end_date}
                  onChange={handleChange}
                  className="mt-1 block w-full bg-gray-700 text-white rounded-md"
                />
              </div>
              <div className='mb-4'>
                <label className="block text-sm font-medium text-gray-400">Categories</label>
                <Select
                  id="categories"
                  name="categories"
                  isMulti
                  options={categories}
                  className="basic-multi-select"
                  classNamePrefix="select"
                  styles={customStyles}
                  onChange={setSelectedCategories}
                />
              </div>
              <div className="mb-4">
                <label className="block text-sm font-medium text-gray-400">Is Active?</label>
                <input
                  type="checkbox"
                  name="is_active"
                  checked={course.is_active}
                  onChange={handleChange}
                  className="mt-1 block"
                />
              </div>
              <div className="mb-4">
                <label className="block text-sm font-medium text-gray-400">Course Image</label>
                <input
                  type="text"
                  name="picture_path"
                  value={course.picture_path}
                  onChange={handleChange}
                  className="mt-1 block w-full bg-gray-700 text-white rounded-md"
                />
              </div>
              <div className="mt-4">
                <Button type="submit">Save Changes</Button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </Section>
  );
};

export default EditCourse;
