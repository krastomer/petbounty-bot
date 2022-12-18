package bounty

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BountyStatus string

const (
	Missing BountyStatus = "MISSING"
	Founded BountyStatus = "FOUNDED"
)

type Bounty struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserID    string             `bson:"user_id"`
	Name      string             `bson:"name"`
	Reward    float64            `bson:"reward"`
	Detail    string             `bson:"detail"`
	Address   string             `bson:"address"`
	Telephone string             `bson:"telephone"`
	Status    BountyStatus       `bson:"status"`
	CreatedAt time.Time          `bson:"createad_at"`
}
