import axios from 'axios';
import {create} from 'zustand';
import { axiosInstance } from '../lib/axios';
import { toast } from 'react-hot-toast';
import { useChatStore } from './useChatStore';

export const useAuthStore = create((set, get) => ({
    authUser: null,
    isSigningUp: false,
    isLoggingIn: false,
    onlineUsers: [],
    socket: null,

    isCheckingAuth: true,

    checkAuth: async () => {
        try {
            const res = await axiosInstance.get('/check');
            set({ authUser: res.data.user || res.data }); // Extract user if it exists, otherwise use res.data
            get().connectSocket();
        } catch (error) {
            console.log('Error checking auth:', error); 
            set({ authUser: null });
        } finally {
            set({ isCheckingAuth: false });
        }
    },

    signup: async (data) => {
    set({ isSigningUp: true });
    try {
      const res = await axiosInstance.post("/signup", data);
      set({ authUser: res.data.user || res.data }); // Extract user if it exists
      toast.success("Account created successfully");
      get().connectSocket();
    } catch (error) {
      toast.error(error.response.data.message || "Username or email already exists");
    } finally {
      set({ isSigningUp: false });
    }
  },

  logout: async () => {
    try {
        await axiosInstance.post('/logout');
        set({authUser: null});
        toast.success("Logged out successfully");

        get().disconnectSocket();
    } catch (error) {
        toast.error(error.response.data.message || "Logout failed");
        
    }
  },
   login: async (data) => {
    set({ isLoggingIn: true });
    try {
      const res = await axiosInstance.post("/login", data);
      set({ authUser: res.data.user || res.data }); // Extract user if it exists
      toast.success("Logged in successfully");

      get().connectSocket();
    } catch (error) {
  const message =
    error.response?.data?.message?.trim() ||
    (error.response?.status === 401
      ? "Invalid email or password"
      : error.message) ||
    "An unexpected error occurred";

  toast.error(message);
}
 finally {
      set({ isLoggingIn: false });
    }
  },
  connectSocket: () => {
    const { authUser, socket } = get();
    if (!authUser || socket) return;

    const ws = new WebSocket(`ws://localhost:8080/ws?userId=${authUser.userId}`);

    ws.onopen = () => {
      console.log("WebSocket connected");
    };

    ws.onmessage = (event) => {
  const payload = JSON.parse(event.data);
  if (payload.event === "getOnlineUsers") {
    set({ onlineUsers: payload.data });
  }
};
    


    ws.onclose = () => {
      console.log("WebSocket disconnected");
      set({ socket: null, onlineUsers: [] });
    };

    set({ socket: ws });
 
  },

  disconnectSocket: () => {
    const { socket } = get();
    if (socket) {
      socket.close();
      set({ socket: null, onlineUsers: [] });
    }
  },

}));