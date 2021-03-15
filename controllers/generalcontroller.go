package controllers

import(
	"net/http"
	"github.com/labstack/echo"
	"github.com/myrachanto/ecommerce/service"
)
//GeneralController ...
var (
	GeneralController generalController = generalController{}
)
type generalController struct{ }
/////////controllers/////////////////

func (controller generalController) View(c echo.Context) error {
	search := c.QueryParam("search")
	createddashboard, err1 := service.Generalservice.View(search)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusOK, createddashboard)	
}