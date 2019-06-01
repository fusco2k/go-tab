package sessions

import (
	"github.com/fusco2k/go-tab/tabs"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID          primitive.ObjectID
	SecretToken uuid.UUID
	Tab         tabs.Tab
	Time        primitive.DateTime
}
