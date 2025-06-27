package utils

import (
	"net/http"
	"os"
	"time"
	"errors"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID string `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateToken(userId string, w http.ResponseWriter) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    signedToken,
		Expires:  expirationTime,
		MaxAge:   int(time.Until(expirationTime).Seconds()),
		HttpOnly: true,                      // Prevent XSS
		SameSite: http.SameSiteStrictMode,   // Prevent CSRF
		Secure:   os.Getenv("NODE_ENV") != "development", // Secure in production
		Path:     "/",
	}
	http.SetCookie(w, cookie)

	return signedToken, nil
}

func VerifyJWT(tokenStr string, collection *mongo.Collection) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid || claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("invalid or expired token")
	}

	// Convert string to ObjectID
	objectID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		return nil, errors.New("invalid user ID in token")
	}

	filter := bson.M{"_id": objectID}
	var result bson.M
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user does not exist")
		}
		return nil, errors.New("database error")
	}

	return claims, nil
}


