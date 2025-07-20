package realtime

import (
	"encoding/json"
	"log"

	"time"

	"github.com/gorilla/websocket"
	"github.com/muskiteer/chat-app/src/lib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	SenderId   primitive.ObjectID `bson:"sender_id" json:"sender_id"`
	ReceiverId primitive.ObjectID `bson:"receiver_id" json:"receiver_id"`
	Content    string             `bson:"content" json:"content"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
}

func SendMessageToUser(receiverId string, message Message) {
	log.Println("Sending message to user:", receiverId)
	lib.Mu.Lock()
	conn, exists := lib.UserSocketMap[receiverId]
	lib.Mu.Unlock()

	if !exists {
		log.Println("Receiver not connected:", receiverId)
		lib.PrintUserSocketMap()
		return
	}

	payload, err := json.Marshal(map[string]interface{}{
		"event": "newMessage",
		"data":  message,
	})
	if err != nil {
		log.Println("Marshal error:", err)
		return
	}

	// log.Printf("Payload being sent: %s\n", string(payload))
	if err := conn.WriteMessage(websocket.TextMessage, payload); err != nil {

		log.Println("WebSocket send error:", err)
		conn.Close()
		lib.Mu.Lock()
		delete(lib.UserSocketMap, receiverId)
		lib.Mu.Unlock()
	}

}
