import { useChatStore } from "../store/useChatStore";
import { useEffect, useRef } from "react";
import { MessageCircle, Sparkles } from "lucide-react";
import avatar from "./../../assets/avatar.png";

import ChatHeader from "./ChatHeader";
import MessageInput from "./MessageInput.jsx";
import MessageSkeleton from "./skeletons/MessageSkeleton";
import { useAuthStore } from "../store/useAuthStore";
import { formatMessageTime } from "../lib/utils";
// import avatar from "../../assets/avatar.png";


const ChatContainer = () => {
  const {
    messages,
    getMessages,
    subscribeToMessages,
    
    isMessagesLoading,
      
    // socket
  } = useChatStore();
  const { socket } = useAuthStore();
  const { selectedUser } = useChatStore();
  const { authUser } = useAuthStore();
  const messageEndRef = useRef(null);

  // Add this debugging
  // console.log('authUser:', authUser);
  // console.log('authUser.id:', authUser.userId);

  useEffect(() => {
  if (!selectedUser?._id) return;
  getMessages(selectedUser._id);
}, [selectedUser?._id]);

  useEffect(() => {
    // console.log({socket});
  if (!selectedUser?._id || !socket) return;
    // console.log("Subscribing to messages for user:", selectedUser._id);
    // console.log(selectedUser._id)
  getMessages(selectedUser._id);
  const unsubscribe = subscribeToMessages();
  // console.log("subscribeToMessages called ✅");

  return () => {
    try {
      // console.log("Unsubscribing from messages for user:", selectedUser._id);
      unsubscribe();
    } catch (err) {
      console.error("Error during unsubscribe:", err);
    }
  };
}, [selectedUser?._id, socket]);// Simplified dependencies

  // console.log('selectedUser:', selectedUser);
  // console.log('selectedUser.id:', selectedUser.id);

  useEffect(() => {
    if (messageEndRef.current && messages) {
      messageEndRef.current.scrollIntoView({ behavior: "smooth" });
    }
  }, [messages]);

  if (isMessagesLoading) {
    return (
      <div className="flex-1 flex flex-col overflow-auto">
        <ChatHeader />
        <MessageSkeleton />
        <MessageInput />          
      </div>
    );
  }

   const EmptyMessageState = () => (
    <div className="flex-1 flex flex-col items-center justify-center p-8 text-center">
      <div className="relative mb-6">
        {/* Animated background circles */}
        <div className="absolute inset-0 animate-pulse">
          <div className="w-32 h-32 bg-primary/10 rounded-full animate-ping"></div>
        </div>
        <div className="relative z-10 w-24 h-24 bg-primary/20 rounded-full flex items-center justify-center">
          <MessageCircle className="w-12 h-12 text-primary animate-bounce" />
        </div>
      </div>
      
      <div className="space-y-4 max-w-md">
        <div className="flex items-center gap-2 justify-center">
          <Sparkles className="w-5 h-5 text-yellow-500 animate-pulse" />
          <h3 className="text-xl font-semibold text-base-content">
            Start the conversation!
          </h3>
          <Sparkles className="w-5 h-5 text-yellow-500 animate-pulse" />
        </div>
        
        <p className="text-base-content/60 text-sm leading-relaxed">
          Say hello to <span className="font-medium text-primary">{selectedUser?.username}</span> and 
          break the ice. Every great conversation starts with a simple message! 💬
        </p>
        
        {/* Fun emoji animation */}
        <div className="flex justify-center gap-2 text-2xl">
          <span className="animate-bounce" style={{animationDelay: '0ms'}}>👋</span>
          <span className="animate-bounce" style={{animationDelay: '100ms'}}>😊</span>
          <span className="animate-bounce" style={{animationDelay: '200ms'}}>💭</span>
        </div>
      </div>
    </div>
  );

  return (
    <div className="flex-1 flex flex-col overflow-auto">
      <ChatHeader />

      <div className="flex-1 overflow-y-auto p-4 space-y-4">

        {messages.length === 0 ? (
          <EmptyMessageState />
        ) : null}
        {messages.map((message, index) => {
          const isMyMessage = message.sender_id === authUser?.userId;
          const isLastMessage = index === messages.length - 1;
          
          return (
            <div
              key={message._id}
              className={`chat ${isMyMessage ? "chat-end" : "chat-start"}`}
              ref={isLastMessage ? messageEndRef : null}
            >
              <div className="chat-image avatar">
                <div className="size-10 rounded-full border">
                  <img
                    src={avatar}
                    alt="profile pic"
                  />
                </div>
              </div>
              <div className="chat-header mb-1">
                <time className="text-xs opacity-50 ml-1">
                  {formatMessageTime(message.created_at)}
                </time>
              </div>
              <div className="chat-bubble flex flex-col">
                {message.image && (
                  <img
                    src={message.image}
                    alt="Attachment"
                    className="sm:max-w-[200px] rounded-md mb-2"
                  />
                )}
                {message.content && <p>{message.content}</p>}
              </div>
            </div>
          );
        })}
      </div>

      <MessageInput />
    </div>
  );
};
export default ChatContainer;