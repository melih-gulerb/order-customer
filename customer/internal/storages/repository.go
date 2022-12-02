package storages

import (
	"context"
	c "customer/internal"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Repository struct {
	mc *mongo.Collection
}

func NewRepository(mc *mongo.Collection) Repository {
	return Repository{mc: mc}
}

func (r Repository) InsertOne(customer c.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.mc.InsertOne(ctx, customer)
	if err != nil {
		return err
	}

	return nil
}
