import { useState, useEffect } from "react";
import Select from "react-select";
import Button from "./Button";
import InputField from "./Input";
import Section from "./Section";

const CourseCreation = () => {
    const [name, setName] = useState('');
    const [description, setDescription] = useState('');
    const [price, setPrice] = useState(0.00);
    const [startDate, setStartDate] = useState(() => {
        const date = new Date();
        const day = String(date.getDate()).padStart(2, '0');
        const month = String(date.getMonth() + 1).padStart(2, '0'); // Months are 0-based in JavaScript
        const year = date.getFullYear();
        return `${year}-${month}-${day}`; // Use YYYY-MM-DD format for date input
    });
    const [endDate, setEndDate] = useState("");
    const [image, setImage] = useState(null);
    const [categories, setCategories] = useState([]);
    const [selectedCategories, setSelectedCategories] = useState([]);

    useEffect(() => {
        const fetchCategories = async () => {
            const response = await fetch("http://localhost:8080/category/all");
            const data = await response.json();
            const formattedData = data.map(category => ({
                value: category.id,
                label: category.name
            }));
            setCategories(formattedData);
        };
        fetchCategories();
    }, []);

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
        e.preventDefault();
        const token = localStorage.getItem("token");

        if (!name || !description || !price || !startDate || !endDate || !image || selectedCategories.length === 0) {
            alert("Please fill in all fields.");
            return;
        }

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
                start_date: startDate,
                end_date: endDate,
                categories_id: selectedCategories.map(category => category.value),
                id_owner: 1, // Replace with actual owner ID if necessary
            }),
        });

        if (response.ok) {
            const data = await response.json();
            console.log("Course created successfully", data);
            // Handle successful register.
        } else {
            const errorData = await response.json();
            console.error("Course creation failed", errorData);
            // Handle register failure
        }
    };

    return (
        <Section className="-mt-[5.25rem]" customPaddings>
            <div className="flex justify-center items-center h-screen ">
                <div className="p-8 rounded-lg shadow-lg max-w-md w-full bg-gray-800">
                    <form className="space-y-4 font-semibold" onSubmit={handleSubmit}>
                        <div>
                            <label htmlFor="name" className="block text-white">Name:</label>
                            <InputField type="text" id="name" name="name" className="w-full" placeholder="Course name" value={name} onChange={(e) => setName(e.target.value)} />
                        </div>
                        <div>
                            <label htmlFor="description" className="block text-white">Description:</label>
                            <InputField type="text" id="description" name="description" className="w-full" placeholder="Description" value={description} onChange={(e) => setDescription(e.target.value)} />
                        </div>
                        <div>
                            <label htmlFor="price" className="block text-white">Price:</label>
                            <InputField type="number" id="price" name="price" className="w-full" placeholder="0.00" value={price} onChange={(e) => setPrice(e.target.value)} />
                        </div>
                        <div>
                            <label htmlFor="start_date" className="block text-white">Start Date:</label>
                            <InputField type="date" id="start_date" name="start_date" className="w-full" value={startDate} onChange={(e) => setStartDate(e.target.value)} />
                        </div>
                        <div>
                            <label htmlFor="end_date" className="block text-white">End Date:</label>
                            <InputField type="date" id="end_date" name="end_date" className="w-full" value={endDate} onChange={(e) => setEndDate(e.target.value)} />
                        </div>
                        <div>
                            <label htmlFor="categories" className="block text-white">Categories:</label>
                            <Select
                                id="categories"
                                name="categories"
                                isMulti
                                options={categories}
                                className="basic-multi-select"
                                classNamePrefix="select"
                                onChange={setSelectedCategories}
                            />
                        </div>
                        <div>
                            <label htmlFor="image" className="block text-white">Image:</label>
                            <input type="file" id="image" name="image" onChange={handleImageChange} />
                        </div>
                        <Button type="submit" className="w-full bg-gray-800 text-white hover:bg-gray-800 rounded text-2xl font-semibold">Create Course</Button>
                    </form>
                </div>
            </div>
        </Section>
    );
};

export default CourseCreation;