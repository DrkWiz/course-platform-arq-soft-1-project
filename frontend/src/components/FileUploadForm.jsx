import React, { useState } from 'react';

const UploadFile = ({ courseId }) => {
  const [file, setFile] = useState(null);

  const handleFileChange = (event) => {
    setFile(event.target.files[0]);
  };

  const handleUpload = async () => {
    if (!file) {
      alert("Please select a file first!");
      return;
    }

    const formData = new FormData();
    formData.append("file", file);

    try {
      const response = await fetch(`/backend/courses/${courseId}/files`, {
        method: "POST",
        headers: {
          "Authorization": `Bearer ${localStorage.getItem("token")}`
        },
        body: formData
      });

      if (response.ok) {
        alert("File uploaded successfully");
      } else {
        const errorData = await response.json();
        alert(`Failed to upload file: ${errorData.error}`);
      }
    } catch (error) {
      alert("Error uploading file: " + error);
    }
  };

  return (
    <div>
      <input type="file" onChange={handleFileChange} />
      <button onClick={handleUpload}>Upload File</button>
    </div>
  );
};

export default UploadFile;
