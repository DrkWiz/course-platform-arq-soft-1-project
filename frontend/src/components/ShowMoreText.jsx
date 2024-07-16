import React, { useState } from 'react';

const ShowMoreText = ({ text, maxLength = 100 }) => {
  const [showFullText, setShowFullText] = useState(false);

  const toggleShowText = () => {
    setShowFullText((prev) => !prev);
  };

  if (text.length <= maxLength) {
    return <p>{text}</p>;
  }

  return (
    <p>
      {showFullText ? text : `${text.substring(0, maxLength)}...`}
      <button onClick={toggleShowText} className="text-blue-500 ml-2">
        {showFullText ? 'Show Less' : 'Show More'}
      </button>
    </p>
  );
};

export default ShowMoreText;
