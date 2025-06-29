import Navbar  from "./components/Navbar"
import { BrowserRouter } from "react-router-dom";


import { Routes, Route } from "react-router-dom";
import HomePage from "./pages/HomePage";  
import SignUpPage from "./pages/SignUpPage";
import SettingsPage from "./pages/SettingsPage";
import LoginPage from "./pages/LoginPage";  


import { useAuthStore } from "./store/useAuthStore";
import { useEffect } from "react";
import { Loader } from "lucide-react";




const App = () => {

    const {authUser,checkAuth,isCheckingAuth} = useAuthStore()
  useEffect(() => {
    checkAuth();
   },[checkAuth]);

     console.log({ authUser });

     if (isCheckingAuth && !authUser) {
        return (
        <div className="flex items-center justify-center h-screen">
          <Loader className="size-10 animate-spin"/>

        </div>
        );
     }

     

  return (
    <BrowserRouter>
    
    <div>
      <Navbar />
     <Routes>
      <Route path="/" element={authUser ? <HomePage /> : <Navigate to="/login" />}/>
      <Route path="/signup" element={!authUser ? <SignUpPage /> : <Navigate to="/" />} />
      <Route path="/login" element={!authUser ? <LoginPage /> : <Navigate to="/" />}/>
      <Route path="/settings" element={<SettingsPage/>} />
     </Routes>

      
    </div>
    </BrowserRouter>
  )
}

export default App