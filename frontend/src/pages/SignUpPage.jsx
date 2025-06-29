import React from 'react'
import {useState} from "react";


const SignUpPage = () => {
  const [showPassword, setShowPassword] = useState(false);
  const [formData,setFormData] = useState({
    fullName:"",
    email: "",
    password:"",
  }); 

  const {signup,isSignUp} = useAuthStore();

  const validateForm =()=>{}
  const handleSubmit =()=>{
    e.preventDefault()
  }


  return (
    <div className="nim-h-screen grid lg:grid-cols-2">
      <div className="flex flex-col justify-center items-center p-6 sm:p-12">

        <div className="w-full max-w-md space-y-8">

        </div>

      </div>

    </div>
  )
}

export default SignUpPage