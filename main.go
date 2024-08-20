package main

import (
	"context"
	"github.com/amiranbari/challenge/config"
	httpserver "github.com/amiranbari/challenge/delivery/http_server"
	"github.com/amiranbari/challenge/repository/mongodb"
	"github.com/amiranbari/challenge/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func main() {
	cfg := Config()
	db := mongodb.New(cfg.MongoDB)

	// Initialize the database
	if err := initializeDatabase(context.Background(), db.Conn()); err != nil {
		log.Fatal(err)
	}

	svc := Service(cfg, db)
	httpServer := HTTPServer(cfg, svc)
	httpServer.Serve()
}

func Config() config.Config {
	return config.C()
}

func Service(cfg config.Config, db *mongodb.DB) *service.Service {
	return service.New(cfg, db)
}

func HTTPServer(cfg config.Config, svc *service.Service) *httpserver.Server {
	return httpserver.New(cfg, svc)
}

func initializeDatabase(ctx context.Context, db *mongo.Database) error {
	// Define the collections and their initial data
	collections := map[string][]interface{}{
		"users": {
			bson.M{"name": "Alice", "age": 30, "email": "alice@example.com"},
			bson.M{"name": "Bob", "age": 25, "email": "bob@example.com"},
		},
	}

	for collName, docs := range collections {
		collection := db.Collection(collName)

		// Check if collection exists and has data (optional)
		count, err := collection.CountDocuments(ctx, bson.M{})
		if err != nil {
			return err
		}
		if count > 0 {
			continue // Skip if collection already has data
		}

		// Insert initial data
		_, err = collection.InsertMany(ctx, docs)
		if err != nil {
			return err
		}
	}

	return nil
}
