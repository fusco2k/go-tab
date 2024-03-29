package items

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Item describe a basic menu item struct
type Item struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty"`
	Price   int8               `json:"price,omitempty" bson:"price,omitempty"`
	Visible bool               `json:"visible,omitempty" bson:"visible,omitempty"`
}

//AllData returns a slice of Items
func AllData(cl *mongo.Collection) []Item {
	//initialize a slice model to get data
	var Items []Item
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
		//initialize a model item to receive data from the cursor
		item := Item{}
		//decode cursor data into item
		err = cursor.Decode(&item)
		if err != nil {
			cancel()
			log.Fatal(err)
		}
		//append the results into the slice of items
		Items = append(Items, item)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	//returns the slice model
	return Items
}

//OneData returns the item from a ObjectID
func OneData(cl *mongo.Collection, id primitive.ObjectID) Item {
	//initialize the model to decoded the mongo data
	item := Item{}
	//creates a context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//cancel de ctx, all jobs done
	defer cancel()
	//gets the item related to id and decode to the pointed item model
	err := cl.FindOne(ctx, bson.M{"_id": id}).Decode(&item)
	if err != nil {
		cancel()
		return item
	}
	//returns the item
	return item
}

//CreateData creates a item and returns the create item
func CreateData(cl *mongo.Collection, i Item) primitive.ObjectID {
	//creates a context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//cancel de ctx, all jobs done
	defer cancel()
	//creates a new item on the collection
	res, err := cl.InsertOne(ctx, i)
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
func ModifyData(cl *mongo.Collection, i []Item) primitive.ObjectID {
	//creates a context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//cancel de ctx, all jobs done
	defer cancel()
	//Replace the data on the collection
	res, err := cl.ReplaceOne(ctx, bson.M{"_id": i[0].ID}, i[1])
	if err != nil {
		cancel()
	}
	//decode response
	obj := res.UpsertedID.(primitive.ObjectID)
	//returns the objectID of the created item
	return obj
	//todo not finished
}
