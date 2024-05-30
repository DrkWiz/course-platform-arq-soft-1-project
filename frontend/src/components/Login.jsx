import { useState } from "react";
import Button from "./Button";
import InputField from "./Input";
import Section from "./Section";

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    const response = await fetch("http://localhost:8080/users/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ username, password }),
    });

    if (response.ok) {
      const data = await response.json();
      console.log("Login successful", data);
      // Handle successful login, e.g., store the token, redirect user, etc.
    } else {
      const errorData = await response.json();
      console.error("Login failed", errorData);
      // Handle login failure, e.g., show an error message
    }
  };

  return (
    <Section className="-mt-[5.25rem]" customPaddings>
      <div className="flex justify-center items-center h-screen">
        <div className="p-8 rounded-lg shadow-lg max-w-md w-full bg-gray-800">
          <form className="space-y-4 font-semibold" onSubmit={handleSubmit}>
            <div>
              <label htmlFor="username" className="block text-withe">Username:</label>
              <InputField
                type="text"
                id="username"
                name="username"
                className="w-full"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
              />
            </div>
            <div>
              <label htmlFor="password" className="block text-withe">Password:</label>
              <InputField
                type="password"
                id="password"
                name="password"
                className="w-full"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </div>
            <div className="rounded bg-gray-800">
              <Button
                type="submit"
                className="w-full bg-gray-800 text-white hover:bg-gray-800 rounded text-2xl font-semibold"
              >
                Login
              </Button>
            </div>
          </form>
        </div>
      </div>
    </Section>
  );
};

export default Login;
