package routes

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/myrachanto/ecommerce/controllers"
)
// ApiMicroservice ...
func ApiMicroservice() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file in routes")
	}
	PORT := os.Getenv("PORT")
	key := os.Getenv("EncryptionKey")
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) 
	e.Use(middleware.CORS())

	e.Static("/", "public")
	JWTgroup := e.Group("/api")
	JWTgroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey: []byte(key),
	}))
	// admin := e.Group("admin/")
	// admin.Use(isAdmin)

	// var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningMethod: "HS256",
	// 	SigningKey: []byte(key),
	// })
	//JwtG := e.Group("/users")
	// JwtG.Use(middleware.JWT([]byte(key)))
	// Routes
	//e.GET("/is-loggedin", h.private, IsLoggedIn)
	//e.POST("/register", IsLoggedIn,isAdmin,isEmployee,isSupervisor, controllers.UserController.Create)
	// Routes
	////////////////////////////////////////////////////////
	/////////////users//////////////////////////////////

	////////////////////////////////////////////////////////
	/////////////front side//////////////////////////////////
	e.GET("/blogi",  controllers.BlogController.GetAll)
	e.GET("/blogi/:code",  controllers.BlogController.GetOne)
	e.GET("/products",  controllers.ProductController.GetAll)
	e.GET("/productlimit",  controllers.ProductController.GetThree)
	e.GET("/product/:code",  controllers.ProductController.GetOne)
	e.POST("/register", controllers.UserController.Create)
	e.POST("/login", controllers.UserController.Login)
	////////////////////////////////////////////////////////
	/////////////users//////////////////////////////////
	JWTgroup.GET("/logout", controllers.UserController.Logout)
	JWTgroup.GET("/users", controllers.UserController.GetAll)
	JWTgroup.GET("/users/:code", controllers.UserController.GetOne)
	JWTgroup.PUT("/users/:code", controllers.UserController.Update)
	JWTgroup.PUT("/usersA/:code", controllers.UserController.AUpdate)
	JWTgroup.DELETE("/users/:code", controllers.UserController.Delete)
	////////////////////////////////////////////////////////
	/////////////general//////////////////////////////////
	JWTgroup.GET("/dashboard", controllers.GeneralController.View)
	////////////////////////////////////////////////////////
	/////////////category//////////////////////////////////
	JWTgroup.POST("/categorys", controllers.CategoryController.Create)
	JWTgroup.GET("/categorys", controllers.CategoryController.GetAll)
	JWTgroup.GET("/categorys/:id", controllers.CategoryController.GetOne)
	JWTgroup.PUT("/categorys/:id", controllers.CategoryController.Update)
	JWTgroup.DELETE("/categorys/:id", controllers.CategoryController.Delete)
	////////////////////////////////////////////////////////
	/////////////category//////////////////////////////////
	JWTgroup.POST("/countys", controllers.CountyController.Create)
	JWTgroup.GET("/countys", controllers.CountyController.GetAll)
	JWTgroup.GET("/countys/:id", controllers.CountyController.GetOne)
	JWTgroup.PUT("/countys/:id", controllers.CountyController.Update)
	JWTgroup.DELETE("/countys/:id", controllers.CountyController.Delete)
	////////////////////////////////////////////////////////
	/////////////category//////////////////////////////////
	JWTgroup.POST("/customers", controllers.CustomerController.Create)
	JWTgroup.GET("/customers", controllers.CustomerController.GetAll)
	JWTgroup.GET("/customers/:id", controllers.CustomerController.GetOne)
	JWTgroup.PUT("/customers/:id", controllers.CustomerController.Update)
	JWTgroup.DELETE("/customers/:id", controllers.CustomerController.Delete)
	////////////////////////////////////////////////////////
	/////////////category//////////////////////////////////
	JWTgroup.POST("/divisions", controllers.DivisionController.Create)
	JWTgroup.GET("/divisions", controllers.DivisionController.GetAll)
	JWTgroup.GET("/divisions/:id", controllers.DivisionController.GetOne)
	JWTgroup.PUT("/divisions/:id", controllers.DivisionController.Update)
	JWTgroup.DELETE("/divisions/:id", controllers.DivisionController.Delete)
	////////////////////////////////////////////////////////
	/////////////category//////////////////////////////////
	JWTgroup.POST("/industrys", controllers.IndustryController.Create)
	JWTgroup.GET("/industrys", controllers.IndustryController.GetAll)
	JWTgroup.GET("/industrys/:id", controllers.IndustryController.GetOne)
	JWTgroup.PUT("/industrys/:id", controllers.IndustryController.Update)
	JWTgroup.DELETE("/industrys/:id", controllers.IndustryController.Delete)
	////////////////////////////////////////////////////////
	/////////////category//////////////////////////////////
	JWTgroup.POST("/majorcategorys", controllers.MajorcategoryController.Create)
	JWTgroup.GET("/majorcategorys", controllers.MajorcategoryController.GetAll)
	JWTgroup.GET("/majorcategorys/:code", controllers.MajorcategoryController.GetOne)
	JWTgroup.PUT("/majorcategorys/:code", controllers.MajorcategoryController.Update)
	JWTgroup.DELETE("/majorcategorys/:code", controllers.MajorcategoryController.Delete)
	////////////////////////////////////////////////////////
	/////////////category//////////////////////////////////
	JWTgroup.POST("/subcategorys", controllers.SubcategoryController.Create)
	JWTgroup.GET("/subcategorys", controllers.SubcategoryController.GetAll)
	JWTgroup.GET("/subcategorys/:id", controllers.SubcategoryController.GetOne)
	JWTgroup.PUT("/subcategorys/:id", controllers.SubcategoryController.Update)
	JWTgroup.DELETE("/subcategorys/:id", controllers.SubcategoryController.Delete)
	////////////////////////////////////////////////////////
	/////////////category//////////////////////////////////
	JWTgroup.POST("/categorys", controllers.CategoryController.Create)
	JWTgroup.GET("/categorys/major/:majorcode", controllers.CategoryController.GetAll)
	JWTgroup.GET("/categorys/:code", controllers.CategoryController.GetOne)
	JWTgroup.PUT("/categorys/:code", controllers.CategoryController.Update)
	JWTgroup.DELETE("/categorys/:code", controllers.CategoryController.Delete)
	////////////////////////////////////////////////////////
	/////////////town//////////////////////////////////
	JWTgroup.POST("/towns", controllers.TownController.Create)
	JWTgroup.GET("/towns", controllers.TownController.GetAll)
	JWTgroup.GET("/towns/:id", controllers.TownController.GetOne)
	JWTgroup.PUT("/towns/:id", controllers.TownController.Update)
	JWTgroup.DELETE("/towns/:id", controllers.TownController.Delete)
	////////////////////////////////////////////////////////
	/////////////Shop//////////////////////////////////
	JWTgroup.POST("/seos", controllers.SeoController.Create)
	JWTgroup.GET("/seos", controllers.SeoController.GetAll)
	JWTgroup.GET("/seos/:code", controllers.SeoController.GetOne)
	JWTgroup.PUT("/seos/:code", controllers.SeoController.Update)
	JWTgroup.DELETE("/seos/:code", controllers.SeoController.Delete)
	////////////////////////////////////////////////////////
	/////////////products//////////////////////////////////
	JWTgroup.POST("/products", controllers.ProductController.Create)
	JWTgroup.GET("/products", controllers.ProductController.GetAll)
	JWTgroup.GET("/products/:code", controllers.ProductController.GetOne)
	JWTgroup.PUT("/products/:code", controllers.ProductController.Update)
	JWTgroup.PUT("/inventory/:code", controllers.ProductController.AUpdate)
	JWTgroup.DELETE("/products/:code", controllers.ProductController.Delete)
	////////////////////////////////////////////////////////
	/////////////tags//////////////////////////////////
	JWTgroup.POST("/blogs", controllers.BlogController.Create)
	JWTgroup.GET("/blogs", controllers.BlogController.GetAll)
	JWTgroup.GET("/blogs/:code", controllers.BlogController.GetOne)
	JWTgroup.PUT("/blogs/:code", controllers.BlogController.Update)
	JWTgroup.DELETE("/blogs/:code", controllers.BlogController.Delete)
	////////////////////////////////////////////////////////
	/////////////ratings//////////////////////////////////
	JWTgroup.POST("/ratings", controllers.RatingController.Create)
	JWTgroup.GET("/ratings", controllers.RatingController.GetAll)
	JWTgroup.GET("/ratings/:id", controllers.RatingController.GetOne)
	JWTgroup.PUT("/ratings/:id", controllers.RatingController.Update)
	JWTgroup.DELETE("/ratings/:id", controllers.RatingController.Delete)
	////////////////////////////////////////////////////////
	/////////////nortificatrions//////////////////////////////////
	JWTgroup.POST("/nortificatrions", controllers.NortificationController.Create)
	JWTgroup.GET("/nortificatrions", controllers.NortificationController.GetAll)
	JWTgroup.GET("/nortificatrions/:id", controllers.NortificationController.GetOne)
	JWTgroup.PUT("/nortificatrions/:id", controllers.NortificationController.Update)
	JWTgroup.DELETE("/nortificatrions/:id", controllers.NortificationController.Delete)
	////////////////////////////////////////////////////////
	/////////////verify//////////////////////////////////
	JWTgroup.POST("/verify", controllers.VerifyController.Create)
	JWTgroup.GET("/verify", controllers.VerifyController.GetAll)
	JWTgroup.GET("/verify/:id", controllers.VerifyController.GetOne)
	JWTgroup.PUT("/verify/:id", controllers.VerifyController.Update)
	JWTgroup.DELETE("/verify/:id", controllers.VerifyController.Delete)
	////////////////////////////////////////////////////////
	/////////////invoice//////////////////////////////////
	e.POST("/invoice", controllers.InvoiceController.Create)
	e.GET("/invoice", controllers.InvoiceController.GetAll)
	e.GET("/invoice/:id", controllers.InvoiceController.GetOne)
	// e.PUT("/invoice/:id", controllers.InvoiceController.Update)
	// e.DELETE("/invoice/:id", controllers.InvoiceController.Delete)

	// Start server
	e.Logger.Fatal(e.Start(PORT))
}
