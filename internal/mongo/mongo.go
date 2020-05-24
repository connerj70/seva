package mongo

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// ErrEmptyConnString will return when Open is called with "" connectionString
	ErrEmptyConnString = errors.New("connection string is empty")
)

// Open will open a connection to a mongo database
func Open(connectionString string) (*mongo.Client, error) {
	if connectionString == "" {
		return nil, ErrEmptyConnString
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		return client, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return client, err
	}

	return client, nil
}
