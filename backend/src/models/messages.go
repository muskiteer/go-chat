package models

import (
	"context"
	"encoding/json"
	"errors"
	
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"time"
)

type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	SenderId  primitive.ObjectID `bson:"sender_id" json:"sender_id"`
	ReceiverId primitive.ObjectID `bson:"receiver_id" json:"receiver_id"`
	Content   string              `bson:"content" json:"content"`
	CreatedAt time.Time           `bson:"created_at" json:"created_at"`
}

func GetOtherUsers(w http.ResponseWriter,r *http.Request,collection *mongo.Collection,currentuser primitive.ObjectID)(error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Filter: all users except the current one
	filter := bson.M{"_id": bson.M{"$ne": currentuser}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return errors.New("failed to retrieve users")
	}
	defer cursor.Close(ctx)
	var users []User
	if err := cursor.All(ctx, &users); err != nil {
		return errors.New("failed to decode users")
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		return errors.New("failed to encode users")
	}
	return nil
}

func GetMessagesForUser(w http.ResponseWriter, r *http.Request, collection *mongo.Collection, userId primitive.ObjectID, otherUserId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Filter: messages between the two users
	filter := bson.M{
		"$or": []bson.M{
			{"sender_id": userId, "receiver_id": otherUserId},
			{"sender_id": otherUserId, "receiver_id": userId},
		},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		
		return errors.New("failed to retrieve messages")
	}
	
	defer cursor.Close(ctx)

	var messages []Message
	if err := cursor.All(ctx, &messages); err != nil {
		return errors.New("failed to decode messages")
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		return errors.New("failed to encode messages")
	}
	return nil
}

func SendMessage(w http.ResponseWriter, r *http.Request, collection *mongo.Collection, userId primitive.ObjectID, otherUserId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var message Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		return errors.New("invalid request payload")
	}

	message.SenderId = userId
	message.ReceiverId = otherUserId
	message.CreatedAt = time.Now()

	// Insert the message into the collection
	_, err := collection.InsertOne(ctx, message)
	if err != nil {
		return errors.New("failed to send message")
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(message); err != nil {
		return errors.New("failed to encode response")
	}
	return nil
}
