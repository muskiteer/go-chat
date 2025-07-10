import { useChatStore } from "../store/useChatStore";
import { useEffect, useRef } from "react";
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
    isMessagesLoading,
    selectedUser,
    subscribeToMessages,
    unsubscribeFromMessages,
  } = useChatStore();
  const { authUser } = useAuthStore();
  const messageEndRef = useRef(null);

  // Add this debugging
  console.log('authUser:', authUser);
  console.log('authUser.id:', authUser.userId);

  useEffect(() => {
    if (selectedUser?.id) { // Add null check
      getMessages(selectedUser.id); // Changed to id
    }
  }, [selectedUser?.id, getMessages]); // Simplified dependencies

  console.log('selectedUser:', selectedUser);
  console.log('selectedUser.id:', selectedUser.id);

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

  return (
    <div className="flex-1 flex flex-col overflow-auto">
      <ChatHeader />

      <div className="flex-1 overflow-y-auto p-4 space-y-4">
        {messages.map((message) => {
          // Add debugging for each message
          console.log('message.sender_id:', message.sender_id);
          console.log('authUser.id:', authUser?.userId);
          console.log('Are they equal?:', message.sender_id === authUser?.userId);

          const isMyMessage = message.sender_id === authUser?.userId;
          
          return (
            <div
              key={message._id}
              className={`chat ${isMyMessage ? "chat-end" : "chat-start"}`}
              ref={messageEndRef}
            >
              <div className=" chat-image avatar">
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