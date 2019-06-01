package sessions

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type session struct {
	id          primitive.ObjectID
	secretToken uuid.UUID
	time        primitive.DateTime
}
