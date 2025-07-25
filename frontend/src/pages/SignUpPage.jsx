import React from 'react'
import {useState} from "react";
import { Eye, EyeOff, Loader2, Lock, Mail, MessageSquare, User } from "lucide-react";
import {useAuthStore} from "./../store/useAuthStore.js";
import { Link } from "react-router-dom";
import AuthImagePattern from './../components/AuthImagePattern.jsx';
import { toast } from 'react-hot-toast';




const SignUpPage = () => {
  const [showPassword, setShowPassword] = useState(false);
  const [formData,setFormData] = useState({
    username:"",
    email: "",
    hashed_password:"",
  }); 

  const {signup,isSigningUp} = useAuthStore();

   const validateForm = () => {
  if (!formData.username.trim()) {
    toast.error("Full name is required");
    return false;
  }
  if (!formData.email.trim()) {
    toast.error("Email is required");
    return false;
  }
  if (!/\S+@\S+\.\S+/.test(formData.email)) {
    toast.error("Invalid email format");
    return false;
  }
  if (!formData.hashed_password) {
    toast.error("Password is required");
    return false;
  }
  if (formData.hashed_password.length < 6) {
    toast.error("Password must be at least 6 characters");
    return false;
  }

  return true;
};

 const handleSubmit = (e) => {
  e.preventDefault();

  if (!validateForm()) return;

  signup(formData);
};



  return (
    
    <div className="min-h-screen grid lg:grid-cols-2">
      {/* left side */}
      <div className="flex flex-col justify-center items-center p-6 sm:p-12">
        <div className="w-full max-w-md space-y-8">
          {/* LOGO */}
          <div className="text-center mb-8">
            <div className="flex flex-col items-center gap-2 group">
              <div
                className="size-12 rounded-xl bg-primary/10 flex items-center justify-center 
              group-hover:bg-primary/20 transition-colors"
              >
                <MessageSquare className="size-6 text-primary" />
              </div>
              <h1 className="text-2xl font-bold mt-Loader22">Create Account</h1>
              <p className="text-base-content/60">Get started with your free account</p>
            </div>
          </div>

          <form onSubmit={handleSubmit} className="space-y-6">
            <div className="form-control">
              <label className="label">
                <span className="label-text font-medium">userName</span>
              </label>
              <div className="relative">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <User className="w-5 h-5 text-gray-400 z-10 visible" />
                </div>
                <input
                  type="text"
                  className={`input input-bordered w-full pl-10`}
                  placeholder="John Doe"
                  value={formData.username}
                  onChange={(e) => setFormData({ ...formData, username: e.target.value })}
                />
              </div>
            </div>

            <div className="form-control">
              <label className="label">
                <span className="label-text font-medium">Email</span>
              </label>
              <div className="relative">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <Mail className="w-5 h-5 text-gray-400 z-10 visible" />
                </div>
                <input
                  type="email"
                  className={`input input-bordered w-full pl-10`}
                  placeholder="you@example.com"
                  value={formData.email}
                  onChange={(e) => setFormData({ ...formData, email: e.target.value })}
                />
              </div>
            </div>

            <div className="form-control">
              <label className="label">
                <span className="label-text font-medium">Password</span>
              </label>
              <div className="relative">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <Lock className="w-5 h-5 text-gray-400 z-10 visible" />
                </div>
                <input
                  type={showPassword ? "text" : "password"}
                  className={`input input-bordered w-full pl-10`}
                  placeholder="••••••••"
                  value={formData.hashed_password}
                  onChange={(e) => setFormData({ ...formData, hashed_password: e.target.value })}
                />
                <button
                  type="button"
                  className="absolute inset-y-0 right-0 pr-3 flex items-center"
                  onClick={() => setShowPassword(!showPassword)}
                >
                  {showPassword ? (
                    <EyeOff className="size-5 text-base-content/40" />
                  ) : (
                    <Eye className="size-5 text-base-content/40" />
                  )}
                </button>
              </div>
            </div>

            <button type="submit" className="btn btn-primary w-full" disabled={isSigningUp}>
              {isSigningUp ? (
                <>
                  <Loader2 className="size-5 animate-spin" />
                  Loading...
                </>
              ) : (
                "Create Account"
              )}
            </button>
          </form>

           <div className="text-center">
            <p className="text-base-content/60">
              Already have an account?{" "}
              <Link to="/login" className="link link-primary">
                Sign in
              </Link>
            </p>
          </div>

          </div>
          </div>
          

          <AuthImagePattern
          title="join our community"
          subtitle="Connect with everyone.."
          />

          </div>
  )
}

export default SignUpPage