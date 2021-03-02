package model


import (
	"github.com/myrachanto/ecommerce/httperrors"
	"go.mongodb.org/mongo-driver/bson/primitive"

)
//Majorcategory ...
type Majorcategory struct {
	ID 	primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string     `json:"name,omitempty" bson:"name,omitempty"`
	Title       string     `json:"title,omitempty" bson:"title,omitempty"`
	Description string     `json:"description,omitempty" bson:"description,omitempty"`
	Code        string     `json:"code,omitempty" bson:"code,omitempty"`
	Category    []*Category `json:"category,omitempty" bson:"category,omitempty"`
	Base        `json:"base,omitempty" bson:"base,omitempty"`
}
//Validate ...
func (majorcategory Majorcategory) Validate() *httperrors.HttpError{
	if majorcategory.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if majorcategory.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if majorcategory.Description == "" {
		return httperrors.NewNotFoundError("Invalid Description")
	}
	return nil
}