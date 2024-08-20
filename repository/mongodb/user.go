package mongodb

import (
	"context"
	"github.com/amiranbari/challenge/entity"
	params "github.com/amiranbari/challenge/param"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (d *DB) GetAllUsers(ctx context.Context, filter params.FilterRequest) ([]entity.User, error) {
	// Get the collection from the database
	collection := d.db.Collection("users")

	users := make([]entity.User, 0)

	bsonFilter := bson.M{}

	for key, value := range filter {
		bsonFilter[key] = value
	}

	findOptions := options.Find()

	cursor, err := collection.Find(ctx, bsonFilter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx) // Ensure the cursor is closed after we're done

	// Iterate over the cursor and decode each document into a User entity
	for cursor.Next(ctx) {
		var user entity.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Check if the cursor encountered any errors during iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
