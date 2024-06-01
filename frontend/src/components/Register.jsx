import { useState } from "react";
import Button from "./Button";
import InputField from "./Input";
import Section from "./Section";
import Alert from "./Alert";


const Register = () => {
  const [name, setName] = useState('');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [email, setEmail] = useState('')
  const [registerFailed, setRegisterFailed] = useState(false);
  const [registerSuccess, setRegisterSuccess] = useState(false);
  const [responseData, setResponseData] = useState({});

  const handleSubmit = async (e) => {
    e.preventDefault()

    const response = await fetch("http://localhost:8080/users", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({name,username,password,email}),
    });

    if (response.ok){
      const data = await response.json();
      console.log("Registered succesfully", data);
      setRegisterFailed(false);
      setRegisterSuccess(true);
      // Handle succesful register.
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
    <Section
      className="-mt-[5.25rem]"
      customPaddings>
      <div className="flex justify-center items-center h-screen ">
        <div className="p-8 rounded-lg shadow-lg max-w-md w-full bg-gray-800">
          <form className="space-y-4 font-semibold" onSubmit={handleSubmit}>
            <div>
              <label htmlFor="name" className="block text-withe">Name:</label>
              <InputField type="text" id="name" name="name" className="w-full" value={name} onChange={(e) => setName(e.target.value)}/>
            </div>
            <div>
              <label htmlFor="username" className="block text-withe">Username:</label>
              <InputField type="text" id="username" name="username" className="w-full" value={username} onChange={(e) => setUsername(e.target.value)}/>
            </div>
            <div>
              <label htmlFor="password" className="block text-withe">Password:</label>
              <InputField type="password" id="password" name="password" className="w-full" value={password} onChange={(e) => setPassword(e.target.value)}/>
            </div>
            <div>
              <label htmlFor="email" className="block text-withe">Email:</label>
              <InputField type="text" id="email" name="email" className="w-full" value={email} onChange={(e) => setEmail(e.target.value)}/>
            </div>
            <div className="rounded bg-gray-800">
              <Button type="submit" className="w-full bg-gray-800 text-white hover:bg-gray-800 rounded text-2xl font-semibold">Register</Button>
            </div>
          </form>
          <div>
              {registerFailed && <Alert message={`Register failed: ${responseData}`} type="error" onClose={() => setRegisterFailed(false)}/>}
              {registerSuccess && <Alert message="Registered successfully!" type="success" onClose={() => setRegisterSuccess(false)}/>}
            </div>
        </div>
      </div>
    </Section>
  );
};

export default Register;
