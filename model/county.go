package model

import(
	"github.com/myrachanto/ecommerce/httperrors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type County struct {
	ID 	primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string  `bson:"name"`
	Title       string  `bson:"title"`
	Description string  `bson:"description"`
	Population  float64 `bson:"float"`
	Base
}

func (county County) Validate() *httperrors.HttpError {
	if county.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if county.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if county.Description == "" {
		return httperrors.NewNotFoundError("Invalid Description")
	}
	return nil
}
