package model


import(
	"github.com/myrachanto/ecommerce/httperrors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Street struct {
	ID 	primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `bson:"name"`
	Title string `bson:"title"`
	Description string `bson:"description"`
	Population float64 `bson:"population"`
	Base
}
func (street Street) Validate() *httperrors.HttpError{
	if street.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if street.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if street.Description == "" {
		return httperrors.NewNotFoundError("Invalid Description")
	}
	return nil
}