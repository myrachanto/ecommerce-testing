package service

import (
	// "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/myrachanto/ecommerce/model" 
	"github.com/myrachanto/ecommerce/httperrors"
)
type MockRepository struct {
	mock.Mock
}
func (mock MockRepository)Create(blog *model.Blog) (*model.Blog,*httperrors.HttpError){
	args := mock.Called()
	result := args.Get(0)
	blog, err := result.(*model.Blog), args.Error(1)
	if blog.Title == "" {
		return nil, httperrors.NewNotFoundError("test failed blog title empty")
	}
	if err != nil {
		return nil, httperrors.NewNotFoundError("test failed")
	}
	return blog, nil
}
func (mock MockRepository)GetOne(id string) (*model.Blog, *httperrors.HttpError){
	args := mock.Called()
	result := args.Get(0)
	blog, err := result.(*model.Blog), args.Error(1)
	if err != nil {
		return nil, httperrors.NewNotFoundError("test failed")
	}
	return blog,nil
}
func (mock MockRepository)GetAll() ([]*model.Blog, *httperrors.HttpError){
	args := mock.Called()
	result := args.Get(0)
	blogs, err := result.([]*model.Blog), args.Error(1)
	if err != nil {
		return nil, httperrors.NewNotFoundError("test failed")
	}
	return blogs, nil
}
func (mock MockRepository)Update(code string, blog *model.Blog) (*httperrors.HttpError){
	args := mock.Called()
	result := args.Get(0)
	blog, err := result.(*model.Blog), args.Error(1)
	if blog.Title == "" {
		return httperrors.NewNotFoundError("test failed blog title empty")
	}
	if err != nil {
		return httperrors.NewNotFoundError("test failed")
	}
	return nil
}
func TestGetAll(t *testing.T){
	mockRepo := new(MockRepository)
	blog := &model.Blog{
		Title: "title",
		Meta: "meta",
		Intro: "intro",
	}
	//set up expecctations
	mockRepo.On("GetAll").Return([]*model.Blog{blog} ,nil)
	results, _ := mockRepo.GetAll()
	//mock assertion: behavioral
	mockRepo.AssertExpectations(t)
	//data assertion
	assert.Equal(t, "titles", results[0].Title)
	assert.Equal(t, "meta", results[0].Meta)
	assert.Equal(t, "intro", results[0].Intro)


}
func TestCreate(t *testing.T){
	mockRepo := new(MockRepository)
	blog := model.Blog{
		Title: "title",
		Meta: "meta",
		Intro: "intro",}
		mockRepo.On("Create").Return(&blog, nil)
		result, err := mockRepo.Create(&blog)
		//mock assertion: behavioral
		mockRepo.AssertExpectations(t)
		//data assertion
		assert.Equal(t, "title", result.Title)
		assert.Equal(t, "meta", result.Meta)
		assert.Equal(t, "intro", result.Intro)
		assert.Nil(t, err)

}
func TestTitleValidate(t *testing.T){ 
	blog := model.Blog{}
	blog.Title = ""
	err := blog.Validate()
	expected := "Invalid title"
	// if err.Message != expected {
	// 	t.Error(err.Message)
	// }
	assert.NotNil(t, err)
	assert.Equal(t, expected, err.Message)
}
func TestMetaValidate(t *testing.T){
	blog := model.Blog{}
	blog.Title = "Title"
	blog.Meta = ""
	err := blog.Validate()
	expected := "Invalid meta"
	// if err.Message != expected {
	// 	t.Error(err.Message)
	// }
	assert.NotNil(t, err)
	assert.Equal(t, expected, err.Message)
}