import { Route, Routes } from "react-router-dom";
import ButtonGradient from './assets/svg/ButtonGradient';
import Benefits from './components/Benefits';
import Footer from './components/Footer';
import Header from './components/Header';
import Hero from './components/Hero';
import Pricing from './components/Pricing';
import Login from './components/Login';

const App = () => {
  return (
    <div className="pt-[4.75rem] lg:pt-[5.25rem] overflow-hidden">
      <Header />
      <Routes>
        <Route path="/" element={
          <>
            <Hero />
            <Benefits />
            <Pricing />
          </>
        } />
        <Route path="/login" element={<Login />} />
      </Routes>
      <Footer />
      <ButtonGradient />
    </div>
  );
};

export default App;