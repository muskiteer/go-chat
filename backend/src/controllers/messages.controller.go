package controllers

import (
	"net/http"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/mongo"
	"github.com/muskiteer/chat-app/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/muskiteer/chat-app/utils"
)

func GetUsersForSidebar(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("id")
		if userID == nil {
			utils.JSONError(w, http.StatusUnauthorized, "User not authenticated")
			return
		}

		err := models.GetOtherUsers(w, r, collection, userID.(primitive.ObjectID))
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