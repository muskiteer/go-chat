import { create } from "zustand";
import { toast } from "react-hot-toast"; // Changed from default import to named import
import { axiosInstance } from "../lib/axios";
import { useAuthStore } from "./useAuthStore";
import { use } from "react";






export const useChatStore = create((set, get) => ({
  messages: [],
  users: [],
  selectedUser: null,
  isUsersLoading: false,
  isMessagesLoading: false,

  getUsers: async () => {
    set({ isUsersLoading: true });
    try {
      const res = await axiosInstance.get("/users");
      set({ users: res.data });
    } catch (error) {
      toast.error(error.response.data.message);
    } finally {
      set({ isUsersLoading: false });
    }
  },

  getMessages: async (userId) => {
    set({ isMessagesLoading: true });

    try {
      const res = await axiosInstance.get(`/${userId}`);
      const data = res.data;

      // Set messages regardless of whether they're empty or not
      set({ messages: data || [] });
     

      // Show friendly message only for empty results
      

    } catch (error) {
      console.error("Error loading messages:", error);
      set({ messages: [] });
      
      // Only show error for actual network/server errors
      toast.error("Failed to load messages. Please try again.");
    } finally {
      set({ isMessagesLoading: false });
    }
  },

  sendMessage: async (messageData) => {
    const { selectedUser, messages } = get();     
    try {
      // console.log({messageData})
      const res = await axiosInstance.post(`/send/${selectedUser._id}`, messageData);
      set({ messages: [...messages, res.data] });
    } catch (error) {
      toast.error(error.response.data.message);
    }
  },

  setSelectedUser: (selectedUser) => set({ selectedUser }),

  subscribeToMessages: () => {
    // console.log("hello");
    const { selectedUser } = get();

  if (!selectedUser) return;
  const socket = useAuthStore.getState().socket;

  // console.log("hello again" )
  const selectedUserId = selectedUser._id;

  

  const handler = (event) => {
    const payload = JSON.parse(event.data);

    // console.log("WebSocket message received:", payload);

    if (payload.event === "getOnlineUsers") {
      // console.log("Online users received:", payload);
      
      set({ onlineUsers: payload.data });
    }

    if (payload.event === "newMessage") {
      console.log("New message event received", payload);
      const newMessage = payload.data;
      const isFromSelectedUser = newMessage.sender_id === selectedUserId;
      console.log("New message received:", { newMessage });
      // console.log({selectedUser});

      if (isFromSelectedUser) {
        set((state) => ({
          messages: [...state.messages, newMessage],
        }));
      }
    }
  };

  socket.addEventListener("message", handler);

  // Return unsubscribe function
  return () => {
        socket.removeEventListener("message", handler);
  };
},



}));