import { useState } from "react";
import Button from "./Button";
import InputField from "./Input";
import Section from "./Section";

const CourseCreation = () => {
    const [name, setName] = useState('');
    const [description, setDescription] = useState('');
    
    //const [category, setCategory] = useState('');
    const [price, setPrice] = useState(0.00);
    const [start_date, setStartDate] = useState(() => {
        const date = new Date();
        const day = String(date.getDate()).padStart(2, '0');
        const month = String(date.getMonth() + 1).padStart(2, '0'); // Months are 0-based in JavaScript
        const year = date.getFullYear();
      
        return `${day}/${month}/${year}`;
      });
    const [end_date, setEndDate] = useState("");
    const [image, setImage] = useState(null);
    

const handleImageChange = (e) => {
  setImage(e.target.files[0]);
};

const handleImageUpload = async () => {
    const token = localStorage.getItem("token");
    const formData = new FormData();
    formData.append('image', image);
  
    const response = await fetch('http://localhost:8080/upload', {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
      body: formData,
    });
  
    const data = await response.json();
    return data.picture_path;
  };

    const handleSubmit = async (e) => {
        e.preventDefault()
        const token = localStorage.getItem("token");
        const picturePath = await handleImageUpload();
        const response = await fetch("http://localhost:8080/courses", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                'Authorization': `Bearer ${token}`, 
            },
            body: JSON.stringify({
                name: name,
                description: description,
                price: parseFloat(price),
                picture_path: picturePath,
                start_date: start_date,
                end_date: end_date,
              }),
        });

        if (response.ok) {
            const data = await response.json();
            console.log("Course created succesfully", data);
            // Handle succesful register.
        } else {
            const errorData = await response.json();
            console.error("Course creation failed", errorData);
            // Handle register failure
        }
    }

    return (
        <Section
            className="-mt-[5.25rem]"
            customPaddings>
            <div className="flex justify-center items-center h-screen ">
                <div className="p-8 rounded-lg shadow-lg max-w-md w-full bg-gray-800">
                    <form className="space-y-4 font-semibold" onSubmit={handleSubmit}>
                        <div>
                            <label htmlFor="name" className="block text-withe">Name:</label>
                            <InputField type="text" id="name" name="name" className="w-full" placeholder={"Course name"}value={name} onChange={(e) => setName(e.target.value)} />
                        </div>
                        <div>
                            <label htmlFor="description" className="block text-withe">Description:</label>
                            <InputField type="text" id="description" name="description" className="w-full" placeholder={"Description"} value={description} onChange={(e) => setDescription(e.target.value)} />
                        </div>
                        <div>
                            <label htmlFor="price" className="block text-withe">Price:</label>
                            <InputField type="text" id="price" name="price" className="w-full" placeholder={"0.00"}value={price} onChange={(e) => setPrice(e.target.value)} />
                        </div>
                        <div>
                            <label htmlFor="end_date" className="block text-withe">End Date:</label>
                            <InputField type="text" id="end_date" name="end_date" className="w-full" placeholder={"DD/MM/YYYY"} value={end_date} onChange={(e) => setEndDate(e.target.value)} />
                        </div>
                        <div className="rounded bg-gray-800">
                        <div>

            <label htmlFor="image" className="block text-withe">Image:</label>
            <input type="file" id="image" name="image" onChange={handleImageChange} />
          </div>
                            <Button type="submit" className="w-full bg-gray-800 text-white hover:bg-gray-800 rounded text-2xl font-semibold">Create Course</Button>
                        </div>
                    </form>
                </div>
            </div>
        </Section>
    );

};

export default CourseCreation;