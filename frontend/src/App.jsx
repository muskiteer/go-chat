import Navbar  from "./components/Navbar"
import { BrowserRouter } from "react-router-dom";

import { Routes, Route } from "react-router-dom";
import HomePage from "./pages/HomePage";  
import SignUpPage from "./pages/SignUpPage";
import SettingsPage from "./pages/SettingsPage";
import LoginPage from "./pages/LoginPage";  



const App = () => {
  return (
    <div >

      <Navbar />
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/signup" element={<SignUpPage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/settings" element={<SettingsPage />} />
       

      </Routes>
    </BrowserRouter>


    
    </div>
  )
}

export default App