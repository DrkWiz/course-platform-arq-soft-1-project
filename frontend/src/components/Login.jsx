import Button from "./Button";
import InputField from "./Input";
import Section from "./Section";

const Login = () => {
  return (
    <Section
      className="-mt-[5.25rem]"
      customPaddings>
      <div className="flex justify-center items-center h-screen ">
        <div className="p-8 rounded-lg shadow-lg max-w-md w-full bg-gray-800">
          <form className="space-y-4 font-semibold">
            <div>
              <label htmlFor="username" className="block text-withe">Username:</label>
              <InputField type="text" id="username" name="username" className="w-full" />
            </div>
            <div>
              <label htmlFor="password" className="block text-withe">Password:</label>
              <InputField type="password" id="password" name="password" className="w-full" />
            </div>
            <div className="rounded bg-gray-800">
              <Button type="submit" className="w-full bg-gray-800 text-white hover:bg-gray-800 rounded text-2xl font-semibold">Login</Button>
            </div>
          </form>
        </div>
      </div>
    </Section>
  );
};

export default Login;
