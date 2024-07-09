package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive" // We always need to add primitive to it
)

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}

// ikde je ID cha type aahe primitive.ObjectID he te aahe jyachat ki aplyala id dili jail
