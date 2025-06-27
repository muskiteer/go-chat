package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/muskiteer/chat-app/src/controllers"
	"github.com/muskiteer/chat-app/src/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthRoutes(r *mux.Router,userCollection *mongo.Collection) {
	

	// Define your authentication routes here
	r.HandleFunc("/login", controllers.LoginHandler(userCollection)).Methods("POST")
	r.HandleFunc("/logout", controllers.LogoutHandler).Methods("POST")
	r.HandleFunc("/signup", controllers.SignupHandler(userCollection)).Methods("POST")


	
	r.Handle("/check", middleware.Authenticate(userCollection)(http.HandlerFunc(controllers.CheckAuthentication))).Methods("GET")

}