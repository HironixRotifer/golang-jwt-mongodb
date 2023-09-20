package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID          primitive.ObjectID `bson:"id"`
	ItemID      int64              `json:"item_id"`
	Status      string             `json:"status"`
	Description string             `json:"description"`
}
