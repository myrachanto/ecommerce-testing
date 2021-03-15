package service

import (
	"github.com/myrachanto/ecommerce/httperrors"
	"github.com/myrachanto/ecommerce/model" 
	r "github.com/myrachanto/ecommerce/repository"
)
//Generalservice ...
var (
	Generalservice generalservice = generalservice{}

) 
type generalservice struct {
	
}

func (service generalservice) View(search string) (*model.General, *httperrors.HttpError) {
	general, err1 := r.Generalrepo.View(search)
	if err1 != nil {
		return nil, err1
	}
	 return general, nil

}
// func (service generalservice) Email() (*model.Email, *httperors.HttpError) {
// 	general, err1 := r.generalrepo.Email()
// 	if err1 != nil {
// 		return nil, err1
// 	}
// 	return general, nil
// }

// func (service generalservice) Send() (*model.Email, *httperors.HttpError) {
// 	generals, err := r.generalrepo.Send()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return generals, nil
// }
//db.Where("age = ?", 20).Delete(&User{})