package bounty

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "bounty"

type Repository interface {
	GetBounty(ctx context.Context, name string) ([]*Bounty, error)
}

type repository struct {
	collection *mongo.Collection
}

func InitializeRepository(database *mongo.Database) Repository {
	return &repository{
		collection: database.Collection(collectionName),
	}
}

func (r *repository) GetBounty(ctx context.Context, name string) ([]*Bounty, error) {
	filter := bson.M{}
	if name != "" {
		filter["name"] = name
	}

	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	bounties := make([]*Bounty, 0)
	if err = cur.All(ctx, &bounties); err != nil {
		return nil, err
	}

	return bounties, nil
}

func (r *repository) CreateBounty(ctx context.Context, bounty Bounty) error {
	bounty.ID = primitive.NewObjectID()
	bounty.CreatedAt = time.Now()
	bounty.Status = Missing
	_, err := r.collection.InsertOne(ctx, bounty)
	return err
}
