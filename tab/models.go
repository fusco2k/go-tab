package tabs

import (
	"github.com/fusco2k/go-request/tables"
	"time"

	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tab struct {
	ID         primitive.ObjectID
	name       uuid.UUID
	timeOpened time.Time
	table      tables.Table
	orders []Order
}
