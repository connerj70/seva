package sandwich

import "go.mongodb.org/mongo-driver/mongo"

type Service struct{ DB *mongo.Client }
