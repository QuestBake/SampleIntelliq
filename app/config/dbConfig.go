package config

import (
	"context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
)

const (
	//URL to connect to mongodb
	URL = "mongodb://localhost:27017"
	//DbName name of mongodb
	DbName = "intelliQ"
)

var client *mongo.Client
var ctx context.Context

//Connect db conn
func Connect() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	client, err = mongo.NewClient(URL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("Successfully connected to DB at ", URL)
	return nil
}

//GetCollection copy of original session
func GetCollection(dbName string, collName string) *mongo.Collection {
	if client == nil {
		return nil
	}
	db := client.Database(dbName)
	if db == nil {
		return nil
	}
	coll := db.Collection(collName)
	if coll == nil {
		return nil
	}
	return coll
}

func CreateIndices() {
	db := client.Database(DbName)
	if db == nil {
		panic("No DB session")
	}

	var searchFields []string
	searchFields = append(searchFields, "city")
	searchFields = append(searchFields, "state")

	addSearchIndex(db, "addresses", searchFields)
	addUniqueIndex(db, "addresses", []string{"city"})
}

func addSearchIndex(db *mongo.Database, collName string, searchFields []string) {
	coll := db.Collection(collName)
	if coll == nil {
		panic("No such Collection in DB" + collName)
	}
	var indexes []mongo.IndexModel

	for _, val := range searchFields {
		indexes = append(indexes, mongo.IndexModel{
			Keys: bsonx.Doc{{
				Key:   val,
				Value: bsonx.Int32(1),
			}},
			Options: options.Index().SetName("textIndex"),
		},
		)
	}
	iv := coll.Indexes()
	_, err := iv.CreateMany(ctx, indexes)
	if err != nil {
		panic("Hi Could not create search index for " + collName + err.Error())
	}
}

func addUniqueIndex(db *mongo.Database, collName string, fields []string) {
	coll := db.Collection(collName)
	if coll == nil {
		panic("No such Collection in DB" + collName)
	}
	var indexes []mongo.IndexModel

	for _, key := range fields {
		indexes = append(indexes, mongo.IndexModel{
			Keys: bsonx.Doc{{Key: key,
				Value: bsonx.Int32(1),
			}},
			Options: options.Index().SetUnique(true),
		})
	}
	iv := coll.Indexes()
	_, err := iv.CreateMany(ctx, indexes)
	if err != nil {
		panic("Could not create unique index for " + collName)
	}
}

//GetContext creates context with timeout
func GetContext() context.Context {
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	return ctx
}
