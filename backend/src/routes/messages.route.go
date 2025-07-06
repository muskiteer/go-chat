package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/muskiteer/chat-app/src/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/muskiteer/chat-app/src/controllers"

)

func MessagesRoutes(r *mux.Router,messageCollection *mongo.Collection,userCollection *mongo.Collection) {
	r.Handle("/users", middleware.Authenticate(userCollection)(http.HandlerFunc(controllers.GetUsersForSidebar(userCollection)))).Methods("GET")
	r.Handle("/{id}", middleware.Authenticate(userCollection)(http.HandlerFunc(controllers.GetMessagesForUser(messageCollection)))).Methods("GET")
	r.Handle("/send/{id}", middleware.Authenticate(userCollection)(http.HandlerFunc(controllers.SendMessage(messageCollection)))).Methods("POST")
}
