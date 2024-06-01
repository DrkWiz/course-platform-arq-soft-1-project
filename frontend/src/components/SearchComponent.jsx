import React, { useState, useCallback } from 'react';
import Autosuggest from 'react-autosuggest';
import axios from 'axios';
import debounce from 'lodash.debounce';

const SearchComponent = () => {
  const [value, setValue] = useState('');
  const [suggestions, setSuggestions] = useState([]);

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

  const getSuggestionValue = (suggestion) => suggestion.Name;

  const renderSuggestion = (suggestion) => (
    
    console.log('Rendering suggestion:', suggestion.name), // Debugging line
    <div style={{ padding: '10px', color: 'black'}}>
      {suggestion.name}
    </div>
  );
  const inputProps = {
    placeholder: 'Type a course name',
    value,
    onChange
  };

  return (
    <Autosuggest
      suggestions={suggestions}
      onSuggestionsFetchRequested={onSuggestionsFetchRequested}
      onSuggestionsClearRequested={onSuggestionsClearRequested}
      getSuggestionValue={getSuggestionValue}
      renderSuggestion={renderSuggestion}
      
      inputProps={inputProps}
      theme={{
        container: 'relative',
        input: 'border p-2 w-full',
        suggestionsContainer: 'react-autosuggest__suggestions-container',
        suggestion: 'react-autosuggest__suggestion',
        suggestionHighlighted: 'react-autosuggest__suggestion--highlighted',
      }}
    />
  );
};

export default SearchComponent;
