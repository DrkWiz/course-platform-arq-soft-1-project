import React, { useState, useCallback } from 'react';
import Autosuggest from 'react-autosuggest';
import axios from 'axios';
import debounce from 'lodash.debounce';
import { useNavigate } from 'react-router-dom';

const SearchComponent = () => {
  const [value, setValue] = useState('');
  const [suggestions, setSuggestions] = useState([]);
  const navigate = useNavigate(); // Initialize navigate

  const fetchSuggestions = async (inputValue) => {
    if (!inputValue) {
      setSuggestions([]);
      return;
    }

    const trimmedValue = inputValue.trim().toLowerCase();
    if (trimmedValue.length === 0) {
      setSuggestions([]);
      return;
    }

    try {
      const response = await axios.get('http://localhost:8080/courses');
      const courses = response.data;
      console.log('Fetched courses:', courses); // Debugging line

      const filteredSuggestions = courses.filter(course => {
        const courseName = course.name;
        if (courseName && typeof courseName === 'string') {
          const isMatch = courseName.toLowerCase().startsWith(trimmedValue);
          console.log(`Checking course "${courseName}": ${isMatch}`); // Debugging line
          return isMatch;
        }
        return false;
      });

      setSuggestions(filteredSuggestions);
      console.log('Filtered suggestions:', filteredSuggestions); // Debugging line
    } catch (error) {
      console.error('Error fetching courses', error);
    }
  };

  const debouncedFetchSuggestions = useCallback(debounce(fetchSuggestions, 300), []);

  const onChange = (event, { newValue }) => {
    setValue(newValue);
  };

  const onSuggestionsFetchRequested = ({ value }) => {
    debouncedFetchSuggestions(value);
  };

  const onSuggestionsClearRequested = () => {
    setSuggestions([]);
  };

  const getSuggestionValue = (suggestion) => suggestion.name;

  const onSuggestionSelected = (event, { suggestion }) => {
    navigate(`/courses/${suggestion.id}`); // Make sure the ID is passed correctly
  };

  const renderSuggestion = (suggestion) => (
    console.log('Rendering suggestion:', suggestion.name), // Debugging line
    <div className="px-4 py-2 text-black">
      {suggestion.name}
    </div>
  );

  const inputProps = {
    placeholder: 'Type a course name',
    value,
    onChange,
    className: 'w-full px-4 py-2 text-white rounded-lg border border-gray-300 bg-gray-800 placeholder-gray-400'
  };

  return (
    <div className="flex justify-center items-center h-screen ">
    <div className="p-8 rounded-lg shadow-lg max-w-md w-full bg-gray-800">
      <h2 className="text-2xl font-bold text-white mb-4">Search Courses</h2>
      <Autosuggest
        suggestions={suggestions}
        onSuggestionsFetchRequested={onSuggestionsFetchRequested}
        onSuggestionsClearRequested={onSuggestionsClearRequested}
        getSuggestionValue={getSuggestionValue}
        renderSuggestion={renderSuggestion}
        inputProps={inputProps}
        onSuggestionSelected={onSuggestionSelected} // Add handler for suggestion selection
        theme={{
          container: 'relative',
          input: 'border p-2 w-full rounded-lg',
          suggestionsContainer: suggestions.length === 0 ? 'hidden' : 'absolute mt-1 w-full bg-white border border-gray-300 rounded-lg z-10',
          suggestion: 'block px-4 py-2 cursor-pointer',
          suggestionHighlighted: 'bg-gray-200',
        }}
      />
    </div>
  </div>
  
  );
};

export default SearchComponent;
