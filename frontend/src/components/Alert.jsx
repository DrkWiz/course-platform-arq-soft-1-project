import React, { useEffect }  from 'react';
import './Alert.css'; // Asegúrate de que la ruta de importación sea correcta
import Ding from "../assets/Ding.mp3"

const Alert = ({ message, type, onClose }) => {
  const alertType = type === 'error' ? 'bg-red-500' : 'bg-green-500';

  useEffect(() => {
    const audio = new Audio(Ding);
    audio.play();
  }, []);

  return (
    <div className={`alert text-white px-6 py-4 border-0 rounded ${alertType}`}>
      <span className="text-xl inline-block mr-5 align-middle">
        <i className="fas fa-bell" />
      </span>
      <span className="inline-block align-middle mr-8">
        {message}
      </span>
      <button className="close-button"
      onClick={onClose}>
        <span>x</span>
      </button>
    </div>
  );
};

export default Alert;