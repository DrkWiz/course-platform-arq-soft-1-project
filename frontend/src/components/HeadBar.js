import React from 'react';
import './HeadBar.css';

const HeadBar = ({ logo, buttons, onLogoClick }) => (
  <nav className="navbar">
    <img src={logo} className="logo" onClick={onLogoClick} alt="logo" />
    <div className="buttons">
      {buttons.map((button, index) => (
        <button key={index} onClick={button.onClick}>
          {button.label}
        </button>
      ))}
    </div>
  </nav>
);

export default HeadBar;