package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Config struct {
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Port     int    `koanf:"port"`
	Host     string `koanf:"host"`
	DBName   string `koanf:"db_name"`
}

type DB struct {
	config Config
	db     *mongo.Database
}

func (m *DB) Conn() *mongo.Database {
	return m.db
}

func New(config Config) *DB {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", config.Username, config.Password, config.Host, config.Port)
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to DB!")

	return &DB{config: config, db: client.Database(config.DBName)}
}
