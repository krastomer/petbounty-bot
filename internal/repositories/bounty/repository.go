package bounty

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "bounty"

type Repository interface {
	GetBounty(ctx context.Context) ([]*Bounty, error)
	CreateBounty(ctx context.Context, bounty Bounty) error
	GetBountyByUserID(ctx context.Context, userID string) ([]*Bounty, error)
}

type repository struct {
	collection *mongo.Collection
}

func InitializeRepository(database *mongo.Database) Repository {
	return &repository{
		collection: database.Collection(collectionName),
	}
}

func (r *repository) GetBounty(ctx context.Context) ([]*Bounty, error) {
	filter := bson.M{}

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
	_, err := r.collection.InsertOne(ctx, bounty)
	return err
}

func (r *repository) GetBountyByUserID(ctx context.Context, userID string) ([]*Bounty, error) {
	filter := bson.M{}
	filter["user_id"] = userID

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
