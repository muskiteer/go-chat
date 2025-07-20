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
	UserSocketMap = make(map[string]*websocket.Conn)
	Mu            sync.Mutex
)


func PrintUserSocketMap() {
	Mu.Lock()
	defer Mu.Unlock()

	log.Println("Current UserSocketMap:")
	for userId := range UserSocketMap {
		log.Println(" -", userId)
	}
}


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
	Mu.Lock()
	UserSocketMap[userId] = conn
	Mu.Unlock()

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
		Mu.Lock()
		delete(UserSocketMap, userId)
		Mu.Unlock()
		broadcastOnlineUsers()
	}()
}



func broadcastOnlineUsers() {
	Mu.Lock()
	defer Mu.Unlock()

	users := make([]string, 0, len(UserSocketMap))
	for uid := range UserSocketMap {
		users = append(users, uid)
	}

	payload, _ := json.Marshal(map[string]interface{}{
		"event": "getOnlineUsers",
		"data":  users,
	})

	for _, conn := range UserSocketMap {
		if err := conn.WriteMessage(websocket.TextMessage, payload); err != nil {
			log.Println("Error sending online users to client:", err)
			conn.Close()
		}
	}
	log.Println("Broadcasted online users:", users)


	// PrintUserSocketMap()
}


