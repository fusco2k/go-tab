package tables

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID     primitive.ObjectID
	number int
	total  float64
}
