package sessions

import (
	"github.com/fusco2k/go-tab/tabs"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	SecretToken uuid.UUID          `json:"secrettoken,omitempty" bson:"secrettoken,omitempty"`
	Tab         tabs.Tab           `json:"tab,omitempty" bson:"tab,omitempty"`
	Time        primitive.DateTime `json:"time,omitempty" bson:"time,omitempty"`
}
