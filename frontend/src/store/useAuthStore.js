import axios from 'axios';
import {create} from 'zustand';
import { axiosInstance } from '../lib/axios';
import { toast } from 'react-hot-toast';


export const useAuthStore = create((set) => ({
    authUser: null,
    isSigningUp: false,
    isLoggingIn: false,

    isCheckingAuth: true,

    checkAuth: async () =>{
        try {
            const res = await axiosInstance.get('/check');

            set ({authUser: res.data});
        } catch (error) {
            console.log('Error checking auth:', error); 
            set({authUser: null});
        } finally {
            set({isCheckingAuth: false});
        }
    },

    signup: async (data) => {
    set({ isSigningUp: true });
    try {
      const res = await axiosInstance.post("/signup", data);
      set({ authUser: res.data });
      toast.success("Account created successfully");
      get().connectSocket();
    } catch (error) {
      toast.error(error.response.data.message);
    } finally {
      set({ isSigningUp: false });
    }
  },

}));