import Button from "./Button";
import InputField from "./Input";
import Section from "./Section";

const Login = () => {
  return (
    <Section
      className="pt-[12rem] -mt-[5.25rem]"
      crosses
      crossesOffset="lg:translate-y-[5.25rem]"
      customPaddings>
      <div className="flex justify-center items-center h-screen">
        <div className="p-8 rounded-lg shadow-lg max-w-md w-full bg-gray-100">
          <h2 className="text-2xl font-semibold text-gray-800 mb-4">Login</h2>
          <form className="space-y-4">
            <div>
              <label htmlFor="username" className="block text-gray-800">Username:</label>
              <InputField type="text" id="username" name="username" className="w-full" />
            </div>
            <div>
              <label htmlFor="password" className="block text-gray-800">Password:</label>
              <InputField type="password" id="password" name="password" className="w-full" />
            </div>
            <div className="rounded bg-gray-800">
              <Button type="submit" className="w-full bg-gray-800 text-white hover:bg-gray-800 rounded">Login</Button>
            </div>
          </form>
        </div>
      </div>
    </Section>
  );
};

export default Login;
