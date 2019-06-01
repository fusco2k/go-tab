package tabs

import (
	"time"

	"github.com/fusco2k/go-tab/orders"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Tab relational struct
type Tab struct {
	ID         primitive.ObjectID
	Number     int8
	timeOpened time.Time
	table      int8
	orders     []orders.Order
}
