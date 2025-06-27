package internals

import(
	"log"
	"errors"
	"regexp"
	"context"
	"go.mongodb.org/mongo-driver/mongo"

)

func ValidateemptyString(str string, fieldName string) error {
	if str == "" {
		err := errors.New(fieldName + " cannot be empty")
		log.Println(err)
		return err
	}
	return nil
}

func ValidateEmail(email string) error {
	if email == "" {
		err := errors.New("email cannot be empty")
		log.Println(err)
		return err
	}
	if len(email) < 5 || len(email) > 50 {
		err := errors.New("email must be between 5 and 50 characters")
		log.Println(err)
		return err
	}
	if !IsValidEmailFormat(email) {
		err := errors.New("invalid email format")
		log.Println(err)
		return err
	}
	return nil
}
func IsValidEmailFormat(email string) bool {
	// Simple RFC 5322 compliant regex for basic validation
	regex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(regex)
	if !re.MatchString(email) {
		// err := errors.New("this is not a valid email address")
		// log.Println(err)
		return false
	}
	return true
}

func AlreadyExists(ctx context.Context, collection *mongo.Collection, username, email string) bool {
	filter := map[string]interface{}{
		"$or": []map[string]interface{}{
			{"username": username},
			{"email": email},
		},
	}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		log.Println("Error checking if user exists:", err)
		return false // Optional: could return error instead for clarity
	}

	return count != 0
}

func ValidateAll(username, email, password string, ctx context.Context, collection *mongo.Collection) error {
	if err := ValidateemptyString(username, "username"); err != nil {
		return err
	}
	if err := ValidateEmail(email); err != nil {
		return err
	}
	if err := ValidateemptyString(password, "password"); err != nil {
		return err
	}
	if AlreadyExists(ctx, collection, username, email) {
		return errors.New("username already exists")
	}
	return nil
}

