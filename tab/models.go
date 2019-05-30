package tabs

import (
	"time"

	"github.com/fusco2k/go-request/orders"

	"github.com/fusco2k/go-request/tables"

	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tab struct {
	ID         primitive.ObjectID
	name       uuid.UUID
	timeOpened time.Time
	table      tables.Table
	orders     []orders.Order
}
