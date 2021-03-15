package model

import(
	"github.com/myrachanto/ecommerce/httperrors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"name,omitempty"`
	Title         string             `json:"title,omitempty"`
	Description   string             `json:"description,omitempty"`
	Code          string             `json:"code,omitempty"`
	Majorcategory string             `json:"majorcat,omitempty"`
	Category      string             `json:"category,omitempty"`
	Subcategory   string             `json:"subcategory,omitempty"`
	Oldprice      float64            `json:"oldprice,omitempty"`
	Newprice      float64            `json:"newprice,omitempty"`
	Buyprice      float64            `json:"buyprice,omitempty"`
	Picture       string             `json:"picture,omitempty"`
	Quantity      float64              `json:"quantity,omitempty"`
	// Tag           []Tag    `json:"tag,omitempty"`
	Rates     []Rating `json:"rates,omitempty"`
	Featured  bool     `json:"featured,omitempty"`
	Promotion bool     `json:"promotion,omitempty"`
	Hotdeals  bool     `json:"hotdeals,omitempty"`
	Base      `json:"base,omitempty"`
}

func (product Product) Validate() *httperrors.HttpError {
	if product.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if product.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if product.Description == "" {
		return httperrors.NewNotFoundError("Invalid Description")
	}
	return nil
}
