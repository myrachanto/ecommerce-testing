package repository

import (
	"fmt"
	"strconv"

	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//productrepository ...
var (
	Productrepository productrepository = productrepository{}
)

type productrepository struct{}

func (r *productrepository) Create(product *model.Product) (*httperrors.HttpError) {
	if err1 := product.Validate(); err1 != nil {
		return err1
	}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	code, err1 := Productrepository.genecode()
	if err1 != nil {
		return err1
	}
	product.Code = code
	collection := db.Collection("product")
	_, err := collection.InsertOne(ctx, product)
		if err != nil {
			return httperrors.NewBadRequestError(fmt.Sprintf("Create product Failed, %d", err))
	}
	DbClose(c)
	return nil
}

func (r *productrepository) GetOne(code string) (product *model.Product, errors *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("product")
	filter := bson.M{"code": code}
	err := collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	DbClose(c)
	return product, nil	
}

func (r *productrepository) GetAll(search string) ([]*model.Product, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("product")
			products := []*model.Product{}
			fmt.Println(search)
	if search != ""{
	// 	filter := bson.D{
	// 		{"name", primitive.Regex{Pattern: search, Options: "i"}},
	// }
		filter := bson.D{
			{"$or", bson.A{
					bson.D{{"name", primitive.Regex{Pattern: search, Options: "i"}}},
					bson.D{{"title", primitive.Regex{Pattern: search, Options: "i"}}},
					bson.D{{"description", primitive.Regex{Pattern: search, Options: "i"}}},
			}},
	}
		// fmt.Println(filter)
				cursor, err := collection.Find(ctx, filter)
				fmt.Println(cursor)
				if err != nil {
					return nil, httperrors.NewNotFoundError("No records found!") 
				}
				if err = cursor.All(ctx, &products); err != nil {
					return nil, httperrors.NewNotFoundError("Error decoding!") 
				}
				DbClose(c)
				fmt.Println(products)
				return products, nil
		}else{
			cursor, err := collection.Find(ctx, bson.M{})
			if err != nil {
				return nil, httperrors.NewNotFoundError("No records found!") 
			}
			if err = cursor.All(ctx, &products); err != nil {
				return nil, httperrors.NewNotFoundError("Error decoding!") 
			}
			DbClose(c)
			return products, nil
		}

}

func (r *productrepository) Update(code string, product *model.Product) (*httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	result, err3 := Productrepository.getuno(code)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(result)
	if product.Name == ""{
		product.Name = result.Name
	}
	if product.Title == ""{
		product.Title = result.Title
	}
	if product.Description == ""{
		product.Description = result.Description
	}
	if product.Code == ""{
		product.Code = result.Code
	}

	if product.Majorcategory == ""{
		product.Majorcategory = result.Majorcategory
	}
	if product.Category == ""{
		product.Category = result.Category
	}
	if product.Oldprice == 0{
		product.Oldprice = result.Oldprice
	}
	if product.Newprice == 0{
		product.Newprice = result.Newprice
	}
	if product.Buyprice == 0{
		product.Buyprice = result.Buyprice
	}

	if product.Quantity == 0{
		product.Quantity = result.Quantity
	}
	collection := db.Collection("product")
	filter := bson.M{"code": code}
	fmt.Println(filter)
	fmt.Println(product)
	update := bson.M{"$set": product}
	_, err := collection.UpdateOne(ctx, filter, update)
		if err != nil {
		return	httperrors.NewNotFoundError("Error updating!") 
		}
	DbClose(c)
	return nil
}


func (r *productrepository) AUpdate(code string, b float64) (*httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	result, err3 := Productrepository.getuno(code)
	if err3 != nil {
		fmt.Println(err3)
	}
	product := &model.Product{}
	product.Quantity = b
	fmt.Println(result)
	if product.Name == ""{
		product.Name = result.Name
	}
	if product.Title == ""{
		product.Title = result.Title
	}
	if product.Description == ""{
		product.Description = result.Description
	}
	if product.Code == ""{
		product.Code = result.Code
	}

	if product.Majorcategory == ""{
		product.Majorcategory = result.Majorcategory
	}
	if product.Category == ""{
		product.Category = result.Category
	}
	if product.Oldprice == 0{
		product.Oldprice = result.Oldprice
	}
	if product.Newprice == 0{
		product.Newprice = result.Newprice
	}
	if product.Buyprice == 0{
		product.Buyprice = result.Buyprice
	}

	if product.Quantity == 0{
		product.Quantity = result.Quantity
	}
	collection := db.Collection("product")
	filter := bson.M{"code": code}
	fmt.Println(filter)
	fmt.Println(product)
	update := bson.M{"$set": product}
	_, err := collection.UpdateOne(ctx, filter, update)
		if err != nil {
		return	httperrors.NewNotFoundError("Error updating!") 
		}
	DbClose(c)
	return nil
}
func (r productrepository) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("product")
	filter := bson.M{"_id": id}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return nil, httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	DbClose(c)
		return httperrors.NewSuccessMessage("deleted successfully"), nil
}
func (r productrepository)genecode()(string, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return "", t
	}
	db, e := Mongodb();if e != nil {
		return "", e
	}
	collection := db.Collection("product")
	filter := bson.M{}
	count, err := collection.CountDocuments(ctx, filter)
	co := count + 1
	if err != nil { 
		return "",	httperrors.NewNotFoundError("no results found")
	}
	code := "ProductCode"+strconv.FormatUint(uint64(co), 10)

	DbClose(c)
	return code, nil
}
func (r productrepository)getuno(code string)(result *model.Product, err *httperrors.HttpError){
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("product")
	filter := bson.M{"code": code}
	err1 := collection.FindOne(ctx, filter).Decode(&result)
	if err1 != nil {
		return nil, httperrors.NewNotFoundError("no results found")
	}
	DbClose(c)
	return result, nil	
}
func (r productrepository)Count()(float64, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return 0, t
	}
	db, e := Mongodb();if e != nil {
		return 0, e
	}
	collection := db.Collection("product")
	filter := bson.M{}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil { 
		return 0,	httperrors.NewNotFoundError("no results found")
	}
	code := float64(count)

	DbClose(c)
	return code, nil
}