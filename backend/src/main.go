package main

import (
	"context"
	"net/http"
	"os"
	"log"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/muskiteer/chat-app/src/lib"
	"github.com/muskiteer/chat-app/src/middleware"
	"github.com/muskiteer/chat-app/src/routes"
)


func main(){

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file in main")
	}
	client,err := lib.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer client.Disconnect(context.TODO())

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}


	userCollection := client.Database("chat_db").Collection("users")
	messageCollection := client.Database("chat_db").Collection("messages")
	
	r:=mux.NewRouter()

	r.HandleFunc("/ws", lib.HandleWebSocket)
	routes.AuthRoutes(r,userCollection)
	routes.MessagesRoutes(r,messageCollection, userCollection)

	handler := middleware.CORSMiddleware(r)

	log.Println("Starting server on PORT:",PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
	
}