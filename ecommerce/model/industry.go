package model


import(
	"github.com/myrachanto/ecommerce/httperrors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Industry struct {
	ID 	primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `bson:"name"`
	Title string `bson:"title"`
	Description string `bson:"description"`
	Picture string `bson:"picture"`
	Base
}
func (industry Industry) Validate() *httperrors.HttpError{
	if industry.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if industry.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if industry.Description == "" {
		return httperrors.NewNotFoundError("Invalid Description")
	}
	return nil
}