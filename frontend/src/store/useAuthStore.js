import axios from 'axios';
import {create} from 'zustand';
import { axiosInstance } from '../lib/axios';


export const useAuthStore = create((set) => ({
    authUser: null,
    isSignningUp: false,
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
        
    }

}));