package controllers

import (
	//"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model"
	"github.com/myrachanto/ecommerce/service"
)

//SeoController ..
var (
	SeoController seoController = seoController{}
)
type seoController struct{ }
/////////controllers/////////////////
func (controller seoController) Create(c echo.Context) error {
	seo := &model.Seo{}
	if err := c.Bind(seo); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	err1 := service.SeoService.Create(seo)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, "created successifuly")
}

func (controller seoController) GetAll(c echo.Context) error {
	search := c.QueryParam("search")
	seos, err3 := service.SeoService.GetAll(search)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, seos)
} 
func (controller seoController) GetOne(c echo.Context) error {
	code := c.Param("code")
	seo, problem := service.SeoService.GetOne(code)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, seo)	
}

func (controller seoController) Update(c echo.Context) error {
	seo :=  &model.Seo{}
	if err := c.Bind(seo); err != nil {
		httperror := httperrors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	id := string(c.Param("id"))
	problem := service.SeoService.Update(id, seo)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusCreated, "updated successifuly")
}

func (controller seoController) Delete(c echo.Context) error {
	id := string(c.Param("id"))
	success, failure := service.SeoService.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}