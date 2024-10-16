package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

//https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/query-document/

type Db struct {
	client   *mongo.Client
	database *mongo.Database
	coll     *mongo.Collection
}

func NewDb(collection string) (*Db, error) {
	// https://github.com/joho/godotenv // use this lib, to load info from .env files
	// MongoDB connection URI
	const uri = "mongodb://root:example@localhost:27017"
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}
	database := client.Database("info")
	if database == nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	//collection would be automatically created if it doesn't exist
	coll := database.Collection(collection)

	return &Db{client: client, coll: coll}, nil

}

func main() {
	db, err := NewDb("users")
	if err != nil {
		panic(err)
	}
	defer db.client.Disconnect(context.Background())
	err = db.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//db.InsertOne()
	//db.InsertMany()
	//db.Get()
	//db.FindAll()
	//db.Update()
	db.Delete()
}

func (db *Db) Ping(ctx context.Context) error {
	err := db.client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("Failed to ping MongoDB: %w ", err)
	}
	return nil
}

type Person struct {
	FirstName string   `bson:"first_name"`
	Email     string   `bson:"email"`
	Age       int      `bson:"age"`
	Marks     int      `bson:"marks"`
	Hobbies   []string `bson:"hobbies"`
}

func (db *Db) InsertOne() {
	u := Person{
		FirstName: "John",
		Email:     "john@email.com",
		Age:       30,
		Hobbies:   []string{"Sports", "Cooking"},
		Marks:     50,
	}

	// Add code to work with the collection...
	ctx := context.Background()
	res, err := db.coll.InsertOne(ctx, u)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Inserted a single document: ", res.InsertedID)

}

// InsertMany inserts multiple documents
func (db *Db) InsertMany() {
	persons := []Person{
		{
			FirstName: "Jane",
			Email:     "jane@email.com",
			Age:       28,
			Hobbies:   []string{"Reading", "Traveling"},
			Marks:     70,
		},
		{
			FirstName: "Alice",
			Email:     "alice@email.com",
			Age:       35,
			Hobbies:   []string{"Swimming", "Photography"},
			Marks:     60,
		},
		{
			FirstName: "John",
			Email:     "john2@email.com",
			Age:       25,
			Hobbies:   []string{"Music", "Photography"},
			Marks:     90,
		},
	}

	ctx := context.Background()
	res, err := db.coll.InsertMany(ctx, persons)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Inserted multiple documents: ", res.InsertedIDs)
}

// Get retrieves a single document based on a filter
func (db *Db) Get() {
	var person Person
	ctx := context.Background()
	//
	filter := bson.D{{"first_name", bson.D{{"$eq", "John"}}}}
	//filter := bson.D{
	//   {"$and",
	//      bson.A{
	//         bson.D{{"marks", bson.D{{"$gt", 7}}}},
	//         bson.D{{"age", bson.D{{"$lte", 30}}}},
	//      },
	//   },
	//}
	//filter := bson.D{{"first_name", "John"}}

	err := db.coll.FindOne(ctx, filter).Decode(&person)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Found a single document: %+v\n", person)
}

// FindAll retrieves all documents
func (db *Db) FindAll() {
	var results []Person
	ctx := context.Background()

	// get everything
	// or specify a specific condition in bson.D{}
	cur, err := db.coll.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var person Person
		err := cur.Decode(&person)
		if err != nil {
			log.Println(err)
			return
		}
		results = append(results, person)
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
		return
	}
	for _, v := range results {
		fmt.Printf("%+v\n\n", v)
	}
	//fmt.Println("Found multiple documents: ", results)
}

// Update modifies a single document based on a filter
func (db *Db) Update() {
	filter := bson.D{{"email", "john@email.com"}}
	update := bson.D{
		{"$set", bson.D{
			{"age", 32},
		}},
	}

	ctx := context.Background()
	res, err := db.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", res.MatchedCount, res.ModifiedCount)
}

// Delete removes a single document based on a filter
func (db *Db) Delete() {
	filter := bson.D{{"email", "john2@email.com"}}

	ctx := context.Background()
	res, err := db.coll.DeleteOne(ctx, filter)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Deleted %v document(s)\n", res.DeletedCount)
}
