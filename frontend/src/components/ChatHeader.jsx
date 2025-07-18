import { X } from "lucide-react";
import { useAuthStore } from "../store/useAuthStore";
import { useChatStore } from "../store/useChatStore";
import avatar from "../../assets/avatar.png";


const ChatHeader = () => {
  const { selectedUser, setSelectedUser } = useChatStore();
  const { onlineUsers } = useAuthStore();
  // if (selectedUser?._id) {
  //   const isOnline = onlineUsers.includes(selectedUser._id);
  //   console.log("🟢 Selected User ID:", selectedUser._id);
  //   console.log("🟢 Online Users:", onlineUsers);
  //   console.log("🟢 User is", isOnline ? "Online ✅" : "Offline ❌");
  // } else {
  //   console.log("⚠️ No selected user");
  // }

  return (
    <div className="p-2.5 border-b border-base-300">
      <div className="flex items-center justify-between">
        <div className="flex items-center gap-3">
          {/* Avatar */}
          <div className="avatar">
            <div className="size-10 rounded-full relative">
              <img src={avatar} alt={selectedUser.username} />
            </div>
          </div>
            
          {/* User info */}
          <div>
            <h3 className="font-medium">{selectedUser.username}</h3>
            <p className="text-sm text-base-content/70">
              {/* console.log({ onlineUsers }); */}
              {onlineUsers.includes(selectedUser._id) ? "Online" : "Offline"}
            </p>
          </div>
        </div>

        {/* Close button */}
        <button onClick={() => setSelectedUser(null)}>
          <X />
        </button>
      </div>
    </div>
  );
};
export default ChatHeader;