package user

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/connerj70/seva/internal/app/seva"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Retrieve will retrieve a user specified by the userID
func Retrieve(db *mongo.Database, userID string) (user seva.User, err error) {
	userCollection := db.Collection("user")

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return user, fmt.Errorf("failed to convert to ObjectID %w", err)
	}

	filter := bson.D{{"_id", objectID}}
	err = userCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, err
}

// Create will create a new user
func Create(db *mongo.Database, newUser seva.NewUser) (userID string, err error) {

	userCollection := db.Collection("user")

	// Check if this user already exists
	userAlreadyExistsFilter := bson.D{{"email", newUser.Email}}
	res := userCollection.FindOne(context.Background(), userAlreadyExistsFilter)
	if res.Err() == nil {
		return "", fmt.Errorf("a user with an email of %q already exists", newUser.Email)
	}
	if res.Err() != nil && res.Err() != mongo.ErrNoDocuments {
		return "", fmt.Errorf("failed to query for existing user %w", res.Err())
	}

	err = newUser.Validate()
	if err != nil {
		return "", fmt.Errorf("user failed validation: %w", err)
	}

	// Hash the users password
	h := sha256.New()
	h.Write([]byte(newUser.Password))
	hashedPassword := fmt.Sprintf("%x", h.Sum(nil))
	newUser.Password = hashedPassword

	newUser.CreatedAt = time.Now().UTC()

	// Insert the user into the DB
	insertResult, err := userCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		return "", fmt.Errorf("database error: %w", err)
	}
	id := insertResult.InsertedID.(primitive.ObjectID)

	return id.Hex(), nil
}

func LogIn(db *mongo.Database, email, password, jwtSecret string) (user seva.User, err error) {
	userCollection := db.Collection("user")
	// Get the user
	userFilter := bson.D{{"email", email}}
	result := userCollection.FindOne(context.Background(), userFilter)
	if result.Err() != nil {
		return user, fmt.Errorf("failed to retrieve user %w", err)
	}

	err = result.Decode(&user)
	if err != nil {
		return user, fmt.Errorf("failed to decode user %w", err)
	}

	// make sure password matches
	// Hash the users password
	h := sha256.New()
	h.Write([]byte(password))
	hashedPassword := fmt.Sprintf("%x", h.Sum(nil))

	if hashedPassword != user.Password {
		return user, fmt.Errorf("passwords do not match")
	}

	oneWeekInFuture := time.Now().AddDate(0, 0, 7).UTC().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: oneWeekInFuture,
		Issuer:    "main",
	})
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return user, fmt.Errorf("failed to sign token %w", err)
	}

	user.JWT = tokenString
	user.Password = ""

	return user, nil
}
