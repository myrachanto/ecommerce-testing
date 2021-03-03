package model

import(
)
//General ...
type General struct {
	Stocks    Module `json:"stocks,omitempty"`
	Inventory Module `json:"inventory,omitempty"`
	Orders    Module `json:"orders,omitempty"`
	Products  Module `json:"products,omitempty"`
	Users     Module `json:"users,omitempty"`
	Blogs     Module `json:"blogs,omitempty"`
}
//Module ...
type Module struct {
	Name        string  `json:"name,omitempty"`
	Total       float64 `json:"total,omitempty"`
	Description string  `json:"description,omitempty"`
	Icon        string  `json:"icon,omitempty"`
}
