package controllers

import (
	"encoding/json"
	"net/http"
	"time"
	 "os"
	
	"github.com/muskiteer/chat-app/internals"
	"github.com/muskiteer/chat-app/utils"
	"github.com/muskiteer/chat-app/src/models"
	"github.com/muskiteer/chat-app/src/middleware"
	
	"go.mongodb.org/mongo-driver/mongo"
)



type SignupRequest struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

type LoginRequest struct {
	Email          string `json:"email_or_username"`
	HashedPassword string `json:"hashed_password"`
}



func SignupHandler(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SignupRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.JSONError(w, http.StatusBadRequest, "Invalid JSON")
			return
		}

		user, err := models.InsertUser(r.Context(), collection, req.Username, req.Email, req.HashedPassword)
		if err != nil {
			utils.JSONError(w, http.StatusBadRequest, err.Error())
			return
		}
		_, err = utils.GenerateToken(user.ID.Hex(), w) 
		if err != nil {
			utils.JSONError(w, http.StatusInternalServerError, "Failed to generate token")
			return
		}

		utils.JSONResponse(w, http.StatusCreated, map[string]interface{}{
			"message": "Signup successful",
			"user": map[string]interface{}{
				
				"username": user.Username,
				"email":    user.Email,
				
			},
		})
	}
}



func LoginHandler(collection *mongo.Collection) http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		if err:= json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.JSONError(w, http.StatusBadRequest, "Invalid JSON")
			return
		}

		user, err := models.GetbyEmailorUsername(r.Context(), collection, req.Email)
		if err != nil {
			utils.JSONError(w, http.StatusUnauthorized, "Invalid credentials")
			return
		}
		if !internals.CheckPasswordHash(req.HashedPassword, user.Hashed_Password) {
			utils.JSONError(w, http.StatusUnauthorized, "Invalid credentials")
			return
		}

		_, err = utils.GenerateToken(user.ID.Hex(), w) 
		if err != nil {
			utils.JSONError(w, http.StatusInternalServerError, "Failed to generate token")
			return
		}

		utils.JSONResponse(w, http.StatusCreated, map[string]interface{}{
			"message": "Signup successful",
			"user": map[string]interface{}{
			
			"username": user.Username,
			"email":    user.Email,
		
			},
		})

	}
}

	

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Path:    "/",
		Value:    "",
		Expires:  time.Unix(0, 0), // Set expiration to the past
		MaxAge:   -1,
		HttpOnly: true, // Prevent XSS
		SameSite: http.SameSiteStrictMode, // Prevent CSRF
		Secure:   os.Getenv("NODE_ENV") != "development",
	})

	utils.JSONResponse(w, http.StatusOK, map[string]string{
		"message": "Logged out successfully",
	})


}

func CheckAuthentication(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey)
	if userID == nil {
		utils.JSONError(w, http.StatusUnauthorized, "Unauthorized: No user found in context")
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"message": "User is authenticated",
		"userId":  userID,
	})
}



