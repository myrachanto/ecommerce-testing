package repository

import (
	"fmt"
	"strconv"
    "go.mongodb.org/mongo-driver/bson"
		"github.com/myrachanto/ecommerce/httperrors"
		"github.com/myrachanto/ecommerce/model" 
		"go.mongodb.org/mongo-driver/bson/primitive"
)
var (
	Blogrepository blogrepository = blogrepository{}
)

type blogrepository struct{}

func (r *blogrepository) Create(blog * model.Blog) (*httperrors.HttpError) {
	if err1 := blog.Validate(); err1 != nil {
		return err1
	}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	code, err1 := Blogrepository.genecode()
	if err1 != nil {
		return err1
	}
	blog.Code = code
	collection := db.Collection("blog")
	_, err := collection.InsertOne(ctx, blog)
		if err != nil {
			return httperrors.NewBadRequestError(fmt.Sprintf("Create blog Failed, %d", err))
	}
	DbClose(c)
	return nil
}

func (r *blogrepository) GetOne(code string) (blog * model.Blog, errors *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("blog")
	filter := bson.M{"code": code}
	err := collection.FindOne(ctx, filter).Decode(&blog)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	DbClose(c)
	return blog, nil	
}

func (r *blogrepository) GetAll(search string) ([]*model.Blog, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("blog")
			results := []*model.Blog{}
			fmt.Println(search)
	if search != ""{
	// 	filter := bson.D{
	// 		{"name", primitive.Regex{Pattern: search, Options: "i"}},
	// }
		filter := bson.D{
			{"$or", bson.A{
					bson.D{{"meta", primitive.Regex{Pattern: search, Options: "i"}}},
					bson.D{{"title", primitive.Regex{Pattern: search, Options: "i"}}},
					bson.D{{"header1", primitive.Regex{Pattern: search, Options: "i"}}},
			}},
	}
		// fmt.Println(filter)
				cursor, err := collection.Find(ctx, filter)
				fmt.Println(cursor)
				if err != nil {
					return nil, httperrors.NewNotFoundError("No records found!") 
				}
				if err = cursor.All(ctx, &results); err != nil {
					return nil, httperrors.NewNotFoundError("Error decoding!") 
				}
				DbClose(c)
				fmt.Println(results)
				return results, nil
		}else{
			cursor, err := collection.Find(ctx, bson.M{})
			if err != nil {
				return nil, httperrors.NewNotFoundError("No records found!") 
			}
			if err = cursor.All(ctx, &results); err != nil {
				return nil, httperrors.NewNotFoundError("Error decoding!") 
			}
			DbClose(c)
			return results, nil
		}

}

func (r *blogrepository) Update(code string, blog * model.Blog) (*httperrors.HttpError) {
	ublog := & model.Blog{}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("blog")
	filter := bson.M{"code": code}
	err := collection.FindOne(ctx, filter).Decode(&ublog)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	if blog.Title  == "" {
		blog.Title = ublog.Title
	}
	if blog.Meta  == "" {
		blog.Meta = ublog.Meta
	}

	if blog.Header1  == "" {
		blog.Header1 = ublog.Header1
	}
	if blog.Header2  == "" {
		blog.Header2 = ublog.Header2
	}
	if blog.Intro  == "" {
		blog.Intro = ublog.Intro
	}
	if blog.Body  == "" {
		blog.Body = ublog.Body
	}
	if blog.Summary  == "" {
		blog.Summary = ublog.Summary
	}
	if blog.Code  == "" {
		blog.Code = ublog.Code
	}
	update := bson.M{"$set": blog}
	_, errs := collection.UpdateOne(ctx, filter, update)
		if errs != nil {
		return	httperrors.NewNotFoundError("Error updating!")
		}
	DbClose(c)
	return nil
}
func (r blogrepository) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("blog")
	filter := bson.M{"_id": id}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return nil, httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	DbClose(c)
		return httperrors.NewSuccessMessage("deleted successfully"), nil
	
}
func (r blogrepository)genecode()(string, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return "", t
	}
	db, e := Mongodb();if e != nil {
		return "", e
	}
	collection := db.Collection("category")
	filter := bson.M{}
	count, err := collection.CountDocuments(ctx, filter)
	co := count + 1
	if err != nil { 
		return "",	httperrors.NewNotFoundError("no results found")
	}
	code := "BlogCode"+strconv.FormatUint(uint64(co), 10)

	DbClose(c)
	return code, nil
}
func (r blogrepository)Count()(float64, *httperrors.HttpError) {
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