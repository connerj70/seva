package noauth

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Service will hold our connection to the database along with all the methods for the service layer
type Service struct {
	DB *mongo.Client
}

// Register will handle putting the user into our database
func (s *Service) Register(user *User) error {
	collection := s.DB.Database("testing").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{"email": user.Email, "password": user.Password})
	if err != nil {
		return fmt.Errorf("error inserting into collection %s", err.Error())
	}
	id := res.InsertedID
	fmt.Println(id)
	return nil
}

//GetUserByEmail will return a user if it finds one
func (s *Service) GetUserByEmail(userEmail string) (*User, error) {
	user := &User{}
	collection := s.DB.Database("testing").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, bson.M{"email": userEmail}).Decode(user)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, fmt.Errorf("GetUserByEmail %s", err.Error())
	}
	return user, nil
}
