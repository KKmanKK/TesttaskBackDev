package initializers

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func GetDbCollection(col string) *mongo.Collection {
	return db.Collection(col)
}

func InitDb() error {
	clientOptions := options.Client().ApplyURI(os.Getenv("URI_DB"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	db = client.Database("testDb")
	return nil
}

func CloseDb() error {
	return db.Client().Disconnect(context.Background())
}
