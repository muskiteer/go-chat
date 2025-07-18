package lib

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/muskiteer/chat-app/utils"
	"encoding/json"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return r.Header.Get("Origin") == "http://localhost:5173" || r.Header.Get("Origin") == "https://your-production-domain.com"
	},
}

var (
	userSocketmap = make(map[string]*websocket.Conn)
	mu sync.Mutex
)

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		utils.JSONError(w, http.StatusBadRequest, "userId query parameter is required")
		log.Println("userId query parameter is missing")
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "Failed to upgrade connection")
		log.Println("Error upgrading connection:", err)
		return
	}
	log.Println("WebSocket connection established for user:", userId)

	// Save the user socket
	mu.Lock()
	userSocketmap[userId] = conn
	mu.Unlock()

	// Broadcast online users to all clients
	broadcastOnlineUsers()

	// Read loop to detect disconnect
	go func() {
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				break
			}
		}

		// Remove socket and broadcast
		log.Println("user disconnected:", userId)
		mu.Lock()
		delete(userSocketmap, userId)
		mu.Unlock()
		broadcastOnlineUsers()
	}()
}



func broadcastOnlineUsers() {
	mu.Lock()
	defer mu.Unlock()

	users:= make([]string, 0, len(userSocketmap))
	for uid := range userSocketmap {
		users = append(users, uid)
	}

	payload,_:= json.Marshal(map[string]interface{}{
		"event": "getOnlineUsers",
		"data":  users,
	})

	for _, conn := range userSocketmap {
		if err := conn.WriteMessage(websocket.TextMessage, payload); err != nil {
			log.Println("Error sending online users to client:", err)
			conn.Close()
		}
	}
	log.Println("Broadcasted online users:", users)
}


