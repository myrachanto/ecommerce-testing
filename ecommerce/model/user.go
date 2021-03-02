package model

import (
	"regexp"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/myrachanto/ecommerce/httperrors"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ExpiresAt = time.Now().Add(time.Minute * 100000).Unix()

type User struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FName         string             `bson:"fname" json:"f_name,omitempty"`
	LName         string             `bson:"lname" json:"l_name,omitempty"`
	UName         string             `bson:"uname" json:"u_name,omitempty"`
	Phone         string             `bson:"phone" json:"phone,omitempty"`
	Address       string             `bson:"address" json:"address,omitempty"`
	Dob           *time.Time         `json:"dob,omitempty"`
	Picture       string             `bson:"picture" json:"picture,omitempty"`
	Email         string             `bson:"email" json:"email,omitempty"`
	Password      string             `bson:"password" json:"password,omitempty"`
	Code          string             `json:"code,omitempty"`
	Admin         bool               `json:"admin,omitempty"`
	Supervisor    bool               `json:"supervisor,omitempty"`
	Employee      bool               `json:"employee,omitempty"`
	Location      Location           `json:"location,omitempty"`
	UserIpAddress []UserIpAddress    `json:"user_ip_address,omitempty"`
	Base          `json:"base,omitempty"`
}
type UserIpAddress struct {
	IpAddress string
	Location Location
	GadgetType string
}
type Verify struct{
	Question string
	Answer string
	Hint string
	Base
}
type Auth struct {
	//User User `gorm:"foreignKey:UserID; not null"`
	UserID string   `json:"userid" bson:"userid"`
	UName string `json:"uname"`
	Token  string `bson:"token"`
	Admin bool
	Supervisor bool
	Employee bool
	Base
}
type LoginUser struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

//Token struct declaration
type Token struct {
	UserID   string
	UName string `json:"uname"`
	Email  string
	Admin bool
	Supervisor bool
	Employee bool
	*jwt.StandardClaims
	Base
}

func (user User) ValidateEmail(email string) (matchedString bool) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&amp;'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	matchedString = re.MatchString(email)
	return matchedString
}
func (user User) ValidatePassword(password string) (bool, *httperrors.HttpError) {
	if len(password) < 5 {
		return false, httperrors.NewBadRequestError("your password need more characters!")
	} else if len(password) > 32 {
		return false, httperrors.NewBadRequestError("your password is way too long!")
	}
	return true, nil
}
func (user User) HashPassword(password string) (string, *httperrors.HttpError) {
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return "", httperrors.NewNotFoundError("type a stronger password!")
	}
	return string(pass), nil

}
func (user LoginUser) Compare(p1, p2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p2), []byte(p1))
	if err != nil {
		return false
	}
	return true
}
func (loginuser LoginUser) Validate() *httperrors.HttpError {
	if loginuser.Email == "" {
		return httperrors.NewNotFoundError("Invalid Email")
	}
	if loginuser.Password == "" {
		return httperrors.NewNotFoundError("Invalid password")
	}
	return nil
}
func (user User) Validate() *httperrors.HttpError {
	if user.FName == "" {
		return httperrors.NewNotFoundError("Invalid first Name")
	}
	if user.LName == "" {
		return httperrors.NewNotFoundError("Invalid last name")
	}
	if user.UName == "" {
		return httperrors.NewNotFoundError("Invalid username")
	}
	if user.Phone == "" {
		return httperrors.NewNotFoundError("Invalid phone number")
	}
	if user.Email == "" {
		return httperrors.NewNotFoundError("Invalid Email")
	}
	// if user.Address == "" {
	// 	return httperrors.NewNotFoundError("Invalid Address")
	// }
	if user.Password == "" {
		return httperrors.NewNotFoundError("Invalid password")
	}
	// if user.Picture == "" {
	// 	return httperrors.NewNotFoundError("Invalid picture")
	// }
	if user.Email == "" {
		return httperrors.NewNotFoundError("Invalid picture")
	}
	return nil
}

func (verify Verify) Validate() *httperrors.HttpError {
	if verify.Question == "" {
		return httperrors.NewNotFoundError("Invalid question")
	}
	if verify.Answer == "" {
		return httperrors.NewNotFoundError("Invalid aswer")
	}
	if verify.Hint == "" && verify.Hint == verify.Answer {
		return httperrors.NewNotFoundError("Invalid hint")
	}
	return nil
}
func (verify Verify) HashAwnser(p string) (string, *httperrors.HttpError) {
	passAnswer, err := bcrypt.GenerateFromPassword([]byte(verify.Answer), 10)
	if err != nil {
		return "", httperrors.NewNotFoundError("type a stronger password!")
	}
	return string(passAnswer), nil

}