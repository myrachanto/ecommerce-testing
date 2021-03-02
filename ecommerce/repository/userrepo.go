package repository

import (
	"fmt"
	"os"
	"log"
	"strconv"
	"github.com/joho/godotenv"
	jwt "github.com/dgrijalva/jwt-go"
    "go.mongodb.org/mongo-driver/bson"
		"github.com/myrachanto/ecommerce/httperrors"
		"github.com/myrachanto/ecommerce/model"
		"go.mongodb.org/mongo-driver/bson/primitive"
)
//Userrepository repository
var (
	Userrepository userrepository = userrepository{}
)

type userrepository struct{}

func (r *userrepository) Create(user *model.User) (*httperrors.HttpError) {
	if err1 := user.Validate(); err1 != nil {
		return err1
	}
	ok, err1 := user.ValidatePassword(user.Password)
	if !ok {
		return err1
	}
	ok = user.ValidateEmail(user.Email)
	if !ok {
		return httperrors.NewNotFoundError("Your email format is wrong!")
	}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	code, err1 := Userrepository.genecode()
	if err1 != nil {
		return err1
	}
	user.Code = code
	
	collection := db.Collection("user")
	result := model.User{}
	filter := bson.M{"email": user.Email}
	errd := collection.FindOne(ctx, filter).Decode(&result)
	if errd == nil {
		return  httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", errd))
	}
	hashpassword, err2 := user.HashPassword(user.Password)
	if err2 != nil {
		return err2
	}
	user.Password = hashpassword
	
	_, err := collection.InsertOne(ctx, user)
		if err != nil {
			return httperrors.NewBadRequestError(fmt.Sprintf("Create user Failed, %d", err))
	}
	DbClose(c)
	return nil
}

func (r *userrepository) Login(user *model.LoginUser) (*model.Auth, *httperrors.HttpError) {
	if err := user.Validate(); err != nil {
		return nil,err
	}
	c, t := Mongoclient();if t != nil {
		return nil,t
	}
	db, e := Mongodb();if e != nil {
		return nil,e
	}
	collection := db.Collection("user")
	filter := bson.M{"email": user.Email,}
	auser := &model.User{}
	err := collection.FindOne(ctx, filter).Decode(&auser)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("User with this email exist @ - , %d", err))
	}
	ok := user.Compare(user.Password, auser.Password)
	if !ok {
		return nil, httperrors.NewNotFoundError("wrong email password combo!")
	}
	tk := &model.Token{
		UserID: auser.Code,
		UName: auser.UName,
		Admin: auser.Admin,
		Supervisor:  auser.Supervisor,
		Employee:  auser.Employee,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: model.ExpiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	err2 := godotenv.Load()
	if err2 != nil {
		log.Fatal("Error loading key")
	}
	encyKey := os.Getenv("EncryptionKey")
	tokenString, error := token.SignedString([]byte(encyKey))
	if error != nil {
		fmt.Println(error)
	}
	auths := &model.Auth{Admin:auser.Admin, UName:auser.UName, Supervisor:auser.Supervisor,Employee:auser.Employee, Token:tokenString}
  // //  fmt.Println(auths)
	// _, err3 := collection.InsertOne(ctx, auths)
	// 	if err3 != nil {
	// 		return nil,httperrors.NewBadRequestError(fmt.Sprintf("Create user Failed, %d", err))
	// }
	
	// filter1 := bson.D{}
	// auth := &model.Auth{}
	// collection2 := db.Collection("auth")
	// err4 := collection2.FindOne(ctx, filter1).Decode(&auth)
	// fmt.Println(filter1)
	// if err4 != nil {
	// 	fmt.Println(err4)
	// 	return nil, httperrors.NewBadRequestError("something went wrong authorizing!")
	// }
	DbClose(c)
	return auths, nil
}
func (r *userrepository) Logout(token string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("auth")
	filter1 := bson.M{"token": token,}
	_, err3 := collection.DeleteOne(ctx, filter1)
	if err3 != nil {
		return nil, httperrors.NewBadRequestError("something went wrong login out!")
	}
	DbClose(c)
 return httperrors.NewSuccessMessage("something went wrong login out!"), nil
}
func (r *userrepository) GetOne(code string) (user *model.User, errors *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("user")
	filter := bson.M{"code": code}
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	DbClose(c)
	return user, nil	
}

func (r *userrepository) GetAll(search string) ([]*model.User, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("user")
			results := []*model.User{}
			fmt.Println(search)
	if search != ""{
	// 	filter := bson.D{
	// 		{"name", primitive.Regex{Pattern: search, Options: "i"}},
	// }
		filter := bson.D{
			{"$or", bson.A{
					bson.D{{"lname", primitive.Regex{Pattern: search, Options: "i"}}},
					bson.D{{"uname", primitive.Regex{Pattern: search, Options: "i"}}},
					bson.D{{"email", primitive.Regex{Pattern: search, Options: "i"}}},
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

func (r *userrepository) Update(code string, user *model.User) (*httperrors.HttpError) {
	uuser := &model.User{}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	
	ok := user.ValidateEmail(user.Email)
	if !ok {
		return httperrors.NewNotFoundError("Your email format is wrong!")
	}
	fmt.Println(code)
	collection := db.Collection("user")
	filter := bson.M{"code": code}
	err := collection.FindOne(ctx, filter).Decode(&uuser)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}

	if user.FName  == "" {
		user.FName = uuser.FName
	}
	if user.LName  == "" {
		user.LName = uuser.LName
	}
	if user.UName  == "" {
		user.UName = uuser.UName
	}
	if user.Phone  == "" {
		user.Phone = uuser.Phone
	}
	if user.Address  == "" {
		user.Address = uuser.Address
	}
	if user.Picture  == "" {
		user.Picture = uuser.Picture
	}
	//////////////////////////////////////////////////////////////////////
	/////////////////get the admin authorisation to handle this///////////////
	if user.Admin  == true {
		user.Admin = uuser.Admin
	}
	if user.Supervisor  == true {
		user.Supervisor = uuser.Supervisor
	}
	if user.Employee  == true {
		user.Employee = uuser.Employee
	}
	if user.Email  == "" {
		user.Email = uuser.Email
	}
	if user.Code == ""{
		user.Code = uuser.Code
	}
	update := bson.M{"$set": user}
	_, errs := collection.UpdateOne(ctx, filter, update)
		if errs != nil {
		return	httperrors.NewNotFoundError("Error updating!")
		}
	DbClose(c)
	return nil
}

func (r *userrepository) AUpdate(id string, user *model.User) (*httperrors.HttpError) {
	uuser := &model.User{}
	c, t := Mongoclient();if t != nil {
		return t
	}
	db, e := Mongodb();if e != nil {
		return e
	}
	ok, err1 := user.ValidatePassword(user.Password)
	if !ok {
		return err1
	}
	ok = user.ValidateEmail(user.Email)
	if !ok {
		return httperrors.NewNotFoundError("Your email format is wrong!")
	}
	hashpassword, err2 := user.HashPassword(user.Password)
	if err2 != nil {
		return err2
	}
	user.Password = hashpassword
	collection := db.Collection("user")
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&uuser)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Could not find resource with this id, %d", err))
	}
	if user.FName  == "" {
		user.FName = uuser.FName
	}
	if user.LName  == "" {
		user.LName = uuser.LName
	}
	if user.UName  == "" {
		user.UName = uuser.UName
	}
	if user.Phone  == "" {
		user.Phone = uuser.Phone
	}
	if user.Address  == "" {
		user.Address = uuser.Address
	}
	if user.Picture  == "" {
		user.Picture = uuser.Picture
	}
	//////////////////////////////////////////////////////////////////////
	/////////////////get the admin authorisation to handle this///////////////
	if !user.Admin {
		user.Admin = uuser.Admin
	}
	if !user.Supervisor {
		user.Supervisor = uuser.Supervisor
	}
	if !user.Employee {
		user.Employee = uuser.Employee
	}
	if user.Email  == "" {
		user.Email = uuser.Email
	}
	if hashpassword == "" {
		user.Password = uuser.Password
	}
	_, err = collection.UpdateOne(ctx, filter, user)
	if err != nil {
		return httperrors.NewBadRequestError(fmt.Sprintf("Update of user Failed, %d", err))
	} 
	DbClose(c)
	return nil
}
func (r userrepository) Delete(id string) (*httperrors.HttpSuccess, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("user")
	filter := bson.M{"_id": id}
	ok, err := collection.DeleteOne(ctx, filter)
	if ok == nil {
		return nil, httperrors.NewNotFoundError(fmt.Sprintf("deletion of %d failed", err))
	}
	DbClose(c)
		return httperrors.NewSuccessMessage("deleted successfully"), nil
	
}
func (r userrepository)genecode()(string, *httperrors.HttpError) {
	c, t := Mongoclient();if t != nil {
		return "", t
	}
	db, e := Mongodb();if e != nil {
		return "", e
	}
	collection := db.Collection("user")
	filter := bson.M{}
	count, err := collection.CountDocuments(ctx, filter)
	co := count + 1
	if err != nil { 
		return "",	httperrors.NewNotFoundError("no results found")
	}
	code := "UserCode"+strconv.FormatUint(uint64(co), 10)

	DbClose(c)
	return code, nil
}
func (r userrepository)getuno(code string)(result *model.User, err *httperrors.HttpError){
	c, t := Mongoclient();if t != nil {
		return nil, t
	}
	db, e := Mongodb();if e != nil {
		return nil, e
	}
	collection := db.Collection("user")
	filter := bson.M{"code": code}
	err1 := collection.FindOne(ctx, filter).Decode(&result)
	if err1 != nil {
		return nil, httperrors.NewNotFoundError("no results found")
	}
	DbClose(c)
	return result, nil	
}
func (r userrepository)emailexist(email string)bool{
	c, t := Mongoclient();if t != nil {
		return false
	}
	db, e := Mongodb();if e != nil {
		return false
	}
	collection := db.Collection("user")
	result := model.User{}
	filter := bson.M{"email": email}
	err1 := collection.FindOne(ctx, filter).Decode(&result)
	if err1 != nil {
		return false
	}
	DbClose(c)
	return true	
}
func (r userrepository)Count()(float64, *httperrors.HttpError) {
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