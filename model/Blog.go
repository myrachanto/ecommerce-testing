package model

import(
	"github.com/myrachanto/ecommerce/httperrors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
//Blog ...
type Blog struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string             `bson:"title" json:"title,omitempty"`
	Meta    string             `bson:"meta" json:"meta,omitempty"`
	Header1 string             `bson:"header1" json:"header_1,omitempty"`
	Header2 string             `bson:"header2" json:"header_2,omitempty"`
	Intro   string             `bson:"into" json:"intro,omitempty"`
	Body    string             `json:"body,omitempty"`
	Summary string             `json:"summary,omitempty"`
	Picture string             `json:"picture,omitempty"`
	Code    string             `json:"code,omitempty"`
	Base    `json:"base,omitempty"`
}
//Validate ...
func (blog Blog) Validate() *httperrors.HttpError {
	
	if blog.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if blog.Meta == "" {
		return httperrors.NewNotFoundError("Invalid meta")
	}
	if blog.Intro == "" {
		return httperrors.NewNotFoundError("Invalid INtro")
	}
	return nil
}
