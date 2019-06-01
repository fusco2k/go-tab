package items

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Item describe a basic menu item struct
type Item struct {
	id      primitive.ObjectID
	name    string
	price   float64
	visible bool
}
