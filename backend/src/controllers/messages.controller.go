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
		// Get the user ID from context (stored as a string by the middleware)
		userIDStr, ok := r.Context().Value(middleware.UserIDKey).(string)
		if !ok {
			utils.JSONError(w, http.StatusUnauthorized, "Invalid user context")
			return
		}

		// Convert to ObjectID
		userID, err := primitive.ObjectIDFromHex(userIDStr)
		if err != nil {
			utils.JSONError(w, http.StatusBadRequest, "Invalid user ID format")
			return
		}

		// Get other user ID from route param
		params := mux.Vars(r)
		otherUserID, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil {
			utils.JSONError(w, http.StatusBadRequest, "Invalid user ID")
			return
		}

		// Call the model function
		err = models.GetMessagesForUser(w, r, collection, userID, otherUserID)
		if err != nil {
			utils.JSONError(w, http.StatusInternalServerError, "Failed to retrieve messages")
			return
		}
	}
}


func SendMessage(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIDStr, ok := r.Context().Value(middleware.UserIDKey).(string)
		if !ok {
			utils.JSONError(w, http.StatusUnauthorized, "Invalid user context")
			return
		}

		// Convert to ObjectID
		userID, err := primitive.ObjectIDFromHex(userIDStr)
		if err != nil {
			utils.JSONError(w, http.StatusBadRequest, "Invalid user ID format")
			return
		}

		
		params := mux.Vars(r)
		otherUserID, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil {
			utils.JSONError(w, http.StatusBadRequest, "Invalid user ID")
			return
		}

		err = models.SendMessage(w, r, collection, userID, otherUserID)
		if err != nil {
			utils.JSONError(w, http.StatusInternalServerError, "Failed to send message")
			return
		}

		
	}
}