package repository

import (
	"authorize/model"
	_"time"
)


type PersonRepository interface {
	Save(*model.Person) error
	Update(string,*model.Person) error
	Delete(string) error
	FindByID(string)(*model.Person, error)
	FindAll()(model.Persons , error)
	FindByNum(string) (*model.Person ,error)
}


//Token Repository
type TokenRepository interface {
	Save(*model.Token) error
	Update(string, *model.Token) error
	Delete(string) error
	FindByID(string) (*model.Token, error)
	FindAll() (model.Tokens, error)

}

//User Repository
type UserRepository interface {
	Save(*model.User) error
	Update(string, *model.User) error
	Delete(string) error
	FindByID(string) (*model.User, error)
	FindByToken(string) (*model.User, error)
	FindByName(string) (*model.User, error)
	FindByUserName(string) (*model.User, error)
	FindAll() (model.Users, error)
	FindByPozId( string ) (*model.User, error)
}
