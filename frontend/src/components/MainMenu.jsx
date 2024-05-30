import React from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
import { Section } from './Section';

const MainMenu = ({ user, actions }) => {
  return (
    <Section>
    <div className="main-menu">
      <span>Welcome, {user.name}!</span>
      <nav>
        <ul>
          {actions.map(action => (
            <li key={action.id}>
              <Link to={action.url}>{action.title}</Link>
            </li>
          ))}
        </ul>
      </nav>
    </div></Section>
    
  );
};

MainMenu.propTypes = {
  user: PropTypes.shape({
    name: PropTypes.string.isRequired,
  }).isRequired,
  actions: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.string.isRequired,
      url: PropTypes.string.isRequired,
      title: PropTypes.string.isRequired,
    })
  ).isRequired,
};

export default MainMenu;
