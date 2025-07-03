package controllers

import (
	"net/http"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/mongo"
	"github.com/muskiteer/chat-app/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/muskiteer/chat-app/utils"
	"github.com/muskiteer/chat-app/src/middleware"
)

func GetUsersForSidebar(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val := r.Context().Value(middleware.UserIDKey)
		userIDStr, ok := val.(string)
		if !ok {
			utils.JSONError(w, http.StatusUnauthorized, "Invalid user ID in context")
			return
		}

		userID, err := primitive.ObjectIDFromHex(userIDStr)
		if err != nil {
			utils.JSONError(w, http.StatusUnauthorized, "Invalid ObjectID format")
			return
		}

		err = models.GetOtherUsers(w, r, collection, userID)
		if err != nil {
			utils.JSONError(w, http.StatusInternalServerError, "Failed to retrieve users")
			return
		}
	}
}


func GetMessagesForUser(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("Id")
		
		
		params := mux.Vars(r)
		otherUserID, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil {
			utils.JSONError(w, http.StatusBadRequest, "Invalid user ID")
			return
		}

		err = models.GetMessagesForUser(w, r, collection, userID.(primitive.ObjectID), otherUserID)
		if err != nil {
			utils.JSONError(w, http.StatusInternalServerError, "Failed to retrieve messages")
			return
		}
	}
}

func SendMessage(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("Id")
		
		params := mux.Vars(r)
		otherUserID, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil {
			utils.JSONError(w, http.StatusBadRequest, "Invalid user ID")
			return
		}

		err = models.SendMessage(w, r, collection, userID.(primitive.ObjectID), otherUserID)
		if err != nil {
			utils.JSONError(w, http.StatusInternalServerError, "Failed to send message")
			return
		}
	}
}