package controllers

import (
	"context"
	"fmt"
	"log"
	"todolist-mongodb/driver"
	"todolist-mongodb/models"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

type Controller struct{}

var client = driver.Driver()
var collection = client.Database("todolist").Collection("errands")

func (c *Controller) InsertOne(t models.Todo) {

	insertResult, err := collection.InsertOne(context.TODO(), t)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func (c *Controller) UpdateOne(t models.Todo) {
	filter := bson.D{{"id", t.ID}}

	update := bson.D{
		{"$set", bson.D{
			{"description", t.Description},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

func (c *Controller) FindOne(t models.Todo) {
	// create a value into which the result can be decoded
	var result models.Todo
	filter := bson.D{{"title", t.Title}}

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)
}

func (c *Controller) Find(setLimit int64) {
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(setLimit)

	// Here's an array in which you can store the decoded documents
	var results []*models.Todo

	// Passing nil as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem models.Todo
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	for _, r := range results {
		fmt.Printf("Found multiple documents (array of pointers): %+v\n", r)
	}

}

// "Finally, you can delete documents using collection.
// DeleteOne() or collection.DeleteMany().
// Here you pass nil as the filter argument, which will match all documents in the collection.
// You could also use collection.Drop() to delete an entire collection.
func (c *Controller) DeleteAll() {
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}
func (c *Controller) DropDB() {
	collection.Drop(context.TODO())
	fmt.Println(collection.Drop(context.TODO()))
}

func (c *Controller) DeleteOne(t models.Todo) {
	filter := bson.D{{"title", t.Title}}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	fmt.Println(deleteResult)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}
