package models

import (
	"context"
	"time"

	"github.com/muskiteer/chat-app/internals"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username        string             `bson:"username" json:"username"`
	Email           string             `bson:"email" json:"email"`
	Hashed_Password string             `bson:"hashed_password" json:"-"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
}

func InsertUser(ctx context.Context, collection *mongo.Collection, username, email, hashedPassword string) (*User, error) {
	if err := internals.ValidateAll(username, email, hashedPassword, ctx, collection); err != nil {
		return nil, err
	}

	hashedPassword, err := internals.HashPassword(hashedPassword)
	if err != nil {
		return nil, err
	}

	newUser := &User{
		ID:              primitive.NewObjectID(),
		Username:        username,
		Email:           email,
		Hashed_Password: hashedPassword,
		CreatedAt:       time.Now(),
	}
	
	_, err = collection.InsertOne(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func GetbyEmailorUsername(ctx context.Context, collection *mongo.Collection, emailOrUsername string) (*User, error) {
	var user User
	filter := map[string]interface{}{
		"$or": []map[string]interface{}{
			{"email": emailOrUsername},
			{"username": emailOrUsername},
		},
	}

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
