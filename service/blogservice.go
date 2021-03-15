package service

import (
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model" 
	r "github.com/myrachanto/ecommerce/repository"
)

var (
	BlogService  = blogService{}
)

type blogService struct {
}

func (service blogService) Create(blog *model.Blog) (*httperrors.HttpError) {
	err1 := r.Blogrepository.Create(blog)
	 return err1

}

func (service blogService) GetOne(id string) (*model.Blog, *httperrors.HttpError) {
	blog, err1 := r.Blogrepository.GetOne(id)
	return blog, err1
}

func (service blogService) GetAll(search string) ([]*model.Blog, *httperrors.HttpError) {
	blogs, err := r.Blogrepository.GetAll(search)
	return blogs, err
}

func (service blogService) Update(code string, blog *model.Blog) (*httperrors.HttpError) {
	err1 := r.Blogrepository.Update(code, blog)
	return err1
}
func (service blogService) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Blogrepository.Delete(id)
		return success, failure
}
