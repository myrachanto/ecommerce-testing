package repository

import (
	"fmt"
	"strconv"
    "go.mongodb.org/mongo-driver/bson"
		"github.com/myrachanto/ecommerce/httperrors"
		"github.com/myrachanto/ecommerce/model" 
		"go.mongodb.org/mongo-driver/bson/primitive"
)
//Seorepository ...
var (
	Seorepository seorepository = seorepository{}
)

type seorepository struct{}

func (r *seorepository) Create(seo * model.Seo) (*httperrors.HttpError) {
	if err1 := seo.Validate(); err1 != nil {
		return err1
	}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	code, err1 := Seorepository.genecode()
	if err1 != nil {
		return err1
	}
	seo.Code = code
	collection := db.Collection("seo")
	_, err := collection.InsertOne(ctx, seo)
		if err != nil {
			return httperrors.NewBadRequestError(fmt.Sprintf("Create seo Failed, %d", err))
	}
	DbClose(c)
	return nil
}

func (r *seorepository) GetOne(code string) (seo * model.Seo, errors *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("seo")
	filter := bson.M{"code": code}
	err := collection.FindOne(ctx, filter).Decode(&seo)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	DbClose(c)
	return seo, nil	
}

func (r *seorepository) GetAll(search string) ([]*model.Seo, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("seo")
			results := []*model.Seo{}
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

func (r *seorepository) Update(code string, seo * model.Seo) (*httperrors.HttpError) {
	useo := & model.Seo{}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	collection := db.Collection("seo")
	filter := bson.M{"code": code}
	err := collection.FindOne(ctx, filter).Decode(&useo)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	if seo.Title  == "" {
		seo.Title = useo.Title
	}
	if seo.Meta  == "" {
		seo.Meta = useo.Meta
	}

	if seo.Header1  == "" {
		seo.Header1 = useo.Header1
	}
	if seo.Header2  == "" {
		seo.Header2 = useo.Header2
	}
	if seo.Code  == "" {
		seo.Code = useo.Code
	}
	update := bson.M{"$set": seo}
	_, errs := collection.UpdateOne(ctx, filter, update)
		if errs != nil {
		return	httperrors.NewNotFoundError("Error updating!")
		}
	DbClose(c)
	return nil
}
func (r seorepository) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("seo")
	filter := bson.M{"_id": id}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return nil, httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	DbClose(c)
		return httperrors.NewSuccessMessage("deleted successfully"), nil
	
}
func (r seorepository)genecode()(string, *httperrors.HttpError) {
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
	code := "SeoCode"+strconv.FormatUint(uint64(co), 10)

	DbClose(c)
	return code, nil
}
