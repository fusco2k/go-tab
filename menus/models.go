package menus

import (
	"context"
	"log"
	"time"

	"github.com/fusco2k/go-tab/items"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Menu represents a menu struct
type Menu struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Category string             `json:"category,omitempty" bson:"category,omitempty"`
	Items    []items.Item       `json:"items,omitempty" bson:"items,omitempty"`
}

//AllData returns a slice of Users
func AllData(cl *mongo.Collection) []Menu {
	//initialize a slice model to get data
	var Menus []Menu
	//creates a context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//cancel de ctx, all jobs done
	defer cancel()
	//gets the cursos with data
	cursor, err := cl.Find(ctx, bson.D{})
	if err != nil {
		cancel()
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	// loop throght the cursor decoding the data and append it to the slice model
	for cursor.Next(ctx) {
		//initialize a model user to receive data from the cursor
		menu := Menu{}
		//decode cursor data into user
		err = cursor.Decode(&menu)
		if err != nil {
			cancel()
			log.Fatal(err)
		}
		//append the results into the slice of menus
		Menus = append(Menus, menu)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	//returns the slice model
	return Menus
}

//OneData returns the item from a ObjectID
func OneData(cl *mongo.Collection, id primitive.ObjectID) Menu {
	//initialize the model to decoded the mongo data
	menu := Menu{}
	//creates a context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//cancel de ctx, all jobs done
	defer cancel()
	//gets the item related to id and decode to the pointed item model
	err := cl.FindOne(ctx, bson.M{"_id": id}).Decode(&menu)
	if err != nil {
		cancel()
		return menu
	}
	//returns the item
	return menu
}

//CreateData creates a item and returns the create item
func CreateData(cl *mongo.Collection, m Menu) primitive.ObjectID {
	//creates a context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//cancel de ctx, all jobs done
	defer cancel()
	//creates a new item on the collection
	res, err := cl.InsertOne(ctx, m)
	if err != nil {
		cancel()
	}
	//decode response
	obj := res.InsertedID.(primitive.ObjectID)
	//returns the objectID of the created item
	return obj
}

//DeleteData deletes the item of given id
func DeleteData(cl *mongo.Collection, id primitive.ObjectID) int64 {
	//creates a context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//cancel de ctx, all jobs done
	defer cancel()
	//delete the item from the collection
	res, err := cl.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		cancel()
	}
	//return the number of deleted documents
	return res.DeletedCount
}

//ModifyData replace the item given on pos 0 from slice by the item on pos 1
func ModifyData(cl *mongo.Collection, m []Menu) primitive.ObjectID {
	//creates a context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//cancel de ctx, all jobs done
	defer cancel()
	//Replace the data on the collection
	res, err := cl.ReplaceOne(ctx, bson.M{"_id": m[0].ID}, m[1])
	if err != nil {
		cancel()
	}
	//decode response
	obj := res.UpsertedID.(primitive.ObjectID)
	//returns the objectID of the created item
	return obj
	//todo not finished
}
