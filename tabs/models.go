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
	TimeOpened time.Time
	Table      int8
	Orders     []orders.Order
}
