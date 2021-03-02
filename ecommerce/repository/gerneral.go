package repository

import (
		"github.com/myrachanto/ecommerce/httperrors"
		"github.com/myrachanto/ecommerce/model" 
)
//Generalrepo repo
var (
	Generalrepo generalrepo = generalrepo{}
)
///curtesy to gorm
type generalrepo struct{}
//////////////
////////////TODO user id///////////
/////////////////////////////////////////
func (generalRepo generalrepo) View(search string)(*model.General, *httperrors.HttpError) {
	
	products,err1 := Productrepository.GetAll(search)
	if err1 != nil {
		return nil, err1
	}
	productcount, errs := Productrepository.Count()
	if errs != nil {
		return nil, errs
	}
	users,err2 := Userrepository.Count() 
	if err2 != nil {
		return nil, err2
	}
	blogs,err3 := Blogrepository.Count()
	if err3 != nil {
		return nil, err3
	}
	var productqty float64 = 0
	for _,s := range products {
		productqty += s.Quantity
	}
	general := model.General{}
	general.Inventory.Total = productqty
	general.Inventory.Name = "Inventory"
	general.Inventory.Description = "Total number of all the stocks in the store"
	general.Products.Total = productcount
	general.Products.Name = "Products"
	general.Products.Description = "Total number of all the products in the store"
	general.Users.Total = users
	general.Users.Name = "Users"
	general.Users.Description = "clients registered"
	general.Blogs.Total = blogs
	general.Blogs.Name = "Blogs"
	general.Blogs.Description = "The  total number of blogs "

	return &general, nil
}
// func (generalRepo generalrepo) Email() (*model.Email, *httperors.HttpError) {
	
// 	email := model.Email{}
// 	email.Email = "Business@gmail.com"
// 	email.To = "example@gmail.com"
// 	email.Subject = "RE:"
// 	email.Message = "this is the email message body"
// 	customers,err4 := Customerrepo.All()
// 	if err4 != nil {
// 		return nil, err4
// 	}
// 	email.Customers = customers
	
// 	return &email, nil
// }

