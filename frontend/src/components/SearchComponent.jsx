import React, { Component } from 'react';
import Autosuggest from 'react-autosuggest';
import axios from 'axios';

class SearchComponent extends Component {
  constructor() {
    super();
    this.state = {
      value: '',
      suggestions: []
    };
  }

  onChange = (event, { newValue }) => {
    this.setState({
      value: newValue
    });
  };

  onSuggestionsFetchRequested = ({ value }) => {
    this.getSuggestions(value);
  };

  onSuggestionsClearRequested = () => {
    this.setState({
      suggestions: []
    });
  };

  getSuggestions = (value) => {
    const inputValue = value.trim().toLowerCase();
    const inputLength = inputValue.length;

    if (inputLength === 0) return;

    axios.get('http://localhost:8080/courses')
      .then(response => {
        const courses = response.data;
        this.setState({
          suggestions: courses.filter(course =>
            course.Name.toLowerCase().slice(0, inputLength) === inputValue
          )
        });
      })
      .catch(error => {
        console.error("Error fetching courses", error);
      });
  };

  getSuggestionValue = (suggestion) => suggestion.Name;

  renderSuggestion = (suggestion) => (
    <div>
      {suggestion.Name}
    </div>
  );

  render() {
    const { value, suggestions } = this.state;

    const inputProps = {
      placeholder: 'Type a course name',
      value,
      onChange: this.onChange
    };

    return (
      <Autosuggest
        suggestions={suggestions}
        onSuggestionsFetchRequested={this.onSuggestionsFetchRequested}
        onSuggestionsClearRequested={this.onSuggestionsClearRequested}
        getSuggestionValue={this.getSuggestionValue}
        renderSuggestion={this.renderSuggestion}
        inputProps={inputProps}
      />
    );
  }
}

export default SearchComponent;