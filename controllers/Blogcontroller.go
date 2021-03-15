package controllers

import(
	"fmt"	
	"os"
	"io"
	"net/http"
	"github.com/labstack/echo"
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model"
	"github.com/myrachanto/ecommerce/service"
)
 //blogController ...
var (
	BlogController blogController = blogController{}
)
type blogController struct{ }
/////////controllers/////////////////
func (controller blogController) Create(c echo.Context) error {

	blog := &model.Blog{}
	blog.Title = c.FormValue("title")
	blog.Meta = c.FormValue("meta")
	blog.Header1 = c.FormValue("header1")
	blog.Header2 = c.FormValue("header2")
	blog.Intro = c.FormValue("intro")
	blog.Body = c.FormValue("body")
	blog.Summary = c.FormValue("summary")
	// user.Business = c.FormValue("business")
fmt.Println("bleeeeeeeeeeeeeeeeeeeeeeeeeeeee")
	pic, err2 := c.FormFile("picture")
	if pic != nil {
		//    fmt.Println(pic.Filename)
			if err2 != nil {
					httperror := httperrors.NewBadRequestError("Invalid picture")
					return c.JSON(httperror.Code, err2)
				}	
			src, err := pic.Open()
			if err != nil {
				httperror := httperrors.NewBadRequestError("the picture is corrupted")
				return c.JSON(httperror.Code, err)
			}	
			defer src.Close()
			// filePath := "./public/imgs/blogs/"
			filePath := "./public/imgs/blogs/" + pic.Filename
			filePath1 := "/imgs/blogs/" + pic.Filename
			// Destination
			dst, err4 := os.Create(filePath)
			if err4 != nil {
				httperror := httperrors.NewBadRequestError("the Directory mess")
				return c.JSON(httperror.Code, err4)
			}
			defer dst.Close()
			// Copy
			if _, err = io.Copy(dst, src); err != nil {
				if err2 != nil {
					httperror := httperrors.NewBadRequestError("error filling")
					return c.JSON(httperror.Code, httperror)
				}
			}
			
		blog.Picture = filePath1
		// fmt.Println(blog)
		err1 := service.BlogService.Create(blog)
		if err1 != nil {
			return c.JSON(err1.Code, err1)
		}
		if _, err = io.Copy(dst, src); err != nil {
			if err2 != nil {
				httperror := httperrors.NewBadRequestError("error filling")
				return c.JSON(httperror.Code, httperror)
			}
		}
		return c.JSON(http.StatusCreated, "user created succesifully")
	}
	err1 := service.BlogService.Create(blog)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	} 
	return c.JSON(http.StatusCreated, "user created succesifully")
}

func (controller blogController) GetAll(c echo.Context) error {
	search := c.QueryParam("search")
	blogs, err3 := service.BlogService.GetAll(search)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, blogs)
} 
func (controller blogController) GetOne(c echo.Context) error {
	
	code := c.Param("code")
	blog, problem := service.BlogService.GetOne(code)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, blog)	
}

func (controller blogController) Update(c echo.Context) error {
	
	code := c.Param("code")
	blog := &model.Blog{}
	blog.Title = c.FormValue("title")
	blog.Meta = c.FormValue("meta")
	blog.Header1 = c.FormValue("header1")
	blog.Header2 = c.FormValue("header2")
	blog.Intro = c.FormValue("intro")
	blog.Body = c.FormValue("body")
	blog.Summary = c.FormValue("summary")
	// user.Business = c.FormValue("business")

	pic, err2 := c.FormFile("picture")
	if pic != nil {
		//    fmt.Println(pic.Filename)
			if err2 != nil {
					httperror := httperrors.NewBadRequestError("Invalid picture")
					return c.JSON(httperror.Code, err2)
				}	
			src, err := pic.Open()
			if err != nil {
				httperror := httperrors.NewBadRequestError("the picture is corrupted")
				return c.JSON(httperror.Code, err)
			}	
			defer src.Close()
			// filePath := "./public/imgs/blogs/"
			filePath := "./public/imgs/blogs/" + pic.Filename
			filePath1 := "/imgs/blogs/" + pic.Filename
			// Destination
			dst, err4 := os.Create(filePath)
			if err4 != nil {
				httperror := httperrors.NewBadRequestError("the Directory mess")
				return c.JSON(httperror.Code, err4)
			}
			defer dst.Close()
			// Copy
			if _, err = io.Copy(dst, src); err != nil {
				if err2 != nil {
					httperror := httperrors.NewBadRequestError("error filling")
					return c.JSON(httperror.Code, httperror)
				}
			}
			
		blog.Picture = filePath1
		err1 := service.BlogService.Update(code,blog)
		if err1 != nil {
			return c.JSON(err1.Code, err1)
		}
		if _, err = io.Copy(dst, src); err != nil {
			if err2 != nil {
				httperror := httperrors.NewBadRequestError("error filling")
				return c.JSON(httperror.Code, httperror)
			}
		}
		return c.JSON(http.StatusCreated, "user created succesifully")
	}
	err1 := service.BlogService.Update(code,blog)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	} 
	return c.JSON(http.StatusCreated, "user created succesifully")
}

func (controller blogController) Delete(c echo.Context) error {
	id := string(c.Param("id"))
	success, failure := service.BlogService.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}