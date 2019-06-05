package tabs

import (
	"time"

	"github.com/fusco2k/go-tab/orders"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Tab relational struct
type Tab struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Active     bool               `json:"active,omitempty" bson:"active,omitempty"`
	Number     int                `json:"number,omitempty" bson:"number,omitempty"`
	TimeOpened time.Time          `json:"timeopened,omitempty" bson:"timeopened,omitempty"`
	Table      int8               `json:"table,omitempty" bson:"table,omitempty"`
	Orders     []orders.Order     `json:"orders,omitempty" bson:"orders,omitempty"`
}
