package service

import (
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model" 
	r "github.com/myrachanto/ecommerce/repository"
)

var (
	BlogService  = blogService{}
)
type BlogInterface interface{
	Create(blog *model.Blog) (*httperrors.HttpError)
	GetOne(id string) (*model.Blog, *httperrors.HttpError)
	GetAll(search string) ([]*model.Blog, *httperrors.HttpError)
	Update(code string, blog *model.Blog) (*httperrors.HttpError)
	Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError)
}
type blogService struct {
}
func NewBlogService(repository r.BlogInterface) BlogInterface {
	return &BlogService
}

func (service blogService) Create(blog *model.Blog) (*httperrors.HttpError) {
	err1 := r.Blogrepository.Create(blog)
	 return err1

}

func (service blogService) GetOne(code string) (*model.Blog, *httperrors.HttpError) {
	blog, err1 := r.Blogrepository.GetOne(code)
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
func (service blogService) Delete(code string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
		success, failure := r.Blogrepository.Delete(code)
		return success, failure
}
