import React, { useState, useCallback } from 'react';
import Autosuggest from 'react-autosuggest';
import axios from 'axios';
import debounce from 'lodash.debounce';
import { useNavigate } from 'react-router-dom';
import Alert from './Alert';

const SearchComponent = () => {
  const [value, setValue] = useState('');
  const [suggestions, setSuggestions] = useState([]);
  const navigate = useNavigate();
  const [showAlert, setShowAlert] = useState(false);
  const [errorMessage, setErrorMessage] = useState('');
  const [alertType, setAlertType] = useState(null);

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
      const response = await axios.get('/backend/courses');
      const courses = response.data;
      console.log('Fetched courses:', courses);
  
      const filteredSuggestions = courses.map(course => {
        const courseName = course.name.toLowerCase();
        const courseDescription = course.description.toLowerCase();
        const isNameMatch = courseName.startsWith(trimmedValue);
        const isDescriptionMatch = courseDescription.includes(trimmedValue);
  
        // Add source field to differentiate
        return {
          ...course,
          source: isNameMatch ? 'name' : isDescriptionMatch ? 'description' : ''
        };
      }).filter(suggestion => suggestion.source); // Filter out suggestions without a source
  
      setSuggestions(filteredSuggestions);
    } catch (error) {
      setErrorMessage('Failed to fetch courses');
      setAlertType('error');
      setShowAlert(true);
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
    navigate(`/courses/${suggestion.id}`);
  };

  const onKeyDown = (event) => {
    if (event.key === 'Enter' && suggestions.length === 0) {
      setErrorMessage('No matches found');
      setAlertType('warning');
      setShowAlert(true);
    }
  };

  const renderSuggestion = (suggestion) => (
    <div className="px-4 py-2 text-black">
      <div>{suggestion.name}</div>
  
      <small className="text-gray-500">Source: {suggestion.source}</small> {/* Add source indication */}
    </div>
  );

  const inputProps = {
    placeholder: 'Type a course name',
    value,
    onChange,
    onKeyDown,
    className: 'w-full px-4 py-2 text-white rounded-lg border border-gray-300 bg-gray-800 placeholder-gray-400'
  };

  return (
    <div className="flex justify-center items-center h-[50vh]">
      {showAlert && <Alert message={errorMessage} type={"error"} onClose={() => setShowAlert(false)} />}
      <div className="p-1 bg-gradient-to-r from-cyan-400 via-yellow-500 to-pink-500 rounded-lg shadow-lg max-w-md w-full">
        <div className="p-8 rounded-lg shadow-lg max-w-md w-full bg-gray-800">
          <h2 className="text-2xl font-bold text-white mb-4">Search Courses</h2>
          <Autosuggest
            suggestions={suggestions}
            onSuggestionsFetchRequested={onSuggestionsFetchRequested}
            onSuggestionsClearRequested={onSuggestionsClearRequested}
            getSuggestionValue={getSuggestionValue}
            renderSuggestion={renderSuggestion}
            inputProps={inputProps}
            onSuggestionSelected={onSuggestionSelected}
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
    </div>
  );
};

export default SearchComponent;
