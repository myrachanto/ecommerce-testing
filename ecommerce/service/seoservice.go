package service

import (
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model" 
	r "github.com/myrachanto/ecommerce/repository"
)

var (
	SeoService  = seoService{}
)

type seoService struct {
}

func (service seoService) Create(seo *model.Seo) (*httperrors.HttpError) {
	err1 := r.Seorepository.Create(seo)
	 return err1

}

func (service seoService) GetOne(code string) (*model.Seo, *httperrors.HttpError) {
	seo, err1 := r.Seorepository.GetOne(code)
	return seo, err1
}

func (service seoService) GetAll(search string) ([]*model.Seo, *httperrors.HttpError) {
	seos, err := r.Seorepository.GetAll(search)
	return seos, err
}

func (service seoService) Update(id string, seo *model.Seo) (*httperrors.HttpError) {
	err1 := r.Seorepository.Update(id, seo)
	return err1
}
func (service seoService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Seorepository.Delete(id)
		return success, failure
}
