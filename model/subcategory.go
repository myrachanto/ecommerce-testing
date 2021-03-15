package model

import(
	"github.com/myrachanto/ecommerce/httperrors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subcategory struct {
	ID 	primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `bson:"name"`
	Title string `bson:"title"`
	Description string `bson:"description"`
	Base
}
func (subcategory Subcategory) Validate() *httperrors.HttpError{
	if subcategory.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if subcategory.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if subcategory.Description == "" {
		return httperrors.NewNotFoundError("Invalid Description")
	}
	return nil
}