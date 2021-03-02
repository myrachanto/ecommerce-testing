package model

import(
	"github.com/myrachanto/ecommerce/httperrors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
//Seo ...
type Seo struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string             `bson:"title" json:"title,omitempty"`
	Meta    string             `bson:"meta" json:"meta,omitempty"`
	Header1 string             `bson:"header1" json:"header_1,omitempty"`
	Header2 string             `bson:"header2" json:"header_2,omitempty"`
	Picture string             `json:"picture,omitempty"`
	Code    string             `json:"code,omitempty"`
	Base    `json:"base,omitempty"`
}
//Validate ...
func (seo Seo) Validate() *httperrors.HttpError {
	if seo.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if seo.Meta == "" {
		return httperrors.NewNotFoundError("Invalid meta")
	}
	if seo.Header1 == "" {
		return httperrors.NewNotFoundError("Invalid INtro")
	}
	return nil
}
