package models

import (
	"time"

	"github.com/fusco2k/go-request/items"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	id            primitive.ObjectID
	items         []items.Item
	time          time.Time
	cost          float64
	beingPrepared bool
	isDone        bool
}
