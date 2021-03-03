package service

import (
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model" 
	r "github.com/myrachanto/ecommerce/repository"
)

var (
	ProductService  = productService{}
)
type productService struct {
}

func (service productService) Create(product *model.Product) (*httperrors.HttpError) {
	err1 := r.Productrepository.Create(product)
	 return err1

}

func (service productService) GetOne(code string) (*model.Product, *httperrors.HttpError) {
	product, err1 := r.Productrepository.GetOne(code)
	return product, err1
}

func (service productService) GetAll(search string) ([]*model.Product, *httperrors.HttpError) {
	products, err := r.Productrepository.GetAll(search) 
	return products, err
}

func (service productService) GetThree() ([]*model.Product, *httperrors.HttpError) {
	products, err := r.Productrepository.GetThree() 
	return products, err
}

func (service productService) Update(code string, product *model.Product) (*httperrors.HttpError) {
	err1 := r.Productrepository.Update(code, product)
	return err1
}
func (service productService) AUpdate(code string, b, old, new,buy float64) (*httperrors.HttpError) {
	err1 := r.Productrepository.AUpdate(code, b, old, new,buy)
	return err1
}
func (service productService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Productrepository.Delete(id)
		return success, failure
}
