package bounty

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "bounty"

type Repository interface {
	GetBounties(ctx context.Context) ([]*Bounty, error)
	CreateBounty(ctx context.Context, bounty Bounty) error
	GetBountyByUserID(ctx context.Context, userID string) ([]*Bounty, error)
	UpdateStatusBountyByID(ctx context.Context, id string, status BountyStatus) error
}

type repository struct {
	collection *mongo.Collection
}

func InitializeRepository(database *mongo.Database) Repository {
	return &repository{
		collection: database.Collection(collectionName),
	}
}

func (r *repository) GetBounties(ctx context.Context) ([]*Bounty, error) {
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

func (r *repository) UpdateStatusBountyByID(ctx context.Context, id string, status BountyStatus) error {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.D{{"$set", bson.D{{"status", status}}}}
	_, err = r.collection.UpdateByID(ctx, primitiveID, update)
	fmt.Println("hit")
	fmt.Println(err)
	return err
}
