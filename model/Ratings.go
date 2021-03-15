package model

import(
	"github.com/myrachanto/ecommerce/httperrors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rating struct {
	ID 	primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string  `bson:"name"`
	Rates       float64  `bson:"rates"`
	Comments string  `bson:"comments"`
	Base
}

func (rating Rating) Validate() *httperrors.HttpError {
	if rating.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if rating.Rates > 0 {
		return httperrors.NewNotFoundError("Invalid rate")
	}
	if rating.Comments == "" {
		return httperrors.NewNotFoundError("Invalid Comments")
	}
	return nil
}
