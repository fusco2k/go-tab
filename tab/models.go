package tabs

import (
	"time"

	"github.com/fusco2k/go-tab/orders"

	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Tab relational struct
type Tab struct {
	ID         primitive.ObjectID
	name       uuid.UUID
	timeOpened time.Time
	table      int8
	orders     []orders.Order
}
