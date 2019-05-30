package items

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	id    primitive.ObjectID
	name  string
	price float64
	stock int8
}
