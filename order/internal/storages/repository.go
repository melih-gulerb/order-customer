package storages

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"order/internal"
	"time"
)

type Repository struct {
	mc *mongo.Collection
}

func NewRepository(mc *mongo.Collection) Repository {
	return Repository{mc: mc}
}

func (r Repository) InsertOne(order order.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.mc.InsertOne(ctx, order)
	if err != nil {
		return err
	}

	return nil
}
