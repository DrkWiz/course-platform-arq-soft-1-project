import { useState } from "react";
import Button from "./Button";
import InputField from "./Input";
import Section from "./Section";
import Alert from "./Alert";
import { useNavigate } from 'react-router-dom';

const Register = () => {
  const [name, setName] = useState('');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [email, setEmail] = useState('')
  const [registerFailed, setRegisterFailed] = useState(false);
  const [registerSuccess, setRegisterSuccess] = useState(false);
  const [responseData, setResponseData] = useState({});
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault()

    const response = await fetch("http://localhost:8080/users", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({name, username, password, email}),
    });

    if (response.ok) {
      const data = await response.json();
      console.log("Registered successfully", data);
      setRegisterFailed(false);
      setRegisterSuccess(true);
      setTimeout(() => {
        navigate('/login');
      }, 1500);
      // Handle successful register.
    } else {
      const errorData = await response.json();
      console.error("Register failed", errorData);
      setRegisterSuccess(false);
      setRegisterFailed(true);
      setResponseData(errorData.message);
      // Handle register failure
    }
  }

  return (
    <Section className="-mt-[3rem]" customPaddings>
      <div className="flex justify-center items-center h-screen">
        <div className="p-1 bg-gradient-to-r from-cyan-400 via-yellow-500 to-pink-500 rounded-lg shadow-lg max-w-md w-full">
          <div className="p-8 rounded-lg shadow-lg max-w-md w-full bg-gray-800">
            <form className="space-y-4 font-semibold" onSubmit={handleSubmit}>
              <div>
                <label htmlFor="name" className="block text-white">Name:</label>
                <InputField type="text" id="name" name="name" className="w-full" value={name} onChange={(e) => setName(e.target.value)} required />
              </div>
              <div>
                <label htmlFor="username" className="block text-white">Username:</label>
                <InputField type="text" id="username" name="username" className="w-full" value={username} onChange={(e) => setUsername(e.target.value)} required />
              </div>
              <div>
                <label htmlFor="password" className="block text-white">Password:</label>
                <InputField type="password" id="password" name="password" className="w-full" value={password} onChange={(e) => setPassword(e.target.value)} required />
              </div>
              <div>
                <label htmlFor="email" className="block text-white">Email:</label>
                <InputField type="email" id="email" name="email" className="w-full" value={email} onChange={(e) => setEmail(e.target.value)} required />
              </div>
              <div className="rounded bg-gray-800">
                <Button type="submit" className="w-full bg-gray-800 text-white hover:bg-gray-800 rounded text-2xl font-semibold">Register</Button>
              </div>
            </form>
            <div>
              {registerFailed && <Alert message={`Register failed: ${responseData}`} type="error" onClose={() => setRegisterFailed(false)} />}
              {registerSuccess && <Alert message="Registered successfully!" type="success" onClose={() => setRegisterSuccess(false)} />}
            </div>
          </div>
        </div>
      </div>
    </Section>
  );
};

export default Register;