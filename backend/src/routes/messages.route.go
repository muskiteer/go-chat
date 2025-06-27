package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/muskiteer/chat-app/src/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/muskiteer/chat-app/src/controllers"

)

func MessagesRoutes(r *mux.Router,userCollection *mongo.Collection) {
	r.Handle("/users", middleware.Authenticate(userCollection)(http.HandlerFunc(controllers.GetUsersForSidebar(userCollection)))).Methods("GET")
	r.Handle("/:id", middleware.Authenticate(userCollection)(http.HandlerFunc(controllers.GetMessagesForUser(userCollection)))).Methods("GET")
	r.Handle("/send/:id", middleware.Authenticate(userCollection)(http.HandlerFunc(controllers.SendMessage(userCollection)))).Methods("POST")
}
