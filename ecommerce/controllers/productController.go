package controllers

import(
	"fmt"	
	"os"
	"io"
	"net/http"
	"strconv"
	"github.com/labstack/echo"
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model"
	"github.com/myrachanto/ecommerce/service"
)
 //ProductController ..
var (
	ProductController productController = productController{}
)
type productController struct{ }
/////////controllers/////////////////
func (controller productController) Create(c echo.Context) error {
	product := &model.Product{}
	
	product.Name = c.FormValue("name")
	product.Description = c.FormValue("description")
	product.Title = c.FormValue("title")
	product.Category = c.FormValue("category")
	product.Majorcategory = c.FormValue("majorcategory")
	b, err := strconv.ParseFloat(c.FormValue("oldprice"), 64)
	if err != nil {
		httperror := httperrors.NewBadRequestError("Invalid buying price")
		return c.JSON(httperror.Code, httperror)
	}
	e, err2 := strconv.ParseFloat(c.FormValue("newprice"), 64)
	if err2 != nil {
		httperror := httperrors.NewBadRequestError("Invalid buying price")
		return c.JSON(httperror.Code, httperror)
	}
	d, err4 := strconv.ParseFloat(c.FormValue("buyprice"), 64)
	if err4 != nil {
		httperror := httperrors.NewBadRequestError("Invalid buying price")
		return c.JSON(httperror.Code, httperror)
	}

	product.Oldprice = b
	product.Newprice = e
	product.Buyprice = d

	pic, err2 := c.FormFile("picture")
			if pic != nil{
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
		filePath := "./public/imgs/products/" + pic.Filename
		// Destination
		dst, err4 := os.Create(filePath)
		if err4 != nil {
			httperror := httperrors.NewBadRequestError("the Directory mess")
			return c.JSON(httperror.Code, err4)
		}
		defer dst.Close()
		//  copy
		if _, err = io.Copy(dst, src); err != nil {
			if err2 != nil {
				httperror := httperrors.NewBadRequestError("error filling")
				return c.JSON(httperror.Code, httperror)
			}
		} 
		
		product.Picture = pic.Filename
		err1 := service.ProductService.Create(product)
		if err1 != nil {
		return c.JSON(err1.Code, err1)
		} 
		return c.JSON(http.StatusCreated, "created successifuly")
		}
	err1 := service.ProductService.Create(product)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, "created successifuly")
}

func (controller productController) GetAll(c echo.Context) error {
	search := c.QueryParam("search")
	products, err3 := service.ProductService.GetAll(search)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, products)
} 
func (controller productController) GetOne(c echo.Context) error {
	code := c.Param("code")
	product, problem := service.ProductService.GetOne(code)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, product)	
}

func (controller productController) Update(c echo.Context) error {
	product :=  &model.Product{}
	product.Name = c.FormValue("name")
	product.Description = c.FormValue("description")
	product.Title = c.FormValue("title")
	product.Category = c.FormValue("category")
	product.Majorcategory = c.FormValue("majorcategory")
	b, err := strconv.ParseFloat(c.FormValue("oldprice"), 64)
	if err != nil {
		httperror := httperrors.NewBadRequestError("Invalid buying price")
		return c.JSON(httperror.Code, httperror)
	}
	e, err2 := strconv.ParseFloat(c.FormValue("newprice"), 64)
	if err2 != nil {
		httperror := httperrors.NewBadRequestError("Invalid buying price")
		return c.JSON(httperror.Code, httperror)
	}
	d, err4 := strconv.ParseFloat(c.FormValue("buyprice"), 64)
	if err4 != nil {
		httperror := httperrors.NewBadRequestError("Invalid buying price")
		return c.JSON(httperror.Code, httperror)
	}

	product.Oldprice = b
	product.Newprice = e
	product.Buyprice = d
	code := c.Param("code")
	// fmt.Println(pcode, "sssssssssssssssssssssssssssssssssss")
	pic, err2 := c.FormFile("picture")
			if pic != nil{
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
		filePath := "./public/imgs/products/" + pic.Filename
		// Destination
		dst, err4 := os.Create(filePath)
		if err4 != nil {
			httperror := httperrors.NewBadRequestError("the Directory mess")
			return c.JSON(httperror.Code, err4)
		}
		defer dst.Close()
		//  copy
		if _, err = io.Copy(dst, src); err != nil {
			if err2 != nil {
				httperror := httperrors.NewBadRequestError("error filling")
				return c.JSON(httperror.Code, httperror)
			}
		} 
		
		product.Picture = pic.Filename
		fmt.Println(product)
		fmt.Println(code, "----==============================")

		err1 := service.ProductService.Update(code,product)
		if err1 != nil {
		return c.JSON(err1.Code, err1)
		} 
		return c.JSON(http.StatusOK, "update successifuly")
		}
	problem := service.ProductService.Update(code, product)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusCreated, "Updated successifuly")
}

func (controller productController) AUpdate(c echo.Context) error {
	b, err := strconv.ParseFloat(c.FormValue("quantity"), 64)
	if err != nil {
		httperror := httperrors.NewBadRequestError("Invalid buying price")
		return c.JSON(httperror.Code, httperror)
	}
	
	code := c.Param("code")
	fmt.Println(code, b)
	problem := service.ProductService.AUpdate(code,b)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusCreated, "Updated successifuly")
}

func (controller productController) Delete(c echo.Context) error {
	id := string(c.Param("id"))
	success, failure := service.ProductService.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}