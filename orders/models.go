package orders

import (
	"time"

	"github.com/fusco2k/go-tab/items"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Order describe a basic order struct
type Order struct {
	id            primitive.ObjectID
	items         []items.Item
	time          time.Time
	cost          float64
	beingPrepared bool
	done          bool
}
